package testcases

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"github.com/jhphon0730/action_manager/internal/projects"
	"github.com/jhphon0730/action_manager/internal/teststatus"
	"gorm.io/gorm"
)

// TestCaseService 는 테스트 케이스 관련 서비스 인터페이스
type TestCaseService interface {
	Create(req *CreateTestCaseRequest) error
	FindOne(projectID, testCaseID uint) (*model.TestCase, error)
	Find(projectID uint) ([]*model.TestCase, error)
	UpdateStatus(testCaseID, projectID, userID, currentStatusID uint) error
	UpdateAssignee(testCaseID, projectID, userID, currentAssigneeID uint) error
}

// testCaseService는 테스트 케이스 관련 서비스 구현체입니다.
type testCaseService struct {
	testCaseRepo         TestCaseRepository
	testStatusService    teststatus.TestStatusService
	projectMemberService projects.ProjectMemberService
}

// NewTestCaseService 함수는 TestCaseRepository를 받아 TestCaseService를 생성합니다.
func NewTestCaseService(testCaseRepo TestCaseRepository, testStatusService teststatus.TestStatusService, projectMemberService projects.ProjectMemberService) TestCaseService {
	return &testCaseService{
		testCaseRepo:         testCaseRepo,
		testStatusService:    testStatusService,
		projectMemberService: projectMemberService,
	}
}

// Create 함수는 테스트 케이스를 생성합니다.
func (s *testCaseService) Create(req *CreateTestCaseRequest) error {
	return s.testCaseRepo.WithTx(func(tx *gorm.DB) error {
		testCase := req.ToModel()

		exists := s.testStatusService.IsProjectStatusTx(tx, req.ProjectID, req.CurrentStatusID)
		if !exists {
			return ErrNotInProjectTestStatus
		}

		// 최초 생성 시에 할당자 사용자 정보가 있다면 체크
		if req.CurrentAssigneeID != nil {
			exists, err := s.projectMemberService.IsMemberTx(tx, req.ProjectID, *req.CurrentAssigneeID)
			if !exists || err != nil {
				return ErrNotInProjectMember
			}
		}

		return tx.Create(testCase).Error

	})
}

// Find 함수는 프로젝트 ID와 테스트 케이스 ID에 해당하는 테스트 케이스를 찾습니다.
func (s *testCaseService) FindOne(projectID, testCaseID uint) (*model.TestCase, error) {
	return s.testCaseRepo.Find(projectID, testCaseID)
}

// FindByProjectID 함수는 프로젝트 ID에 해당하는 테스트 케이스를 찾습니다.
func (s *testCaseService) Find(projectID uint) ([]*model.TestCase, error) {
	return s.testCaseRepo.FindByProjectID(projectID)
}

// UpdateStatus 함수는 테스트 케이스의 상태를 업데이트합니다.
func (s *testCaseService) UpdateStatus(testCaseID, projectID, userID, currentStatusID uint) error {
	return s.testCaseRepo.WithTx(func(tx *gorm.DB) error {
		var testCase *model.TestCase

		// 테스트케이스가 존재하는지 확인
		if err := tx.Where("project_id = ? AND id = ?", projectID, testCaseID).First(&testCase).Error; err != nil || testCase == nil {
			if err == nil {
				return ErrNotFound
			}
			return err
		}

		// 프로젝트에 해당하는 status 인지 확인
		exists := s.testStatusService.IsProjectStatusTx(tx, projectID, currentStatusID)
		if !exists {
			return ErrNotInProjectTestStatus
		}

		// 현재 status랑 동일한 status로 변경하려고 하면 에러 반환
		if testCase.CurrentStatusID == currentStatusID {
			return ErrSameStatus
		}

		// status 업데이트
		if err := tx.Model(&model.TestCase{}).Where("project_id = ? AND id = ?", projectID, testCaseID).Update("current_status_id", currentStatusID).Error; err != nil {
			return err
		}

		// 업데이트 내역 추가
		log := model.TestStatusHistory{
			TestCaseID:   testCaseID,
			FromStatusID: testCase.CurrentStatusID,
			ToStatusID:   currentStatusID,
			ChangedByID:  userID,
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}

		return nil
	})
}

// UpdateAssignee 함수는 현재 테스트케이스에 할당된 사용자를 다른 사용자로 변경해주는 함수
func (s *testCaseService) UpdateAssignee(testCaseID, projectID, userID, currentAssigneeID uint) error {
	return s.testCaseRepo.WithTx(func(tx *gorm.DB) error {
		var testCase *model.TestCase
		if err := tx.Where("project_id = ? AND id = ?", projectID, testCaseID).First(&testCase).Error; err != nil || testCase == nil {
			if err == nil {
				return ErrNotFound
			}
			return err
		}

		// 프로젝트에 해당하는 사용자인지 여부 확인
		exists, err := s.projectMemberService.IsMemberTx(tx, projectID, currentAssigneeID)
		if err != nil || !exists {
			return ErrNotInProjectMember
		}

		// 기존이랑 동일하다면 ...
		if testCase.CurrentAssigneeID != nil && *testCase.CurrentAssigneeID == currentAssigneeID {
			return ErrSameAssignee
		}

		// assignee 업데이트
		if err := tx.Model(&model.TestCase{}).Where("project_id = ? AND id = ?", projectID, testCaseID).Update("assignee_id", currentAssigneeID).Error; err != nil {
			return err
		}

		log := model.TestAssigneeHistory{
			TestCaseID:     testCaseID,
			FromAssigneeID: testCase.CurrentAssigneeID,
			ToAssigneeID:   &currentAssigneeID,
			ChangedByID:    userID,
		}

		if err := tx.Create(&log).Error; err != nil {
			return err
		}

		return nil
	})
}

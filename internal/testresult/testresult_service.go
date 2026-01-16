package testresults

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"github.com/jhphon0730/action_manager/internal/testcases"
)

// TestResultService 인터페이스는 테스트 결과를 관리하는 서비스를 정의합니다.
type TestResultService interface {
	Create(req *CreateTestResultRequest, projectID uint) error
	Find(projectID, testCaseID uint) ([]*model.TestResult, error)
}

// testResultService 구조체는 TestResultService를 구현합니다.
type testResultService struct {
	testResultRepo  TestResultRepository
	testCaseService testcases.TestCaseService
}

// NewTestResultService 함수는 새로운 TestResultService를 생성합니다.
func NewTestResultService(testResultRepo TestResultRepository, testCaseService testcases.TestCaseService) TestResultService {
	return &testResultService{
		testResultRepo:  testResultRepo,
		testCaseService: testCaseService,
	}
}

// Create 함수는 새로운 테스트 결과를 생성합니다.
func (s *testResultService) Create(req *CreateTestResultRequest, projectID uint) error {
	if testCase, err := s.testCaseService.Find(projectID, req.TestCaseID); err != nil || testCase == nil {
		return ErrNotFoundTestCase
	}

	testResult := req.ToModel()
	return s.testResultRepo.Create(testResult)
}

// Find 함수는 프로젝트 아이디와 테스트케이스 아이디로 테스트케이스 결과를 모두 조회
func (s *testResultService) Find(projectID, testCaseID uint) ([]*model.TestResult, error) {
	if testCase, err := s.testCaseService.Find(projectID, testCaseID); err != nil || testCase == nil {
		return nil, ErrNotFoundTestCase
	}

	return s.testResultRepo.Find(testCaseID)
}

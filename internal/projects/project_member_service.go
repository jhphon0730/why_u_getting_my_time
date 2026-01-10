package projects

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// ProjectMemberService는 프로젝트 멤버 관리를 위한 인터페이스입니다.
type ProjectMemberService interface {
	Create(req *CreateProjectMemberRequest) error
	FindByProjectID(projectID uint) ([]*model.ProjectMember, error)
	UpdateRoleToManager(projectID, userID uint) error
	UpdateRoleToMember(projectID, userID uint) error
	Delete(projectID, userID uint) error
	IsMember(projectID, userID uint) (bool, error)
	IsManager(projectID, userID uint) (bool, error)

	IsMemberTx(tx *gorm.DB, projectID, userID uint) (bool, error)
}

// NewProjectMemberService는 새로운 ProjectMemberService 인스턴스를 생성합니다.
type projectMemberService struct {
	projectMemberRepo ProjectMemberRepository
}

// NewProjectMemberService는 새로운 ProjectMemberService 인스턴스를 생성합니다.
func NewProjectMemberService(projectMemberRepo ProjectMemberRepository) ProjectMemberService {
	return &projectMemberService{
		projectMemberRepo: projectMemberRepo,
	}
}

// Create 프로젝트 멤버를 추가합니다.
func (s *projectMemberService) Create(req *CreateProjectMemberRequest) error {
	if exists, _ := s.projectMemberRepo.IsMember(req.ProjectID, req.UserID); exists {
		return ErrAlreadyMember
	}

	project := req.ToModel()
	return s.projectMemberRepo.Create(project)
}

// FindByProjectID 프로젝트에 속한 멤버 목록을 조회합니다.
func (s *projectMemberService) FindByProjectID(projectID uint) ([]*model.ProjectMember, error) {
	return s.projectMemberRepo.FindByProjectID(projectID)
}

// UpdateRoleToManager 함수는 프로젝트 멤버의 역할을 관리자로 업데이트합니다.
func (s *projectMemberService) UpdateRoleToManager(projectID, userID uint) error {
	return s.projectMemberRepo.UpdateRoleToManager(projectID, userID)
}

// UpdateRoleToMember 함수는 프로젝트 멤버의 역할을 멤버로 업데이트합니다.
func (s *projectMemberService) UpdateRoleToMember(projectID, userID uint) error {
	// 프로젝트에 해당하는 사용자인지 확인
	isMember, err := s.projectMemberRepo.IsMember(projectID, userID)
	if err != nil {
		return err
	}

	if !isMember {
		return ErrNotMember
	}

	isManager, err := s.projectMemberRepo.IsManager(projectID, userID)
	if err != nil {
		return err
	}

	// 매니저의 경우 마지막 매니저인이 확인
	if isManager {
		memberCtn, err := s.projectMemberRepo.CountManagers(projectID)
		if err != nil {
			return err
		}

		if memberCtn == 1 {
			return ErrLastManager
		}
	}

	return s.projectMemberRepo.UpdateRoleToMember(projectID, userID)
}

// Delete 프로젝트 멤버를 삭제합니다.
func (s *projectMemberService) Delete(projectID, userID uint) error {
	// 프로젝트에 해당하는 사용자인지 확인
	isMember, err := s.projectMemberRepo.IsMember(projectID, userID)
	if err != nil {
		return err
	}

	if !isMember {
		return ErrNotMember
	}

	isManager, err := s.projectMemberRepo.IsManager(projectID, userID)
	if err != nil {
		return err
	}

	// 매니저의 경우 마지막 매니저인이 확인
	if isManager {
		memberCtn, err := s.projectMemberRepo.CountManagers(projectID)
		if err != nil {
			return err
		}

		if memberCtn == 1 {
			return ErrLastManager
		}
	}

	return s.projectMemberRepo.Delete(projectID, userID)
}

// IsMember는 사용자가 프로젝트에 멤버인지 확인합니다.
func (s *projectMemberService) IsMember(projectID, userID uint) (bool, error) {
	return s.projectMemberRepo.IsMember(projectID, userID)
}

// IsManager는 사용자가 프로젝트 관리자인지 확인합니다.
func (s *projectMemberService) IsManager(projectID, userID uint) (bool, error) {
	return s.projectMemberRepo.IsManager(projectID, userID)
}

// IsMember는 사용자가 프로젝트에 멤버인지 확인합니다.
func (s *projectMemberService) IsMemberTx(tx *gorm.DB, projectID, userID uint) (bool, error) {
	return s.projectMemberRepo.NewWithTx(tx).IsMember(projectID, userID)
}

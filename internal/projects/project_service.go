package projects

import (
	"time"

	"github.com/jhphon0730/action_manager/internal/model"
	"github.com/jhphon0730/action_manager/internal/teststatus"
	"gorm.io/gorm"
)

// ProjectService는 프로젝트 관련 서비스 인터페이스입니다.
type ProjectService interface {
	Create(req *CreateProjectRequest, userID uint) error
	FindAllByUserID(userID uint) ([]*model.Project, error)
}

// projectService는 프로젝트 관련 서비스 구현체입니다.
type projectService struct {
	projectRepo       ProjectRepository
	teststatusService teststatus.TestStatusService
}

// NewProjectService는 프로젝트 서비스를 생성합니다.
func NewProjectService(projectRepo ProjectRepository, teststatusService teststatus.TestStatusService) ProjectService {
	return &projectService{
		projectRepo:       projectRepo,
		teststatusService: teststatusService,
	}
}

// Create는 프로젝트를 생성합니다.
func (s *projectService) Create(req *CreateProjectRequest, userID uint) error {
	return s.projectRepo.WithTx(func(tx *gorm.DB) error {
		project := req.ToModel()

		if err := tx.Create(project).Error; err != nil {
			return err
		}

		member := &model.ProjectMember{
			ProjectID:   project.ID,
			UserID:      userID,
			ProjectRole: model.RoleManager,
			JoinedAt:    time.Now(),
		}

		if err := tx.Create(member).Error; err != nil {
			return err
		}

		return s.teststatusService.CreateDefaultTx(tx, project.ID)
	})
}

// FindAllByUserID는 사용자 ID로 프로젝트를 조회합니다.
func (s *projectService) FindAllByUserID(userID uint) ([]*model.Project, error) {
	return s.projectRepo.FindAllByUserID(userID)
}

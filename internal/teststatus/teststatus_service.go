package teststatus

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// TestStatusService는 테스트 상태를 관리하는 서비스 인터페이스입니다.
type TestStatusService interface {
	CreateDefault(projectID uint) error
	CreateDefaultTx(tx *gorm.DB, projectID uint) error
	Create(req *CreateTestStatusRequest) error
	FindByProjectID(projectID uint) ([]*model.TestStatus, error)
	Delete(projectID, statusID uint) error
}

// testStatusService는 테스트 상태를 관리하는 구현체입니다.
type testStatusService struct {
	repo TestStatusRepository
}

// NewTestStatusService는 테스트 상태 서비스를 생성합니다.
func NewTestStatusService(repo TestStatusRepository) TestStatusService {
	return &testStatusService{
		repo: repo,
	}
}

// CreateDefault는 기본 테스트 상태를 생성합니다.
func (s *testStatusService) CreateDefault(projectID uint) error {
	return s.repo.CreateDefault(projectID)
}

// CreateDefaultTx는 기본 테스트 상태를 생성합니다.
func (s *testStatusService) CreateDefaultTx(tx *gorm.DB, projectID uint) error {
	return s.repo.WithTx(tx).CreateDefault(projectID)
}

// Create는 테스트 상태를 생성합니다.
func (s *testStatusService) Create(req *CreateTestStatusRequest) error {
	status := req.ToModel()

	return s.repo.Create(status)
}

// FindByProjectID는 프로젝트 ID에 해당하는 테스트 상태를 조회합니다.
func (s *testStatusService) FindByProjectID(projectID uint) ([]*model.TestStatus, error) {
	return s.repo.FindByProjectID(projectID)
}

// Delete는 프로젝트 ID와 상태 ID에 해당하는 테스트 상태를 삭제합니다.
func (s *testStatusService) Delete(projectID, statusID uint) error {
	return s.repo.Delete(projectID, statusID)
}

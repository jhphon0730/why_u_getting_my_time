package teststatus

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// TestStatusService는 테스트 상태를 관리하는 서비스 인터페이스입니다.
type TestStatusService interface {
	CreateDefault(projectID uint) error
	CreateDefaultTx(tx *gorm.DB, projectID uint) error
	FindByProjectID(projectID uint) ([]*model.TestStatus, error)
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

// FindByProjectID는 프로젝트 ID에 해당하는 테스트 상태를 조회합니다.
func (s *testStatusService) FindByProjectID(projectID uint) ([]*model.TestStatus, error) {
	return s.repo.FindByProjectID(projectID)
}

//

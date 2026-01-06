package teststatus

import "gorm.io/gorm"

// TestStatusService는 테스트 상태를 관리하는 서비스 인터페이스입니다.
type TestStatusService interface {
	CreateDefault(projectID uint) error
	CreateDefaultTx(tx *gorm.DB, projectID uint) error
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

package teststatus

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// TestStatusRepository는 테스트 상태를 관리하는 인터페이스입니다.
type TestStatusRepository interface {
	WithTx(tx *gorm.DB) TestStatusRepository

	CreateDefault(projectID uint) error
	FindByProjectID(projectID uint) ([]*model.TestStatus, error)
}

// testStatusRepository는 테스트 상태를 관리하는 구현체입니다.
type testStatusRepository struct {
	db *gorm.DB
}

// NewTestStatusRepository는 TestStatusRepository 인스턴스를 생성합니다.
func NewTestStatusRepository(db *gorm.DB) TestStatusRepository {
	return &testStatusRepository{
		db: db,
	}
}

// WithTx는 트랜잭션을 사용하여 새로운 TestStatusRepository 인스턴스를 생성합니다.
func (r *testStatusRepository) WithTx(tx *gorm.DB) TestStatusRepository {
	return &testStatusRepository{
		db: tx,
	}
}

// CreateDefault는 프로젝트 생성 시에 기본으로 추가되는 Status를 생성함.
func (r *testStatusRepository) CreateDefault(projectID uint) error {
	for _, status := range model.DEFAULT_TESTSTATUS_DATA {
		newStatus := &model.TestStatus{
			Name:      status,
			ProjectID: projectID,
			IsActive:  true,
		}

		if err := r.db.Create(newStatus).Error; err != nil {
			return err
		}
	}

	return nil
}

// FindByProjectID 함수는 프로젝트 ID를 기반으로 테스트 상태를 조회합니다.
func (r *testStatusRepository) FindByProjectID(projectID uint) ([]*model.TestStatus, error) {
	var testCases []*model.TestStatus
	if err := r.db.Where("project_id = ?", projectID).Find(&testCases).Error; err != nil {
		return nil, err
	}

	return testCases, nil
}

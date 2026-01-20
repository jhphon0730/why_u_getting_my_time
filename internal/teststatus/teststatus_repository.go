package teststatus

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// TestStatusRepository는 테스트 상태를 관리하는 인터페이스입니다.
type TestStatusRepository interface {
	NewWithTx(tx *gorm.DB) TestStatusRepository

	CreateDefault(projectID uint) error
	Create(status *model.TestStatus) error
	Delete(projectID, statusID uint) error
	Find(projectID uint) ([]*model.TestStatus, error)
	IsProjectStatus(projectID, statusID uint) bool
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
func (r *testStatusRepository) NewWithTx(tx *gorm.DB) TestStatusRepository {
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

// Create 함수는 새로운 테스트 상태를 생성합니다.
func (r *testStatusRepository) Create(status *model.TestStatus) error {
	return r.db.Create(status).Error
}

// Find 함수는 프로젝트 ID를 기반으로 테스트 상태를 조회합니다.
func (r *testStatusRepository) Find(projectID uint) ([]*model.TestStatus, error) {
	var testCases []*model.TestStatus
	if err := r.db.Where("project_id = ?", projectID).Find(&testCases).Error; err != nil {
		return nil, err
	}

	return testCases, nil
}

// Delete 함수는 프로젝트 ID와 상태 ID를 기반으로 테스트 상태를 삭제합니다.
func (r *testStatusRepository) Delete(projectID, statusID uint) error {
	return r.db.Where("project_id = ? AND id = ?", projectID, statusID).Delete(&model.TestStatus{}).Error
}

// IsProjectStatus 함수는 프로젝트 ID와 상태 ID를 기반으로 테스트 상태가 존재하는지 확인합니다.
func (r *testStatusRepository) IsProjectStatus(projectID, statusID uint) bool {
	var count int64
	r.db.Model(&model.TestStatus{}).Where("project_id = ? AND id = ?", projectID, statusID).Count(&count)
	return count > 0
}

package testcases

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// TestCaseRepository 는 테스트 케이스를 관리하는 인터페이스입니다.
type TestCaseRepository interface {
	WithTx(fn func(tx *gorm.DB) error) error

	Create(testcase *model.TestCase) error
	FindByProjectID(projectID uint) ([]*model.TestCase, error)
}

// TestCaseRepository 는 테스트 케이스를 관리하는 구현체입니다.
type testCaseRepository struct {
	db *gorm.DB
}

// NewTestCaseRepository 함수는 새로운 TestCaseRepository를 생성합니다.
func NewTestCaseRepository(db *gorm.DB) TestCaseRepository {
	return &testCaseRepository{
		db: db,
	}
}

// WithTx는 트랜잭션을 사용하여 함수를 실행하는 함수입니다.
func (r *testCaseRepository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

// Create 함수는 새로운 테스트 케이스를 생성합니다.
func (r *testCaseRepository) Create(testcase *model.TestCase) error {
	return r.db.Create(testcase).Error
}

// FindByProjectID 함수는 프로젝트 ID에 해당하는 테스트 케이스를 찾습니다.
func (r *testCaseRepository) FindByProjectID(projectID uint) ([]*model.TestCase, error) {
	var testcases []*model.TestCase
	err := r.db.Where("project_id = ?", projectID).Find(&testcases).Error
	return testcases, err
}

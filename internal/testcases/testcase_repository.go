package testcases

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// TestCaseRepository 는 테스트 케이스를 관리하는 인터페이스입니다.
type TestCaseRepository interface {
	Create(testcase *model.TestCase) error
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

// Create 함수는 새로운 테스트 케이스를 생성합니다.
func (r *testCaseRepository) Create(testcase *model.TestCase) error {
	return r.db.Create(testcase).Error
}

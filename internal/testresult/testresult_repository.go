package testresults

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// TestResultRepository 인터페이스는 TestResultRepository의 메서드를 정의합니다.
type TestResultRepository interface {
	Create(testResult *model.TestResult) error
}

// TestResultRepository 인터페이스를 구현하는 구조체입니다.
type testResultRepository struct {
	db *gorm.DB
}

// NewTestResultRepository 함수는 새로운 TestResultRepository를 생성합니다.
func NewTestResultRepository(db *gorm.DB) TestResultRepository {
	return &testResultRepository{db: db}
}

// Create 함수는 새로운 테스트 결과를 생성합니다
func (r *testResultRepository) Create(testResult *model.TestResult) error {
	return r.db.Create(testResult).Error
}

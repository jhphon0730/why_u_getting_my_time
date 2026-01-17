package testresults

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// TestResultRepository 인터페이스는 TestResultRepository의 메서드를 정의합니다.
type TestResultRepository interface {
	Create(testResult *model.TestResult) error
	Find(testCaseID uint) ([]*model.TestResult, error)
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

// Find 함수는 프로젝트 아이디와 테스트케이스 아이디로 테스트케이스 결과를 모두 조회
func (r *testResultRepository) Find(testCaseID uint) ([]*model.TestResult, error) {
	var testResults []*model.TestResult
	if err := r.db.Where("test_case_id = ?", testCaseID).Find(&testResults).Error; err != nil {
		return nil, err
	}
	return testResults, nil
}

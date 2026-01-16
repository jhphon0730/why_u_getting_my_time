package testresults

import "github.com/jhphon0730/action_manager/internal/testcases"

// TestResultService 인터페이스는 테스트 결과를 관리하는 서비스를 정의합니다.
type TestResultService interface {
	Create(req *CreateTestResultRequest, projectID uint) error
}

// testResultService 구조체는 TestResultService를 구현합니다.
type testResultService struct {
	testResultRepo  TestResultRepository
	testCaseService testcases.TestCaseService
}

// NewTestResultService 함수는 새로운 TestResultService를 생성합니다.
func NewTestResultService(testResultRepo TestResultRepository, testCaseService testcases.TestCaseService) TestResultService {
	return &testResultService{
		testResultRepo:  testResultRepo,
		testCaseService: testCaseService,
	}
}

// Create 함수는 새로운 테스트 결과를 생성합니다.
func (s *testResultService) Create(req *CreateTestResultRequest, projectID uint) error {
	if testCase, err := s.testCaseService.Find(projectID, req.TestCaseID); err != nil || testCase == nil {
		return ErrNotFoundTestCase
	}

	testResult := req.ToModel()
	return s.testResultRepo.Create(testResult)
}

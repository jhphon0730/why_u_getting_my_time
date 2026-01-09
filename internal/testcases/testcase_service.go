package testcases

// TestCaseService 는 테스트 케이스 관련 서비스 인터페이스
type TestCaseService interface {
	Create(req *CreateTestCaseRequest) error
}

// testCaseService는 테스트 케이스 관련 서비스 구현체입니다.
type testCaseService struct {
	testCaseRepo TestCaseRepository
}

// NewTestCaseService 함수는 TestCaseRepository를 받아 TestCaseService를 생성합니다.
func NewTestCaseService(testCaseRepo TestCaseRepository) TestCaseService {
	return &testCaseService{
		testCaseRepo: testCaseRepo,
	}
}

// Create 함수는 테스트 케이스를 생성합니다.
func (s *testCaseService) Create(req *CreateTestCaseRequest) error {
	return s.testCaseRepo.Create(req.ToModel())
}

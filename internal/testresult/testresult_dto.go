package testresults

import "github.com/jhphon0730/action_manager/internal/model"

// CreateTestResultRequest 구조체는 새로운 테스트 결과를 생성하는데 사용되는 DTO Request
type CreateTestResultRequest struct {
	TestCaseID uint   `json:"test_case_id"`
	Result     string `json:"result"`
	Comment    string `json:"comment"`
}

// ToModel 함수는 CreateTestResultRequest를 model.TestResult로 변환하는 함수
func (d *CreateTestResultRequest) ToModel() *model.TestResult {
	return &model.TestResult{
		TestCaseID: d.TestCaseID,
		Result:     d.Result,
		Comment:    d.Comment,
	}
}

// TestResultResponse 구조체는 응답에 사용되는 구조체
type TestResultResponse struct {
	ID         uint   `json:"id"`
	TestCaseID uint   `json:"test_case_id"`
	Result     string `json:"result"`
	Comment    string `json:"comment"`
}

// ToModelTestResultResponse 함수는 원본 테스크결과 모델 데이터를 TestResultResponse로 변환하는 함수
func ToModelTestResultResponse(testResult *model.TestResult) *TestResultResponse {
	return &TestResultResponse{
		ID:         testResult.ID,
		TestCaseID: testResult.TestCaseID,
		Result:     testResult.Result,
		Comment:    testResult.Comment,
	}
}

// ToModelTestResultResponseList 함수는 원본 테스크결과 모델 데이터 리스트를 TestResultResponse 리스트로 변환하는 함수
func ToModelTestResultResponseList(testResults []*model.TestResult) []*TestResultResponse {
	res := make([]*TestResultResponse, len(testResults))
	for i, testResult := range testResults {
		res[i] = ToModelTestResultResponse(testResult)
	}
	return res
}

package teststatus

import "github.com/jhphon0730/action_manager/internal/model"

// CreateTestStatusRequest 구조체는 테스트 상태 생성 요청을 나타냅니다.
type CreateTestStatusRequest struct {
	ProjectID uint   `json:"project_id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
}

// ToModel 함수는 CreateTestStatusRequest를 model.TestStatus로 변환합니다.
func (req *CreateTestStatusRequest) ToModel() *model.TestStatus {
	return &model.TestStatus{
		ProjectID: req.ProjectID,
		Name:      req.Name,
		IsActive:  req.IsActive, // default : TRUE
	}
}

// TestStatusResponse 구조체는 테스트 상태에 대한 응답을 나타냅니다.
type TestStatusResponse struct {
	ID        uint   `json:"id"`
	ProjectID uint   `json:"project_id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
}

// ToModelTestStatusResponse 함수는 테스트 상태 모델을 테스트 상태 응답으로 변환합니다.
func ToModelTestStatusResponse(status *model.TestStatus) *TestStatusResponse {
	return &TestStatusResponse{
		ID:        status.ID,
		ProjectID: status.ProjectID,
		Name:      status.Name,
		IsActive:  status.IsActive,
	}
}

// ToModelTestStatusResponseList 함수는 테스트 상태 모델 리스트를 테스트 상태 응답 리스트로 변환합니다.
func ToModelTestStatusResponseList(statuses []*model.TestStatus) []*TestStatusResponse {
	var responses []*TestStatusResponse
	for _, status := range statuses {
		responses = append(responses, ToModelTestStatusResponse(status))
	}
	return responses
}

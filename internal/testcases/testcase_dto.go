package testcases

import (
	"time"

	"github.com/jhphon0730/action_manager/internal/model"
)

// CreateTestCaseRequest 는 테스트케이스 생성 시에 사용되는 DTO입니다.
type CreateTestCaseRequest struct {
	ProjectID   uint   `json:"project_id"` // project__id
	Title       string `json:"title"`
	Description string `json:"description"`

	CurrentStatusID   uint  `json:"current_status_id"`   // teststatus__id
	CurrentAssigneeID *uint `json:"current_assignee_id"` // user__id

	DueDate *time.Time `json:"due_date"`
}

// ToModel 함수는 CreateTestCaseRequest를 model.TestCase로 변환합니다.
func (d *CreateTestCaseRequest) ToModel() *model.TestCase {
	return &model.TestCase{
		ProjectID:         d.ProjectID,
		Title:             d.Title,
		Description:       d.Description,
		CurrentStatusID:   d.CurrentStatusID,
		CurrentAssigneeID: d.CurrentAssigneeID,
		DueDate:           d.DueDate,
	}
}

// TestCaseResponse 구조체는 테스트케이스의 응답 정보를 담는 DTO입니다.
type TestCaseResponse struct {
	ID                uint       `json:"id"`
	ProjectID         uint       `json:"project_id"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	CurrentStatusID   uint       `json:"current_status_id"`
	CurrentAssigneeID *uint      `json:"current_assignee_id"`
	DueDate           *time.Time `json:"due_date"`
}

// ToModelTestCaseResponse 함수는 model.TestCase를 TestCaseResponse로 변환합니다.
func ToModelTestCaseResponse(testcase *model.TestCase) *TestCaseResponse {
	return &TestCaseResponse{
		ID:                testcase.ID,
		ProjectID:         testcase.ProjectID,
		Title:             testcase.Title,
		Description:       testcase.Description,
		CurrentStatusID:   testcase.CurrentStatusID,
		CurrentAssigneeID: testcase.CurrentAssigneeID,
		DueDate:           testcase.DueDate,
	}
}

// ToModelTestCaseResponseList 함수는 model.TestCase 리스트를 TestCaseResponse 리스트로 변환합니다.
func ToModelTestCaseResponseList(testcases []*model.TestCase) []*TestCaseResponse {
	res := make([]*TestCaseResponse, len(testcases))
	for i, testcase := range testcases {
		res[i] = ToModelTestCaseResponse(testcase)
	}
	return res
}

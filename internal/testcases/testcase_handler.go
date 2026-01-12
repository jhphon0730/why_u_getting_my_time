package testcases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

// TestCaseHandler 는 테스트 케이스를 관리하는 인터페이스입니다.
type TestCaseHandler interface {
	Create(c *gin.Context)
	FindByProjectID(c *gin.Context)
}

// TestCaseHandler 는 테스트 케이스를 관리하는 구현체입니다.
type testCaseHandler struct {
	testCaseService TestCaseService
}

// NewTestCaseHandler 함수는 새로운 TestCaseHandler를 생성합니다.
func NewTestCaseHandler(testCaseService TestCaseService) TestCaseHandler {
	return &testCaseHandler{
		testCaseService: testCaseService,
	}
}

// Create 함수는 새로운 테스트 케이스를 생성합니다.
func (h *testCaseHandler) Create(c *gin.Context) {
	var req CreateTestCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error()+" / "+err.Error())
		return
	}

	// middleware에서 이미 projectID 유무 확인 완료
	projectID, _ := contextutils.GetProjectIDByParam(c)
	req.ProjectID = projectID

	if err := h.testCaseService.Create(&req); err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondCreated(c, nil)
}

// FindByProjectID 함수는 특정 프로젝트의 테스트 케이스를 조회합니다.
func (h *testCaseHandler) FindByProjectID(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)

	testCases, err := h.testCaseService.FindByProjectID(projectID)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondOK(c, ToModelTestCaseResponseList(testCases))
}

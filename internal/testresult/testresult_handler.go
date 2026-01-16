package testresults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

// TestResultHandler 인터페이스는 테스트 결과를 처리 관련 요청을 처리하는 인터페이스입니다.
type TestResultHandler interface {
	Create(c *gin.Context)
}

// testResultHandler 구조체는 테스트 결과를 처리하는 구조체입니다.
type testResultHandler struct {
	testResultService TestResultService
}

// NewTestResultHandler 함수는 테스트 결과를 처리하는 인터페이스를 생성합니다.
func NewTestResultHandler(testResultService TestResultService) TestResultHandler {
	return &testResultHandler{
		testResultService: testResultService,
	}
}

// Create 함수는 새로운 테스트 결과를 생성합니다.
func (h *testResultHandler) Create(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)

	var req CreateTestResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.testResultService.Create(&req, projectID); err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondCreated(c, nil)
}

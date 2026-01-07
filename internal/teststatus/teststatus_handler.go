package teststatus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

// TestStatusHandler는 테스트 상태를 관리하는 기능을 제공
type TestStatusHandler interface {
	FindByProjectID(c *gin.Context)
}

// testStatusHandler 구조체는 TestStatusHandler 인터페이스를 구현합니다.
type testStatusHandler struct {
	testStatusService TestStatusService
}

// NewTestStatusHandler 함수는 TestStatusHandler 인터페이스를 반환합니다.
func NewTestStatusHandler(testStatusService TestStatusService) TestStatusHandler {
	return &testStatusHandler{
		testStatusService: testStatusService,
	}
}

// FindByProjectID 함수는 특정 프로젝트의 테스트 상태를 조회합니다.
func (h *testStatusHandler) FindByProjectID(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)

	status, err := h.testStatusService.FindByProjectID(projectID)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
	}

	response.RespondOK(c, gin.H{
		"test_status": ToModelTestStatusResponseList(status),
	})
}

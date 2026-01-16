package teststatus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

// TestStatusHandler는 테스트 상태를 관리하는 기능을 제공
type TestStatusHandler interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
	Delete(c *gin.Context)
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

// Create 함수는 새로운 테스트케이스 상태를 추가합니다.
func (h *testStatusHandler) Create(c *gin.Context) {
	var createTestStatusRequst CreateTestStatusRequest
	if err := c.ShouldBindJSON(&createTestStatusRequst); err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}

	projectID, _ := contextutils.GetProjectIDByParam(c)
	createTestStatusRequst.ProjectID = projectID
	createTestStatusRequst.IsActive = true

	err := h.testStatusService.Create(&createTestStatusRequst)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondCreated(c, gin.H{
		"message": "Create Success.",
	})
}

// FindByProjectID 함수는 특정 프로젝트의 테스트 상태를 조회합니다.
func (h *testStatusHandler) Find(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)

	status, err := h.testStatusService.FindByProjectID(projectID)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
	}

	response.RespondOK(c, gin.H{
		"test_status": ToModelTestStatusResponseList(status),
		"message":     "Find Success.",
	})
}

// Delete 함수는 특정 테스트 상태를 삭제합니다.
func (h *testStatusHandler) Delete(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)
	statusID, err := contextutils.GetStatusIDByParam(c)
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
	}

	err = h.testStatusService.Delete(projectID, statusID)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondNoContent(c)
}

package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

// ProjectMemberHandler는 프로젝트 멤버를 관리하는 인터페이스입니다.
type ProjectMemberHandler interface {
	CreateMember(c *gin.Context)
}

// ProjectMemberHandler는 프로젝트 멤버를 관리하는 구현체입니다.
type projectMemberHandler struct {
	projectMemberSer ProjectMemberService
}

// NewProjectMemberHandler는 ProjectMemberHandler를 생성합니다.
func NewProjectMemberHandler(projectMemberSer ProjectMemberService) ProjectMemberHandler {
	return &projectMemberHandler{
		projectMemberSer: projectMemberSer,
	}
}

// CreateMember는 프로젝트 멤버를 생성합니다.
func (h *projectMemberHandler) CreateMember(c *gin.Context) {
	var createMemberRequest CreateProjectMemberRequest
	if err := c.ShouldBindJSON(&createMemberRequest); err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}

	// checked middleware
	projectID, _ := contextutils.GetProjectIDByParam(c)
	createMemberRequest.ProjectID = projectID

	if err := h.projectMemberSer.Create(&createMemberRequest); err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondCreated(c, nil)
}

package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

// ProjectMemberHandler는 프로젝트 멤버를 관리하는 인터페이스입니다.
type ProjectMemberHandler interface {
	AddMember(c *gin.Context)
	UpdateRoleToManager(c *gin.Context)
	UpdateRoleToMember(c *gin.Context)
	DeleteMember(c *gin.Context)
	FindByProjectID(c *gin.Context)
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
func (h *projectMemberHandler) AddMember(c *gin.Context) {
	var createMemberRequest CreateProjectMemberRequest

	userID, ok := contextutils.GetUserIDIDByParam(c)
	if !ok {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidUserID.Error())
		return
	}
	createMemberRequest.UserID = userID

	// checked middleware
	projectID, _ := contextutils.GetProjectIDByParam(c)
	createMemberRequest.ProjectID = projectID

	if err := h.projectMemberSer.Create(&createMemberRequest); err != nil {
		if err == ErrAlreadyMember {
			response.RespondError(c, http.StatusConflict, ErrAlreadyMember.Error())
			return
		}
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondCreated(c, nil)
}

// UpdateRoleToManager는 프로젝트 멤버의 역할을 관리자로 업데이트합니다.
func (h *projectMemberHandler) UpdateRoleToManager(c *gin.Context) {
	projectID, exists := contextutils.GetProjectIDByParam(c)
	if !exists {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidProjectID.Error())
		return
	}
	userID, exists := contextutils.GetUserIDIDByParam(c)
	if !exists {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidUserID.Error())
		return
	}

	if err := h.projectMemberSer.UpdateRoleToManager(projectID, userID); err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondCreated(c, nil)
}

// UpdateRoleToMember는 프로젝트 멤버의 역할을 멤버로 업데이트합니다.
func (h *projectMemberHandler) UpdateRoleToMember(c *gin.Context) {
	projectID, exists := contextutils.GetProjectIDByParam(c)
	if !exists {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidProjectID.Error())
		return
	}
	userID, exists := contextutils.GetUserIDIDByParam(c)
	if !exists {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidUserID.Error())
		return
	}

	if err := h.projectMemberSer.UpdateRoleToMember(projectID, userID); err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondCreated(c, nil)
}

// DeleteMember는 프로젝트 멤버를 삭제합니다.
func (h *projectMemberHandler) DeleteMember(c *gin.Context) {
	projectID, exists := contextutils.GetProjectIDByParam(c)
	if !exists {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidProjectID.Error())
		return
	}
	userID, exists := contextutils.GetUserIDIDByParam(c)
	if !exists {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidUserID.Error())
		return
	}

	if err := h.projectMemberSer.Delete(projectID, userID); err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondNoContent(c)
}

// FindByProjectID는 프로젝트 멤버 목록을 조회합니다.
func (h *projectMemberHandler) FindByProjectID(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)

	members, err := h.projectMemberSer.FindByProjectID(projectID)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondOK(c, gin.H{
		"members": ToModelMemberResponseList(members),
	})
}

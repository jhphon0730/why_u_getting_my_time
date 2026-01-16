package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

// ProjectHandler는 프로젝트를 관리하는 인터페이스입니다.
type ProjectHandler interface {
	CreateProject(c *gin.Context)
	GetAllProjects(c *gin.Context)
}

// projectHandler는 ProjectHandler를 구현하는 구조체입니다.
type projectHandler struct {
	projectService       ProjectService
	projectMemberService ProjectMemberService
}

// NewProjectHandler는 새로운 ProjectHandler를 생성합니다.
func NewProjectHandler(projectService ProjectService, projectMemberService ProjectMemberService) ProjectHandler {
	return &projectHandler{
		projectService:       projectService,
		projectMemberService: projectMemberService,
	}
}

// CreateProject는 새로운 프로젝트를 생성합니다.
func (h *projectHandler) CreateProject(c *gin.Context) {
	userID, exists := contextutils.GetUserID(c)
	if !exists {
		response.RespondError(c, http.StatusUnauthorized, ErrUnauthorized.Error())
		return
	}

	var createProjectRequest CreateProjectRequest
	if err := c.ShouldBindJSON(&createProjectRequest); err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}

	if err := h.projectService.Create(&createProjectRequest, userID); err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondSuccess(c, http.StatusCreated, gin.H{
		"message": "Create Success.",
	})
}

// GetProjects는 사용자의 모든 프로젝트를 가져옵니다.
func (h *projectHandler) GetAllProjects(c *gin.Context) {
	userID, exists := contextutils.GetUserID(c)
	if !exists {
		response.RespondError(c, http.StatusUnauthorized, ErrUnauthorized.Error())
		return
	}

	projects, err := h.projectService.FindAllByUserID(userID)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondSuccess(c, http.StatusOK, gin.H{
		"projects": ToModelListResponse(projects),
	})
}

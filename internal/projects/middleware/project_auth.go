package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/projects"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

// RequireProjectManager는 프로젝트 관리자 권한을 요구하는 미들웨어입니다.
func RequireProjectManager(projectMemberService projects.ProjectMemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := contextutils.GetUserID(c)
		if !ok {
			response.RespondError(c, http.StatusUnauthorized, UnauthorizedAuth.Error())
			c.Abort()
			return
		}

		projectID, ok := contextutils.GetProjectIDByParam(c)
		if !ok {
			log.Println(ok, projectID)
			response.RespondError(c, http.StatusUnauthorized, UnauthorizedProject.Error())
			c.Abort()
			return
		}

		isManager, err := projectMemberService.IsManager(projectID, userID)
		if err != nil {
			response.RespondError(c, http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		if !isManager {
			response.RespondError(c, http.StatusForbidden, PermissionRequired.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireProjectMember는 프로젝트 멤버 권한을 요구하는 미들웨어입니다.
func RequireProjectMember(projectMemberService projects.ProjectMemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := contextutils.GetUserID(c)
		if !ok {
			response.RespondError(c, http.StatusUnauthorized, UnauthorizedAuth.Error())
			c.Abort()
			return
		}

		projectID, ok := contextutils.GetProjectIDByParam(c)
		if !ok {
			response.RespondError(c, http.StatusUnauthorized, UnauthorizedProject.Error())
			c.Abort()
			return
		}

		isMember, err := projectMemberService.IsMember(projectID, userID)
		if err != nil {
			response.RespondError(c, http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		if !isMember {
			response.RespondError(c, http.StatusForbidden, PermissionRequired.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

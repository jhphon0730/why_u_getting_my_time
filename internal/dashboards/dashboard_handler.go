package dashboards

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
)

type DashboardHandler interface {
	Find(c *gin.Context)
}

type dashboardHandler struct {
	dashboardService DashboardService
}

func NewDashboardHandler(dashboardService DashboardService) DashboardHandler {
	return &dashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *dashboardHandler) Find(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)

	dashboard, err := h.dashboardService.Find(projectID)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondOK(c, gin.H{
		"dashboard": dashboard,
	}, "Find Success.")
}

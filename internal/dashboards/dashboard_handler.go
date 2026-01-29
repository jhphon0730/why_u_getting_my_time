package dashboards

import "github.com/gin-gonic/gin"

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

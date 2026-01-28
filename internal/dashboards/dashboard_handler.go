package dashboards

type DashboardHandler interface {
}

type dashboardHandler struct {
	dashboardService DashboardService
}

func NewDashboardHandler(dashboardService DashboardService) DashboardHandler {
	return &dashboardHandler{
		dashboardService: dashboardService,
	}
}

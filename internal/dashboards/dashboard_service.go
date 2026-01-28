package dashboards

type DashboardService interface {
}

type dashboardService struct {
	dashboardRepo DashboardRepository
}

func NewDashboardService(repo DashboardRepository) DashboardService {
	return &dashboardService{
		dashboardRepo: repo,
	}
}

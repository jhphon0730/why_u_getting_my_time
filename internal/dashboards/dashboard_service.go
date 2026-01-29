package dashboards

type DashboardService interface {
	Find(projectID uint) (*Dashboard, error)
}

type dashboardService struct {
	dashboardRepo DashboardRepository
}

func NewDashboardService(repo DashboardRepository) DashboardService {
	return &dashboardService{
		dashboardRepo: repo,
	}
}

// 대시보드 페이지의 모든 정보를 하나의 Service 함수에서 제공하여 편리하게 접근할 수 있도록 하는 것이 목표
func (s *dashboardService) Find(projectID uint) (*Dashboard, error) {
	countTestCasesByStatus := s.dashboardRepo.CountTestCasesByStatus(projectID)
	countTestCasesByAssignee := s.dashboardRepo.CountTestCasesByAssignee(projectID)
	findOverdueTestCases := s.dashboardRepo.FindOverdueTestCases(projectID)

	return &Dashboard{
		CountTestCasesByStatus:   countTestCasesByStatus,
		CountTestCasesByAssignee: countTestCasesByAssignee,
		FindOverdueTestCases:     findOverdueTestCases,
	}, nil
}

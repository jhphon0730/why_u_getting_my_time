package dashboards

import (
	"time"

	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	WithTx(fn func(tx *gorm.DB) error) error

	// GetProjectSummary(projcetID uint)
	CountTestCasesByStatus(projectID uint) []CountTestCasesByStatus     // 테스트 케이스 상태별 현황
	CountTestCasesByAssignee(projectID uint) []CountTestCasesByAssignee // 담당자별 진행 상황
	FindOverdueTestCases(projectID uint) []FindOverdueTestCases         // 마감일이 지난 테스트 목록
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *dashboardRepository) CountTestCasesByStatus(projectID uint) []CountTestCasesByStatus {
	var cases []CountTestCasesByStatus
	r.db.Model(&model.TestCase{}).Where("project_id = ?", projectID).Select("status, count(*) as count").Group("status").Scan(&cases)
	return cases
}

func (r *dashboardRepository) CountTestCasesByAssignee(projectID uint) []CountTestCasesByAssignee {
	var cases []CountTestCasesByAssignee
	r.db.Model(&model.TestCase{}).Where("project_id = ?", projectID).Select("assignee_id, count(*) as count").Group("assignee_id").Scan(&cases)
	return cases
}

func (r *dashboardRepository) FindOverdueTestCases(projectID uint) []FindOverdueTestCases {
	var cases []FindOverdueTestCases
	r.db.Model(&model.TestCase{}).Where("project_id = ? AND due_date < ?", projectID, time.Now()).Select("id, title, due_date").Scan(&cases)
	return cases
}

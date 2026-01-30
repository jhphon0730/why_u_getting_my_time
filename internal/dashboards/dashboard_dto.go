package dashboards

import "time"

type CountTestCasesByStatus struct {
	StatusID uint  `json:"status_id"`
	Count    int64 `json:"count"`
}

type CountTestCasesByAssignee struct {
	CurrentAssigneeID uint  `json:"current_assignee_id"`
	Count             int64 `json:"count"`
}

type FindOverdueTestCases struct {
	ID      uint      `json:"id"`
	Title   string    `json:"title"`
	DueDate time.Time `json:"due_date"`
}

type Dashboard struct {
	CountTestCasesByStatus   []CountTestCasesByStatus   `json:"count_test_cases_by_status"`
	CountTestCasesByAssignee []CountTestCasesByAssignee `json:"count_test_cases_by_assignee"`
	FindOverdueTestCases     []FindOverdueTestCases     `json:"find_overdue_test_cases"`
}

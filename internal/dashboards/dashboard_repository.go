package dashboards

import "gorm.io/gorm"

type DashboardRepository interface {
	WithTx(fn func(tx *gorm.DB) error) error
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

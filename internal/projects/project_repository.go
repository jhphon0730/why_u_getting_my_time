package projects

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// ProjectRepository는 프로젝트를 관리하는 인터페이스입니다.
type ProjectRepository interface {
	DB() *gorm.DB
	WithTx(fn func(tx *gorm.DB) error) error

	Create(project *model.Project) error
	FindAllByUserID(userID uint) ([]*model.Project, error)

	CreateMember(member *model.ProjectMember) error
}

// projectRepository는 프로젝트를 관리하는 구조체입니다.
type projectRepository struct {
	db *gorm.DB
}

// NewProjectRepository는 새로운 프로젝트 리포지토리를 생성하는 함수입니다.
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

// DB는 프로젝트 리포지토리의 데이터베이스를 반환하는 함수입니다.
func (r *projectRepository) DB() *gorm.DB {
	return r.db
}

// WithTx는 트랜잭션을 사용하여 함수를 실행하는 함수입니다.
func (r *projectRepository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

// Create는 새로운 프로젝트를 생성하는 함수입니다.
func (r *projectRepository) Create(project *model.Project) error {
	return r.db.Create(project).Error
}

// FindAll는 특정 멤버의 모든 프로젝트를 조회하는 함수입니다.
func (r *projectRepository) FindAllByUserID(userID uint) ([]*model.Project, error) {
	var projects []*model.Project
	r.db.
		Joins("JOIN project_members ON project_members.project_id = projects.id").
		Where("project_members.user_id = ?", userID).
		Find(&projects)

	return projects, nil
}

// CreateMember는 새로운 프로젝트 멤버를 생성하는 함수입니다.
func (r *projectRepository) CreateMember(member *model.ProjectMember) error {
	return r.db.Create(member).Error
}

package projects

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// ProjectMemberRepository는 프로젝트 멤버 관련 데이터베이스 작업을 수행하는 인터페이스입니다.
type ProjectMemberRepository interface {
	Create(projectMember *model.ProjectMember) error
	FindByProjectID(projectID uint) ([]*model.ProjectMember, error)
	FindByProjectIDAndUserID(projectID, userID uint) (*model.ProjectMember, error)
	UpdateRoleToManager(projectID, userID uint) error
	UpdateRoleToMember(projectID, userID uint) error
	Delete(projectID, userID uint) error
	IsMember(projectID, userID uint) (bool, error)
	IsManager(projectID, userID uint) (bool, error)
	CountManagers(projectID uint) (int64, error)
}

// projectMemberRepository는 프로젝트 멤버 관련 데이터베이스 작업을 수행하는 구현체입니다.
type projectMemberRepository struct {
	db *gorm.DB
}

// NewProjectMemberRepository는 프로젝트 멤버 관련 데이터베이스 작업을 수행하는 인터페이스를 반환합니다.
func NewProjectMemberRepository(db *gorm.DB) ProjectMemberRepository {
	return &projectMemberRepository{db: db}
}

// Create는 프로젝트 멤버를 생성합니다.
func (r *projectMemberRepository) Create(projectMember *model.ProjectMember) error {
	return r.db.Create(projectMember).Error
}

// FindByProjectID는 프로젝트 ID를 기반으로 프로젝트 멤버 목록을 찾습니다.
func (r *projectMemberRepository) FindByProjectID(projectID uint) ([]*model.ProjectMember, error) {
	var projectMembers []*model.ProjectMember
	if err := r.db.Where("project_id = ?", projectID).Find(&projectMembers).Error; err != nil {
		return nil, err
	}
	return projectMembers, nil
}

// UpdateRoleToManager 함수는 프로젝트 멤버의 역할을 관리자로 업데이트합니다.
func (r *projectMemberRepository) UpdateRoleToManager(projectID, userID uint) error {
	return r.db.Model(&model.ProjectMember{}).Where("project_id = ? AND user_id = ?", projectID, userID).Update("project_role", model.RoleManager).Error
}

// UpdateRoleToMember 함수는 프로젝트 멤버의 역할을 멤버로 업데이트합니다.
func (r *projectMemberRepository) UpdateRoleToMember(projectID, userID uint) error {
	return r.db.Model(&model.ProjectMember{}).Where("project_id = ? AND user_id = ?", projectID, userID).Update("project_role", model.RoleMember).Error
}

// Delete는 프로젝트 멤버를 삭제합니다.
func (r *projectMemberRepository) Delete(projectID, userID uint) error {
	return r.db.Where("project_id = ? AND user_id = ?", projectID, userID).Delete(&model.ProjectMember{}).Error
}

// FindByProjectIDAndUserID는 프로젝트 ID와 사용자 ID를 기반으로 프로젝트 멤버를 찾습니다.
func (r *projectMemberRepository) FindByProjectIDAndUserID(projectID, userID uint) (*model.ProjectMember, error) {
	var projectMember model.ProjectMember
	if err := r.db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&projectMember).Error; err != nil {
		return nil, err
	}
	return &projectMember, nil
}

// IsMember는 프로젝트 ID와 사용자 ID를 기반으로 프로젝트 멤버가 존재하는지 확인합니다.
func (r *projectMemberRepository) IsMember(projectID, userID uint) (bool, error) {
	var projectMember model.ProjectMember
	if err := r.db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&projectMember).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}
	return true, nil
}

// IsManager는 프로젝트 ID와 사용자 ID를 기반으로 프로젝트 관리자가 존재하는지 확인합니다.
func (r *projectMemberRepository) IsManager(projectID, userID uint) (bool, error) {
	var projectMember model.ProjectMember
	if err := r.db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&projectMember).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	if projectMember.ProjectRole == model.RoleManager {
		return true, nil
	}
	return false, nil
}

// CountManagers 함수는 프로젝트에 해당하는 매니저의 수를 반환함
func (r *projectMemberRepository) CountManagers(projectID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&model.ProjectMember{}).Where("project_id = ? AND project_role = ?", projectID, model.RoleManager).Count(&count).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}

		return 0, err
	}
	return count, nil
}

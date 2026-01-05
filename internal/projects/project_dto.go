package projects

import (
	"time"

	"github.com/jhphon0730/action_manager/internal/model"
)

// CreateProjectRequest는 프로젝트 생성 요청을 구조체입니다.
type CreateProjectRequest struct {
	Name        string `json:"name"`        // required
	Description string `json:"description"` // required
}

// ToModel 함수는 프로젝트 생성 요청을 프로젝트 모델로 변환합니다.
func (r *CreateProjectRequest) ToModel() *model.Project {
	return &model.Project{
		Name:        r.Name,
		Description: r.Description,
	}
}

// CreateProjectMemberRequest는 프로젝트 멤버 생성 요청을 구조체입니다.
type CreateProjectMemberRequest struct {
	ProjectID uint `json:"project_id"`
	UserID    uint `json:"user_id"`
}

// ToModel 함수는 프로젝트 멤버 모델로 변환합니다.
func (r *CreateProjectMemberRequest) ToModel() *model.ProjectMember {
	return &model.ProjectMember{
		ProjectID:   r.ProjectID,
		UserID:      r.UserID,
		ProjectRole: model.RoleMember,
		JoinedAt:    time.Now(),
	}
}

// ProjectResponse 구조체는 프로젝트 정보를 나타냅니다.
type ProjectResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToModelResponse 함수는 프로젝트 모델을 ProjectResponse로 변환합니다.
func ToModelResponse(project *model.Project) *ProjectResponse {
	return &ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

// ToModelListResponse 함수는 model.Project 리스트를 ProjectResponse 리스트로 변환합니다.
func ToModelListResponse(projects []*model.Project) []*ProjectResponse {
	var responses []*ProjectResponse
	for _, project := range projects {
		responses = append(responses, ToModelResponse(project))
	}
	return responses
}

// ProjectMemberResponse 구조체는 프로젝트 멤버 정보를 나타냅니다.
type ProjectMemberResponse struct {
	ID          uint      `json:"id"`
	ProjectID   uint      `json:"project_id"`
	UserID      uint      `json:"user_id"`
	ProjectRole string    `json:"project_role"`
	JoinedAt    time.Time `json:"joined_at"`
}

// ToModelMemberResponse 함수는 model.ProjectMember를 ProjectMemberResponse로 변환합니다.
func ToModelMemberResponse(member *model.ProjectMember) *ProjectMemberResponse {
	return &ProjectMemberResponse{
		ID:          member.ID,
		ProjectID:   member.ProjectID,
		UserID:      member.UserID,
		ProjectRole: member.ProjectRole,
		JoinedAt:    member.JoinedAt,
	}
}

// ToModelMemberResponseList 함수는 model.ProjectMember 리스트를 ProjectMemberResponse 리스트로 변환합니다.
func ToModelMemberResponseList(members []*model.ProjectMember) []*ProjectMemberResponse {
	var responses []*ProjectMemberResponse
	for _, member := range members {
		responses = append(responses, ToModelMemberResponse(member))
	}
	return responses
}

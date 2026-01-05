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

type ProjectResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToModelResponse(project *model.Project) *ProjectResponse {
	return &ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func ToModelListResponse(projects []*model.Project) []*ProjectResponse {
	var responses []*ProjectResponse
	for _, project := range projects {
		responses = append(responses, ToModelResponse(project))
	}
	return responses
}

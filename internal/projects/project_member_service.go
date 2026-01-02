package projects

// ProjectMemberService는 프로젝트 멤버 관리를 위한 인터페이스입니다.
type ProjectMemberService interface {
}

// NewProjectMemberService는 새로운 ProjectMemberService 인스턴스를 생성합니다.
type projectMemberService struct {
	projectMemberRepo ProjectMemberRepository
}

// NewProjectMemberService는 새로운 ProjectMemberService 인스턴스를 생성합니다.
func NewProjectMemberService(projectMemberRepo ProjectMemberRepository) ProjectMemberService {
	return &projectMemberService{
		projectMemberRepo: projectMemberRepo,
	}
}

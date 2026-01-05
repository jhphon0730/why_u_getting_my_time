package projects

// ProjectMemberService는 프로젝트 멤버 관리를 위한 인터페이스입니다.
type ProjectMemberService interface {
	Create(req *CreateProjectMemberRequest) error
	IsMember(projectID, userID uint) (bool, error)
	IsManager(projectID, userID uint) (bool, error)
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

// Create 프로젝트 멤버를 추가합니다.
func (s *projectMemberService) Create(req *CreateProjectMemberRequest) error {
	if exists, _ := s.projectMemberRepo.IsMember(req.ProjectID, req.UserID); exists {
		return ErrAlreadyMember
	}

	project := req.ToModel()
	return s.projectMemberRepo.Create(project)
}

// IsMember는 사용자가 프로젝트에 멤버인지 확인합니다.
func (s *projectMemberService) IsMember(projectID, userID uint) (bool, error) {
	return s.projectMemberRepo.IsMember(projectID, userID)
}

// IsManager는 사용자가 프로젝트 관리자인지 확인합니다.
func (s *projectMemberService) IsManager(projectID, userID uint) (bool, error) {
	return s.projectMemberRepo.IsManager(projectID, userID)
}

package users

// UserService는 사용자 관련 서비스를 제공하는 인터페이스입니다.
type UserService interface {
	CreateUser(req CreateUserRequest) error
}

// UserService는 사용자 관련 서비스를 제공하는 구현체입니다.
type userService struct {
	userRepo UserRepository
}

// NewUserService는 새로운 UserService를 생성합니다.
func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

// CreateUser는 새로운 사용자를 생성합니다.
func (s *userService) CreateUser(req CreateUserRequest) error {
	user := req.ToUser()
	return s.userRepo.Create(user)
}

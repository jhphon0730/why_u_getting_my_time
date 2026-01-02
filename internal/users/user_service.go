package users

import "github.com/jhphon0730/action_manager/internal/model"

// UserService는 사용자 관련 서비스를 제공하는 인터페이스입니다.
type UserService interface {
	CreateUser(req SignUpRequest) error
	GetUserByEmail(req SignInRequest) (*model.User, error)
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
func (s *userService) CreateUser(req SignUpRequest) error {
	user := req.ToModel()
	return s.userRepo.Create(user)
}

// GetUserByEmail 함수는 이메일로 사용자를 조회합니다.
func (s *userService) GetUserByEmail(req SignInRequest) (*model.User, error) {
	return s.userRepo.FindByEmail(req.Email)
}

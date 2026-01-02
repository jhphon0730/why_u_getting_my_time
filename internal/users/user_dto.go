package users

import "github.com/jhphon0730/action_manager/internal/model"

// SignUpRequest 사용자를 생성하기 위한 요청
type SignUpRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ToModel는 CreateUserRequest를 User로 변환하는 메서드
func (r *SignUpRequest) ToModel() *model.User {
	return &model.User{
		Email: r.Email,
		Name:  r.Name,
	}
}

// SignInRequest 사용자를 로그인하기 위한 요청
type SignInRequest struct {
	Email string `json:"email"`
}

// UserResponse는 사용자 정보를 나타내는 응답
type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ToModelResponse는 User를 UserResponse로 변환하는 메서드
func ToModelResponse(user *model.User) *UserResponse {
	if user == nil {
		return nil
	}

	return &UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}

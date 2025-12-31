package users

import "github.com/jhphon0730/action_manager/internal/model"

// CreateUserRequest는 사용자를 생성하기 위한 요청
type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ToUser는 CreateUserRequest를 User로 변환하는 메서드
func (r *CreateUserRequest) ToUser() *model.User {
	return &model.User{
		Email: r.Email,
		Name:  r.Name,
	}
}

// LoginRequest는 사용자를 로그인하기 위한 요청
type LoginRequest struct {
	Email string `json:"email"`
}

// UserResponse는 사용자 정보를 나타내는 응답
type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ToUserResponse는 User를 UserResponse로 변환하는 메서드
func ToUserResponse(user *model.User) *UserResponse {
	if user == nil {
		return nil
	}

	return &UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}

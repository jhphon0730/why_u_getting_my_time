package server

import (
	"github.com/jhphon0730/action_manager/internal/database"
	"github.com/jhphon0730/action_manager/internal/users"
)

// RegisterRoutes는 초기에 설정되는 라우트들을 등록합니다.
func (s *server) RegisterRoutes() {
	db := database.GetDB().DB

	userRepo := users.NewUserRepository(db)
	userSer := users.NewUserService(userRepo)
	userHan := users.NewUserHandler(userSer)

	v1 := s.engine.Group("/api/v1")
	usersGroup := v1.Group("/users")
	{
		usersGroup.POST("", userHan.SignUp)
		usersGroup.POST("/login", userHan.SignIn)
	}

}

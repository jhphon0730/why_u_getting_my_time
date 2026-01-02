package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/database"
	"github.com/jhphon0730/action_manager/internal/middleware"
	"github.com/jhphon0730/action_manager/internal/projects"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/internal/users"
)

// RegisterRoutes는 초기에 설정되는 라우트들을 등록합니다.
func (s *server) RegisterRoutes() {
	db := database.GetDB().DB

	userRepo := users.NewUserRepository(db)
	userSer := users.NewUserService(userRepo)
	userHan := users.NewUserHandler(userSer)

	projectRepo := projects.NewProjectRepository(db)
	projectMemberRepo := projects.NewProjectMemberRepository(db)
	projectSer := projects.NewProjectService(projectRepo)
	projectMemberSer := projects.NewProjectMemberService(projectMemberRepo)
	projectHan := projects.NewProjectHandler(projectSer, projectMemberSer)

	v1 := s.engine.Group("/api/v1")

	/* USER & AUTH */
	usersGroup := v1.Group("/users")
	{
		usersGroup.POST("", userHan.SignUp)
		usersGroup.POST("/sign-in", userHan.SignIn)
		usersGroup.POST("/sign-out", middleware.AuthMiddleware(), userHan.SignOut)
	}

	/* PROJECT */
	projectGroup := v1.Group("/projects")
	projectGroup.Use(middleware.AuthMiddleware())
	{
		projectGroup.POST("", projectHan.CreateProject)
		projectGroup.GET("", projectHan.GetAllProjects)
	}

	/* PING TEST */
	testGroup := v1.Group("ping")
	testGroup.Use(middleware.AuthMiddleware())
	{
		testGroup.GET("", func(c *gin.Context) {
			response.RespondSuccess(c, 200, "pong")
		})
	}

}

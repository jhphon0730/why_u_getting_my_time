package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/database"
	"github.com/jhphon0730/action_manager/internal/middleware"
	authmw "github.com/jhphon0730/action_manager/internal/middleware"
	"github.com/jhphon0730/action_manager/internal/projects"
	projectmw "github.com/jhphon0730/action_manager/internal/projects/middleware"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/internal/testcases"
	testresults "github.com/jhphon0730/action_manager/internal/testresult"
	"github.com/jhphon0730/action_manager/internal/teststatus"
	"github.com/jhphon0730/action_manager/internal/users"
)

// RegisterRoutes는 초기에 설정되는 라우트들을 등록합니다.
func (s *server) RegisterRoutes() {
	db := database.GetDB().DB

	userRepo := users.NewUserRepository(db)
	userSer := users.NewUserService(userRepo)
	userHan := users.NewUserHandler(userSer)

	teststatusRepo := teststatus.NewTestStatusRepository(db)
	teststatusSer := teststatus.NewTestStatusService(teststatusRepo)

	projectRepo := projects.NewProjectRepository(db)
	projectMemberRepo := projects.NewProjectMemberRepository(db)
	projectSer := projects.NewProjectService(projectRepo, teststatusSer)
	projectMemberSer := projects.NewProjectMemberService(projectMemberRepo)
	projectHan := projects.NewProjectHandler(projectSer, projectMemberSer)
	projectMemberHan := projects.NewProjectMemberHandler(projectMemberSer)

	testStatusRepo := teststatus.NewTestStatusRepository(db)
	testStatusSer := teststatus.NewTestStatusService(testStatusRepo)
	testStatusHan := teststatus.NewTestStatusHandler(testStatusSer)

	testCaseRepo := testcases.NewTestCaseRepository(db)
	testCaseSer := testcases.NewTestCaseService(testCaseRepo, testStatusSer, projectMemberSer)
	testCaseHan := testcases.NewTestCaseHandler(testCaseSer)

	testResultRepo := testresults.NewTestResultRepository(db)
	testResultSer := testresults.NewTestResultService(testResultRepo, testCaseSer)
	testResultHan := testresults.NewTestResultHandler(testResultSer)

	v1 := s.engine.Group("/api/v1")

	/* USER & AUTH */
	usersGroup := v1.Group("/users")
	{
		usersGroup.POST("", userHan.SignUp)
		usersGroup.POST("/sign-in", userHan.SignIn)
		usersGroup.POST("/sign-out", authmw.AuthMiddleware(), userHan.SignOut)
	}

	/* PROJECT */
	projectGroup := v1.Group("/projects")
	projectGroup.Use(authmw.AuthMiddleware())
	{
		projectGroup.POST("", projectHan.CreateProject)
		projectGroup.GET("", projectHan.GetAllProjects)

		/* PROJECT MEMBER */
		projectMemberGroup := projectGroup.Group("/members/:projectID")
		{
			projectMemberGroup.POST("/:userID", projectmw.RequireProjectManager(projectMemberSer), projectMemberHan.AddMember)
			projectMemberGroup.DELETE("/:userID", projectmw.RequireProjectManager(projectMemberSer), projectMemberHan.DeleteMember)
			projectMemberGroup.GET("", projectmw.RequireProjectMember(projectMemberSer), projectMemberHan.FindByProjectID)
			projectMemberGroup.PATCH("/:userID/manager", projectmw.RequireProjectManager(projectMemberSer), projectMemberHan.UpdateRoleToManager)
			projectMemberGroup.PATCH("/:userID/member", projectmw.RequireProjectManager(projectMemberSer), projectMemberHan.UpdateRoleToMember)
		}

		/* TEST STATUS */
		testStatusGroup := projectGroup.Group("/test-status/:projectID")
		{
			testStatusGroup.GET("", projectmw.RequireProjectMember(projectMemberSer), testStatusHan.FindByProjectID)
			testStatusGroup.POST("", projectmw.RequireProjectManager(projectMemberSer), testStatusHan.Create)
			testStatusGroup.DELETE("/:statusID", projectmw.RequireProjectManager(projectMemberSer), testStatusHan.Delete)
		}

		/* TEST CASE */
		testCaseGroup := projectGroup.Group("/test-cases/:projectID")
		{
			testCaseGroup.POST("", projectmw.RequireProjectMember(projectMemberSer), testCaseHan.Create)
			testCaseGroup.GET("", projectmw.RequireProjectMember(projectMemberSer), testCaseHan.FindByProjectID)
			testCaseGroup.PATCH("/:testCaseID/status", projectmw.RequireProjectMember(projectMemberSer), testCaseHan.UpdateStatus)
		}

		/* TEST RESULT */
		testResultGroup := projectGroup.Group("/test-results/:projectID")
		{
			testResultGroup.POST("", projectmw.RequireProjectMember(projectMemberSer), testResultHan.Create)
		}

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

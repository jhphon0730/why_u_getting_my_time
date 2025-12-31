package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server는 사용자 요청을 처리하는 인터페이스입니다.
type Server interface {
	Start() error
	ShutDown(ctx context.Context) error

	RegisterRoutes()
}

// Server는 사용자 요청을 처리하는 구현체입니다.
type server struct {
	engine     *gin.Engine
	httpServer *http.Server
}

// NewServer는 새로운 서버 인스턴스를 생성합니다.
func NewServer(PORT string, APP_ENV string) Server {
	if APP_ENV == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin"},
		ExposeHeaders:    []string{"Content-Type", "Authorization", "Origin"},
		AllowCredentials: true, // 해당 서버에서 사용자 인증 정보를 허용합니다.
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}

	return &server{
		engine:     r,
		httpServer: srv,
	}
}

// Start는 서버를 시작합니다.
func (s *server) Start() error {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// ShutDown은 서버를 종료합니다.
func (s *server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

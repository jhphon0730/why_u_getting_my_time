package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/auth"
	"github.com/jhphon0730/action_manager/internal/response"
)

// UserHandler는 사용자 관련 요청을 처리하는 인터페이스입니다.
type UserHandler interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
}

// userHandler는 사용자 관련 요청을 처리하는 구조체입니다.
type userHandler struct {
	userSer UserService
}

// NewUserHandler는 UserHandler를 생성하는 함수입니다.
func NewUserHandler(userSer UserService) UserHandler {
	return &userHandler{
		userSer: userSer,
	}
}

// SignUp는 사용자를 생성하는 함수입니다.
func (h *userHandler) SignUp(c *gin.Context) {
	var signUpReq SignUpRequest
	if err := c.ShouldBindJSON(&signUpReq); err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}
	err := h.userSer.CreateUser(signUpReq)
	if err != nil {
		if err == ErrDuplicateEmail {
			response.RespondError(c, http.StatusBadRequest, ErrDuplicateEmail.Error())
		} else {
			response.RespondError(c, http.StatusInternalServerError, ErrInternal.Error())
		}
		return
	}

	response.RespondCreated(c, nil)
}

// SignIn 함수는 사용자를 찾아 로그인합니다.
func (h *userHandler) SignIn(c *gin.Context) {
	var signInReq SignInRequest
	if err := c.ShouldBindJSON(&signInReq); err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}

	user, err := h.userSer.GetUserByEmail(signInReq)
	if err != nil {
		if err == ErrNotFound {
			response.RespondError(c, http.StatusNotFound, ErrNotFound.Error())
		} else {
			response.RespondError(c, http.StatusInternalServerError, ErrInternal.Error())
		}
		return
	}

	if user == nil {
		response.RespondError(c, http.StatusNotFound, ErrNotFound.Error())
		return
	}

	token, err := auth.GenerateJWTToken(user.ID)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, ErrInternal.Error())
		return
	}

	c.SetCookie("token", token, 43200, "/", "", false, true)

	response.RespondOK(c, gin.H{
		"token": token,
		"user":  ToModelResponse(user),
	})
}

// SignOut 함수는 사용자를 로그아웃합니다.
func (h *userHandler) SignOut(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil || token == "" {
		response.RespondError(c, http.StatusUnauthorized, ErrUnauthorized.Error())
		return
	}

	c.SetCookie("token", "", -1, "/", "", false, true)

	response.RespondOK(c, gin.H{
		"message": "Logged out successfully",
	})
}

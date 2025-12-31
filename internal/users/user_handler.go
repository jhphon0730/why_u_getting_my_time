package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
)

// UserHandler는 사용자 관련 요청을 처리하는 인터페이스입니다.
type UserHandler interface {
	SignUp(c *gin.Context)
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
	var createUserRequest CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}
	err := h.userSer.CreateUser(createUserRequest)
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

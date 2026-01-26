package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// successResponse는 성공 응답 구조체
type successResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// errorResponse는 실패 응답 구조체
type errorResponse struct {
	Message string `json:"message"`
}

// RespondSuccess는 성공 응답을 반환
func RespondSuccess(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, successResponse{
		Data:    data,
		Message: message,
	})
}

// RespondError는 실패 응답을 반환
func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, errorResponse{
		Message: message,
	})
}

// RespondOK는 200 OK 응답
func RespondOK(c *gin.Context, data interface{}, message string) {
	RespondSuccess(c, http.StatusOK, data, message)
}

// RespondCreated는 201 Created 응답
func RespondCreated(c *gin.Context, data interface{}, message string) {
	RespondSuccess(c, http.StatusCreated, data, message)
}

// RespondNoContent는 204 No Content 응답
func RespondNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func RespondFile(c *gin.Context, filePath string, downloadName string) {
	c.FileAttachment(filePath, downloadName)
}

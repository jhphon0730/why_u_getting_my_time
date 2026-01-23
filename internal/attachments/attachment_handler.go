package attachments

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/response"
	"github.com/jhphon0730/action_manager/pkg/contextutils"
	"github.com/jhphon0730/action_manager/pkg/utils"
)

type AttachmentHandler interface {
	Create(c *gin.Context)
	FindOne(c *gin.Context)
	Find(c *gin.Context)
	Download(c *gin.Context)
}

type attachmentHandler struct {
	attachmentService AttachmentService
}

func NewAttachmentHandler(attachmentService AttachmentService) AttachmentHandler {
	return &attachmentHandler{
		attachmentService: attachmentService,
	}
}

func (h *attachmentHandler) Create(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)
	uploadedBy, _ := contextutils.GetUserID(c)

	// form 값
	targetType := c.PostForm("target_type")
	targetIDStr := c.PostForm("target_id")

	targetID := utils.InterfaceToUint(targetIDStr)

	req := CreateAttachmentRequest{
		TargetType: targetType,
		TargetID:   uint(targetID),
	}

	// files
	form, err := c.MultipartForm()
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, ErrInvalidRequest.Error())
		return
	}

	fileHeaders := form.File["files"]
	if len(fileHeaders) == 0 {
		response.RespondError(c, http.StatusBadRequest, ErrNoFilesProvided.Error())
		return
	}

	var files []UploadedFile
	for _, fh := range fileHeaders {
		f, err := fh.Open()
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, err.Error())
			return
		}
		defer f.Close()

		bytes, err := io.ReadAll(f)
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, err.Error())
			return
		}

		files = append(files, UploadedFile{
			Bytes:       bytes,
			Filename:    fh.Filename,
			ContentType: fh.Header.Get("Content-Type"),
			Size:        fh.Size,
		})
	}

	if err := h.attachmentService.Create(req, projectID, uploadedBy, files); err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondSuccess(c, http.StatusCreated, nil)
}

func (h *attachmentHandler) FindOne(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)
	attachmentID, _ := contextutils.GetAttachmentIDByParam(c)

	targetType := contextutils.GetQueryValue(c, "target_type")
	targetIDStr := contextutils.GetQueryValue(c, "target_id")
	targetID := utils.InterfaceToUint(targetIDStr)

	attachment, err := h.attachmentService.FindOne(targetType, projectID, targetID, attachmentID)
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondSuccess(c, http.StatusOK, gin.H{
		"message":    "Attachment found successfully",
		"attachment": ToModelAttachmentResponse(attachment),
	})
}

func (h *attachmentHandler) Find(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)

	targetType := contextutils.GetQueryValue(c, "target_type")
	targetIDStr := contextutils.GetQueryValue(c, "target_id")
	targetID := utils.InterfaceToUint(targetIDStr)

	attachments, err := h.attachmentService.Find(targetType, projectID, targetID)
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondSuccess(c, http.StatusOK, gin.H{
		"message":     "Attachments found successfully",
		"attachments": ToModelAttachmentResponseList(attachments),
	})
}

func (h *attachmentHandler) Download(c *gin.Context) {
	projectID, _ := contextutils.GetProjectIDByParam(c)
	attachmentID, _ := contextutils.GetAttachmentIDByParam(c)

	targetType := contextutils.GetQueryValue(c, "target_type")
	targetIDStr := contextutils.GetQueryValue(c, "target_id")
	targetID := utils.InterfaceToUint(targetIDStr)

	attachment, err := h.attachmentService.FindOne(targetType, projectID, targetID, attachmentID)
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	response.RespondFile(c, attachment.FilePath, attachment.Filename)
}

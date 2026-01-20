package attachments

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

type AttachmentRepository interface {
	WithTx(fn func(tx *gorm.DB) error) error

	Create(attachment *model.Attachment) error
}

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{
		db: db,
	}
}

func (r *attachmentRepository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *attachmentRepository) Create(attachment *model.Attachment) error {
	return r.db.Create(attachment).Error
}

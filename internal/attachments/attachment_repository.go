package attachments

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

type AttachmentRepository interface {
	WithTx(fn func(tx *gorm.DB) error) error

	Create(attachment *model.Attachment) error
	FindOne(targetType string, targetID, attachmentID uint) (*model.Attachment, error)
	Find(targetType string, targetID uint) ([]*model.Attachment, error)
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

func (r *attachmentRepository) FindOne(targetType string, targetID, attachmentID uint) (*model.Attachment, error) {
	var attachment model.Attachment
	if err := r.db.Where("target_type = ? AND target_id = ? AND id = ?", targetType, targetID, attachmentID).First(&attachment).Error; err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (r *attachmentRepository) Find(targetType string, targetID uint) ([]*model.Attachment, error) {
	var attachments []*model.Attachment
	if err := r.db.Where("target_type = ? AND target_id = ?", targetType, targetID).Find(&attachments).Error; err != nil {
		return nil, err
	}
	return attachments, nil
}

package attachments

import "github.com/jhphon0730/action_manager/internal/model"

// CreateAttachmentRequest 구조체는 새로운 첨부 파일을 생성하기 위한 요청 정보를 포함
type CreateAttachmentRequest struct {
	TargetType string `json:"target_type"`
	TargetID   uint   `json:"target_id"`
}

// Attachment 생성 시에 같이 오는 파일 메타 데이터 구조체
type UploadedFile struct {
	Bytes       []byte
	Filename    string
	ContentType string
	Size        int64
}

// ToModel 함수는 CreateAttachmentRequest를 model.Attachment 모델로 변환하는 함수
func (d *CreateAttachmentRequest) ToModel() *model.Attachment {
	return &model.Attachment{
		TargetType: d.TargetType,
		TargetID:   d.TargetID,
	}
}

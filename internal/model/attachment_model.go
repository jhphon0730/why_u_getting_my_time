package model

type Attachment struct {
	BaseModel

	TargetType string `gorm:"not null"` // TEST_CASE / TEST_RESULT
	TargetID   uint   `gorm:"not null;index"`

	FilePath   string `gorm:"not null"`
	UploadedBy uint   `gorm:"not null"`

	Uploader User `gorm:"foreignKey:UploadedBy"`
}

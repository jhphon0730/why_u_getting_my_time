package model

import "time"

type Project struct {
	BaseModel
	Name        string `gorm:"not null"`
	Description string
}

type ProjectMember struct {
	BaseModel

	ProjectID   uint   `gorm:"not null;index"`
	UserID      uint   `gorm:"not null;index"`
	ProjectRole string `gorm:"not null"` // MANAGER / MEMBER

	JoinedAt time.Time `gorm:"not null"`

	Project Project
	User    User
}

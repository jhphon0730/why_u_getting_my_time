package model

type User struct {
	BaseModel
	Email string `gorm:"uniqueIndex;not null"`
	Name  string `gorm:"not null"`
}

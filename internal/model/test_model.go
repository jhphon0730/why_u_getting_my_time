package model

import "time"

// 테스트 케이스 상태 모델
type TestStatus struct {
	BaseModel
	ProjectID uint   `gorm:"not null;index"`
	Name      string `gorm:"not null"`
	IsActive  bool   `gorm:"default:true"`

	Project Project
}

// 테스트 케이스 모델
type TestCase struct {
	BaseModel
	ProjectID   uint   `gorm:"not null;index"`
	Title       string `gorm:"not null"`
	Description string

	CurrentStatusID   uint `gorm:"not null"`
	CurrentAssigneeID *uint

	DueDate *time.Time

	Project         Project
	CurrentStatus   TestStatus `gorm:"foreignKey:CurrentStatusID"`
	CurrentAssignee *User      `gorm:"foreignKey:CurrentAssigneeID"`
}

// 테스트 케이스 실행 결과 기록
type TestResult struct {
	BaseModel
	TestCaseID uint   `gorm:"not null;index"`
	Result     string `gorm:"not null"` // SUCCESS / FAIL / ABORT
	Comment    string

	TestCase TestCase
}

// 테스트 케이스 상태 변경 이력
type TestStatusHistory struct {
	BaseModel

	TestCaseID   uint `gorm:"not null;index"`
	FromStatusID uint `gorm:"not null"`
	ToStatusID   uint `gorm:"not null"`
	ChangedByID  uint `gorm:"not null"`

	TestCase   TestCase
	FromStatus TestStatus `gorm:"foreignKey:FromStatusID"`
	ToStatus   TestStatus `gorm:"foreignKey:ToStatusID"`
	ChangedBy  User       `gorm:"foreignKey:ChangedByID"`
}

// 테스트 케이스 담당자 변경 이력
type TestAssigneeHistory struct {
	BaseModel

	TestCaseID     uint `gorm:"not null;index"`
	FromAssigneeID *uint
	ToAssigneeID   *uint
	ChangedByID    uint `gorm:"not null"`

	TestCase     TestCase
	FromAssignee *User `gorm:"foreignKey:FromAssigneeID"`
	ToAssignee   *User `gorm:"foreignKey:ToAssigneeID"`
	ChangedBy    User  `gorm:"foreignKey:ChangedByID"`
}

package database

import (
	"fmt"
	"sync"

	"github.com/jhphon0730/action_manager/internal/config"
	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var (
	db   *DB
	once sync.Once
)

// NewDB 함수는 새로운 DB 인스턴스를 생성하고 반환
func NewDB() (*DB, error) {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.POSTGRES.DB_HOST,
		cfg.POSTGRES.DB_PORT,
		cfg.POSTGRES.DB_USER,
		cfg.POSTGRES.DB_PASSWORD,
		cfg.POSTGRES.DB_NAME,
		cfg.POSTGRES.SSL_MODE,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{gormDB}, nil
}

// InitDatabase 함수는 DB를 초기화하고 마이그레이션을 수행
func InitDatabase() {
	db := GetDB()
	if err := AutoMigrate(db.DB); err != nil {
		panic(err.Error())
	}
}

// GetDB 함수는 DB 인스턴스를 반환
func GetDB() *DB {
	once.Do(func() {
		var err error
		if db, err = NewDB(); err != nil {
			panic(err.Error())
		}
	})

	return db
}

// AutoMigrate 함수는 DB 스키마를 마이그레이션
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Project{},

		&model.ProjectMember{},
		&model.TestStatus{},

		&model.TestCase{},
		&model.TestResult{},

		&model.TestStatusHistory{},
		&model.TestAssigneeHistory{},

		&model.Attachment{},
	)
}

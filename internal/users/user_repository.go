package users

import (
	"strings"

	"github.com/jhphon0730/action_manager/internal/model"
	"gorm.io/gorm"
)

// UserRepository는 사용자 관련 데이터를 관리하는 인터페이스입니다.
type UserRepository interface {
	Create(user *model.User) error
	FindByID(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

// userRepository는 UserRepository를 구현하는 구조체입니다.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository는 userRepository를 생성하는 함수입니다.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create는 새로운 사용자를 생성하는 함수입니다.
func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return ErrDuplicateEmail
		}

		return err
	}

	return nil
}

// FindByID는 ID로 사용자를 찾는 함수입니다.
func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

// FindByEmail는 이메일로 사용자를 찾는 함수입니다.
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

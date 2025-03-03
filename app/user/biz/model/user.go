package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}
func Create(db *gorm.DB, user *User) error {
	tx := db.Begin() // 开启事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 发生异常时回滚
		}
	}()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback() // 创建失败时回滚
		return fmt.Errorf("failed to create user: %w", err)
	}

	return tx.Commit().Error // 提交事务
}
func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

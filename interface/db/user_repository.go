package db

import (
	"github.com/haru0017/go-clean-architecture/domain/model"
)

// UserRepository は構造体ではなくinterfaceのSQLHandler
type UserRepository struct {
	SQLHandler
}

// FindByID はidでuserを取得するメソッド
func (userRepository *UserRepository) FindByID(id int) (user model.User, err error) {
	if err = userRepository.Find(&user, id).Error; err != nil {
		return
	}
	return
}

// Store は新規userを作成するメソッド
func (userRepository *UserRepository) Store(u model.User) (user model.User, err error) {
	if err = userRepository.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// Update はuserを更新するメソッド
func (userRepository *UserRepository) Update(u model.User) (user model.User, err error) {
	if err = userRepository.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// DeleteByID はidでuserを削除するメソッド
func (userRepository *UserRepository) DeleteByID(u model.User) (err error) {
	if err = userRepository.Delete(&u).Error; err != nil {
		return
	}
	return
}

// FindAll は全userを取得するメソッド
func (userRepository *UserRepository) FindAll() (users model.Users, err error) {
	if err = userRepository.Find(&users).Error; err != nil {
		return
	}
	return
}

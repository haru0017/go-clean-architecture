package usecase

import (
	"github.com/haru0017/go-clean-architecture/domain/model"
)

// UserInteractor はinterfaceのUserRepository
type UserInteractor struct {
	UserRepository UserRepository
}

// UserByID はidでuserを取得するメソッド
func (interactor *UserInteractor) UserByID(id int) (user model.User, err error) {
	user, err = interactor.UserRepository.FindByID(id)
	return
}

// Users は全userを取得するメソッド
func (interactor *UserInteractor) Users() (users model.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

// Add は新規userを作成するメソッド
func (interactor *UserInteractor) Add(u model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

// Update はuserを更新するメソッド
func (interactor *UserInteractor) Update(u model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

// DeleteByID はidでuserを削除するメソッド
func (interactor *UserInteractor) DeleteByID(user model.User) (err error) {
	err = interactor.UserRepository.DeleteByID(user)
	return
}
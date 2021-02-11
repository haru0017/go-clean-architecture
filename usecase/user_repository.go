package usecase

import (
	"github.com/haru0017/go-clean-architecture/domain/model"
)

// UserRepository はinterfaceのメソッドをusecase層で使うためのinterface
type UserRepository interface {
	FindByID(int) (model.User, error)
	Store(model.User) (model.User, error)
	Update(model.User) (model.User, error)
	DeleteByID(model.User) error
	FindAll() (model.Users, error)
}

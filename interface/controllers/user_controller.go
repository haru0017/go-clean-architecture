package controllers

import (
	"strconv"
	"github.com/haru0017/go-clean-architecture/domain/model"
	"github.com/haru0017/go-clean-architecture/interface/db"
	"github.com/haru0017/go-clean-architecture/usecase"
)

// UserController はusecaseの構造体UserInteractorを持つ
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewController はcontrollerを生成する
func NewController(SQLHandler db.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &db.UserRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// CreateUser は新規userを作成するcontrollerのmethod
func (controller *UserController) CreateUser(c Context) (err error) {
	u := model.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// GetUser はidによりuserを取得するcontrollerのmethod
func (controller *UserController) GetUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserByID(id)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// GetUsers は全userを取得するcontrollerのmethod
func (controller *UserController) GetUsers(c Context) (err error) {
	users, err := controller.Interactor.Users()

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// UpdateUser はuserを更新するcontrollerのmethod
func (controller *UserController) UpdateUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := model.User{ID: id}
	c.Bind(&u)
	user, err := controller.Interactor.Update(u)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// DeleteUser はuserを削除するcontrollerのmethod
func (controller *UserController) DeleteUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := model.User{ID: id}
	err = controller.Interactor.DeleteByID(user)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
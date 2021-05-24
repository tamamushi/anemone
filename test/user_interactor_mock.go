/* vim: set ts=4 sw=4: */

package test

import (
	//"anemone/adapter/controllers"
	"anemone/application/usecase"
	//"anemone/adapter/helper"
	//"anemone/codes"
)

// User Create UseCase Mocking Class
type userUseCaseMock struct {
	usecase.IUserUseCase
	MockCreate func()
}

func NewUserUseCaseMock() *userUseCaseMock {
	return &userUseCaseMock{}
}

func (u *userUseCaseMock) Create() {
	u.MockCreate()
}

type UserUseCaseMethod userUseCaseMethod

type userUseCaseMethod struct {
	Create func()
	Remove func(id string) error
	Update func()
	//	FindById func(id string) (*model.User, error)
	Finds func()
}

func GetUserUseCaseMethodStruct() *userUseCaseMethod {
	return &userUseCaseMethod{}
}

func (m *userUseCaseMethod) SetCreate(f func()) *userUseCaseMethod {
	m.Create = f
	return m
}

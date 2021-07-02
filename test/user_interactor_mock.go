/* vim: set ts=4 sw=4: */

package test

import (
	"anemone/application/usecase"
)

// User Create UseCase Mocking Class
type userUseCaseMock struct {
	usecase.IUserUseCase
	MockCreate   func() error
	MockRemove   func(id string) error
	MockUpdate   func() error
	MockFindById func(id string) error
	//MockFindById func(id string) (*model.User, error)
	MockFinds func() error
}

func NewUserUseCaseMock() *userUseCaseMock {
	return &userUseCaseMock{}
}
func (u *userUseCaseMock) Create() error {
	return u.MockCreate()
}
func (u *userUseCaseMock) Update() error {
	return u.MockUpdate()
}
func (u *userUseCaseMock) Remove(id string) error {
	return u.MockRemove(id)
}
func (u *userUseCaseMock) FindById(id string) error {
	return u.MockFindById(id)
}
func (u *userUseCaseMock) Finds() error {
	return u.MockFinds()
}

type UserUseCaseMethod struct {
	Create func() error
	Remove func(id string) error
	Update func() error
	//	FindById func(id string) (*model.User, error)
	FindById func(id string) error
	Finds    func() error
}

func GetUserUseCaseMethodStruct() *UserUseCaseMethod {
	return &UserUseCaseMethod{}
}
func (m *UserUseCaseMethod) SetCreate(f func() error) *UserUseCaseMethod {
	m.Create = f
	return m
}
func (m *UserUseCaseMethod) SetRemove(f func(id string) error) *UserUseCaseMethod {
	m.Remove = f
	return m
}
func (m *UserUseCaseMethod) SetUpdate(f func() error) *UserUseCaseMethod {
	m.Update = f
	return m
}
func (m *UserUseCaseMethod) SetFindById(f func(id string) error) *UserUseCaseMethod {
	m.FindById = f
	return m
}
func (m *UserUseCaseMethod) SetFinds(f func() error) *UserUseCaseMethod {
	m.Finds = f
	return m
}

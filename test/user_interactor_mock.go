/* vim: set ts=4 sw=4: */

package test

import (
	"anemone/application/usecase"
	"anemone/model"
)

// User Create UseCase Mocking Class
type userUseCaseMock struct {
	usecase.IUserUseCase
	MockCreate   func(interface{}) (*model.User, error)
	MockRemove   func(id string) error
	MockUpdate   func(interface{}) (*model.User, error)
	MockFindById func(id string) (*model.User, error)
	MockFinds    func() ([]*model.User, error)
}

func NewUserUseCaseMock() *userUseCaseMock {
	return &userUseCaseMock{}
}
func (u *userUseCaseMock) Create(i interface{}) (*model.User, error) {
	return u.MockCreate(i)
}
func (u *userUseCaseMock) Update(i interface{}) (*model.User, error) {
	return u.MockUpdate(i)
}
func (u *userUseCaseMock) Remove(id string) error {
	return u.MockRemove(id)
}
func (u *userUseCaseMock) FindById(id string) (*model.User, error) {
	return u.MockFindById(id)
}
func (u *userUseCaseMock) Finds() ([]*model.User, error) {
	return u.MockFinds()
}

type UserUseCaseMethod struct {
	Create   func(interface{}) (*model.User, error)
	Remove   func(id string) error
	Update   func(interface{}) (*model.User, error)
	FindById func(id string) (*model.User, error)
	Finds    func() ([]*model.User, error)
}

func GetUserUseCaseMethodStruct() *UserUseCaseMethod {
	return &UserUseCaseMethod{nil, nil, nil, nil, nil}
}
func (m *UserUseCaseMethod) SetCreate(
	f func(interface{}) (*model.User, error)) *UserUseCaseMethod {
	m.Create = f
	return m
}
func (m *UserUseCaseMethod) SetRemove(
	f func(id string) error) *UserUseCaseMethod {
	m.Remove = f
	return m
}
func (m *UserUseCaseMethod) SetUpdate(
	f func(interface{}) (*model.User, error)) *UserUseCaseMethod {
	m.Update = f
	return m
}
func (m *UserUseCaseMethod) SetFindById(
	f func(id string) (*model.User, error)) *UserUseCaseMethod {
	m.FindById = f
	return m
}
func (m *UserUseCaseMethod) SetFinds(
	f func() ([]*model.User, error)) *UserUseCaseMethod {
	m.Finds = f
	return m
}

func PrepareUseCaseMock(tt *TCase) usecase.IUserUseCase {
	usecase := NewUserUseCaseMock()
	method, _ := tt.GetMethod()
	inter, ok := method.(*UserUseCaseMethod)
	if ok {
		usecase.MockFindById = inter.FindById
		usecase.MockCreate = inter.Create
		usecase.MockRemove = inter.Remove
		usecase.MockUpdate = inter.Update
	} else {
		f1 := func(id string) (*model.User, error) { return new(model.User), nil }
		usecase.MockFindById = f1

		f2 := func(_ interface{}) (*model.User, error) { return nil, nil }
		usecase.MockCreate = f2

		f3 := func(id string) error { return nil }
		usecase.MockRemove = f3

		f4 := func(interface{}) (*model.User, error) { return nil, nil }
		usecase.MockUpdate = f4
	}
	return usecase
}

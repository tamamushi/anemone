/* vim: set ts=4 sw=4: */

package usecase

import (
	"anemone/application/interactor"
	"anemone/model"
)

func NewUserInteractor() IUserUseCase {
	return interactor.NewUserInteractor()
}

type IUserUseCase interface {
	Create(interface{}) (*model.User, error)
	Remove(id string) error
	Update(interface{}) (*model.User, error)
	FindById(id string) (*model.User, error)
	Finds() ([]*model.User, error)
}

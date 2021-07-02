/* vim: set ts=4 sw=4: */

package usecase

import (
	"anemone/application/interactor"
)

func NewUserInteractor() IUserUseCase {
	return interactor.NewUserInteractor()
}

type IUserUseCase interface {
	Create() error
	Remove(id string) error
	Update() error
	FindById(id string) error
	Finds() error
}

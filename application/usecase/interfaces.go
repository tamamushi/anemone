/* vim: set ts=4 sw=4: */

package usecase

type IUserUseCase interface {
	Create() error
	Remove(id string) error
	Update() error
	FindById(id string) error
	Finds() error
}

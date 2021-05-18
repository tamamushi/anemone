/* vim: set ts=4 sw=4: */

package usecase

type UserInteractor interface {
}

type userInteractor struct {
}

func NewUserInteractor() UserInteractor {
	return &userInteractor{}
}

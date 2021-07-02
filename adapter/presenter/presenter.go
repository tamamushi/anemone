/* vim: set ts=4 sw=4: */

package presenter

import (
	"github.com/spf13/cobra"
)

type Presenter struct {
	interactor usecase.IUserUseCase
	response   gateway.IResponse
}

func NewUserController(u usecase.IUserUseCase, r gateway.IResponse) UserController {
	return &userController{u, r}
}

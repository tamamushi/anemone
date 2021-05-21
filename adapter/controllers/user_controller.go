/* vim: set ts=4 sw=4: */

package controllers

import (
	"fmt"
	"github.com/spf13/cobra"

	"anemone/application"
	"anemone/application/usecase"
)

type UserController interface {
	Handler() *cobra.Command
}

type userController struct {
	interactor usecase.IUserUseCase
}

func newUserController(u usecase.IUserUseCase) UserController {
	return &userController{u}
}

func (s *userController) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("user called")
			fmt.Println("%v", args)
		},
	}
	return cmd
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	controller := newUserController(usecase)
	blder.AddCommand(controller.Handler())
}

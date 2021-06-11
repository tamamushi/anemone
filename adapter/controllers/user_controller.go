/* vim: set ts=4 sw=4: */

package controllers

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
	_ "anemone/adapter/handler"
	"anemone/application"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// UserController のinterface定義
type UserController interface {
	Handler() *cobra.Command
}

type userController struct {
	interactor usecase.IUserUseCase
	response   gateway.IResponse
}

func NewUserController(u usecase.IUserUseCase, r gateway.IResponse) UserController {
	return &userController{u, r}
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	response := gateway.NewResponse()
	controller := NewUserController(usecase, response)
	blder.AddCommand(controller.Handler())
}

func (s *userController) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "A brief description of your command",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				// Error コードを返す
				return errors.New(
					codes.NotEnoughArgument,
					"Required target sub command",
				)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New(
				codes.UnSupportedMethod,
				errors.Messagef("UnSupported called method: %s", args[0]),
			)
		},
	}
	return cmd
}

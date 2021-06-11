/* vim: set ts=4 sw=4: */

package handler

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
	"anemone/application"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// UserController のinterface定義
type UserCreateHandler interface {
	Handler() *cobra.Command
	Create() error
}

type userCreateHandler struct {
	interactor usecase.IUserUseCase
	response   gateway.IResponse
}

func NewUserCreateHandler(u usecase.IUserUseCase, r gateway.IResponse) UserController {
	return &userCreateHandler{u, r}
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	response := gateway.NewResponse()
	controller := NewUserController(usecase, response)
	blder.AddCommand(controller.Handler())
}

func (s *userCreateHandler) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create",
		Short: "User Create Command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.Create()
		},
	}
	cmd.Flags().String("data", "", "Your name")
	return cmd
}

func (s *userController) Create() error {
	s.interactor.Create()
	return nil
}

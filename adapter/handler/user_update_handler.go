/* vim: set ts=4 sw=4: */

package handler

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// UserUpdateHandler のinterface定義
type UserUpdateHandler interface {
	Handler() *cobra.Command
	Update() error
}

type handler struct {
	interactor usecase.IUserUseCase
	response   gateway.IResponse
}

func NewUserUpdateHandler(u usecase.IUserUseCase, r gateway.IResponse) UserUpdateHandler {
	return &handler{u, r}
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	response := gateway.NewResponse()
	controller := NewUserUpdateHandler(usecase, response)
	blder.AddCommand(controller.Handler())
}

func (s *handler) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "update",
		Short: "User update command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.Update()
		},
	}
	cmd.Flags().String("data", "", "Your name")
	return cmd
}

func (s *handler) Update() error {
	return nil
}

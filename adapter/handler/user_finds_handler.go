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

// UserFindsHandler のinterface定義
type UserFindsHandler interface {
	Handler() *cobra.Command
	Finds() error
}

type handler struct {
	interactor usecase.IUserUseCase
	response   gateway.IResponse
}

func NewUserFindsHandler(u usecase.IUserUseCase, r gateway.IResponse) UserFindsHandler {
	return &handler{u, r}
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	response := gateway.NewResponse()
	controller := NewUserFindsHandler(usecase, response)
	blder.AddCommand(controller.Handler())
}

func (s *handler) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "finds",
		Short: "User finds command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.Finds()
		},
	}
	return cmd
}

func (s *handler) Finds() error {
	return nil
}

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

// UserFindByIdHandler のinterface定義
type UserFindByIdHandler interface {
	Handler() *cobra.Command
	FindById(id string) error
}

type handler struct {
	interactor usecase.IUserUseCase
	response   gateway.IResponse
}

func NewUserFindByIdHandler(u usecase.IUserUseCase, r gateway.IResponse) UserFindByIdHandler {
	return &handler{u, r}
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	response := gateway.NewResponse()
	controller := NewUserFindByIdHandler(usecase, response)
	blder.AddCommand(controller.Handler())
}

func (s *handler) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "findbyid",
		Short: "User findbyid command",
		Args: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetString("id")
			if len(id) < 1 {
				return errors.New(
					codes.NotEnoughArgument,
					"Need to specified \x1b[3mID\x1b[0m",
				)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetString("id")
			err := s.FindById(id)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("id", "", "Your name")
	return cmd
}

func (s *handler) FindById(id string) error {

	_, err := s.interactor.FindById(id)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}

/* vim: set ts=4 sw=4: */

package controllers

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
	"anemone/application"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// UserRemoveHanlder のinterface定義
type UserRemoveHanlder interface {
	Handler() *cobra.Command
	Remove(id string) error
}

type handler struct {
	interactor usecase.IUserUseCase
	response   gateway.IResponse
}

func NewUserRemoveHanlder(u usecase.IUserUseCase, r gateway.IResponse) UserRemoveHanlder {
	return &userRemoveHandler{u, r}
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	response := gateway.NewResponse()
	controller := NewUserRemoveHanlder(usecase, response)
	blder.AddCommand(controller.Handler())
}

func (s *handler) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "User Remove Command",
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
			err := s.Remove(id)
			if err != nil {
				return err
			}
			s.response.
				cmd.Printf("%s", id)
		},
	}
	cmd.Flags().String("id", "", "Your name")
	return cmd
}

func (s *handler) Remove(id string) error {

	err := s.interactor.Remove(id)
	if err != nil {
		return errors.Newf(
			codes.InternalServerError,
			"Internal Server Error: %s",
			err,
		)
	}
	return nil
}

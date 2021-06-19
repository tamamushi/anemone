/* vim: set ts=4 sw=4: */

package handler

import (
	"github.com/spf13/cobra"

	"anemone/adapter/gateway"
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
}

func NewUserRemoveHanlder(u usecase.IUserUseCase) UserRemoveHanlder {
	return &userRemoveHandler{u}
}

func init() {
	blder := application.GetBuilderInstance()
	usecase := usecase.NewUserInteractor()
	handler := NewUserRemoveHanlder(usecase)
	blder.AddCommand(handler.Handler())
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

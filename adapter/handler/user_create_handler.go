/* vim: set ts=4 sw=4: */

package handler

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"anemone/adapter/helper"
	"anemone/application/usecase"
	//"anemone/adapter/gateway"
	//"anemone/codes"
	//"anemone/errors"
)

type UserCreateHandler interface {
	Handle() *cobra.Command
	Create() error
}

type handler struct {
	interactor usecase.IUserUseCase
}

func NewUserCreateHandler(u usecase.IUserUseCase) UserCreateHandler {
	return &handler{u}
}

func init() {
	blder, err := helper.GetBuilderInstance("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	usecase := usecase.NewUserInteractor()
	handler := NewUserCreateHandler(usecase)
	blder.AddCommand(handler.Handle())
}

func (s *handler) Handle() *cobra.Command {

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

func (s *handler) Create() error {
	//s.interactor.Create()
	return nil
}

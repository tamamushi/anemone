/* vim: set ts=4 sw=4: */

package controllers

import (
	"fmt"

	"github.com/spf13/cobra"

	//_ "anemone/adapter/handler"
	"anemone/adapter/helper"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

// UserController のinterface定義
type UserController interface {
	Handler() *cobra.Command
}

type controller struct {
	interactor usecase.IUserUseCase
	blder      helper.Builder
}

func NewUserController(u usecase.IUserUseCase) UserController {
	return &controller{u, helper.GetBuilderInstance("user")}
}

func init() {
	blder := helper.GetBuilderInstance("root")

	fmt.Printf("%#v\n", blder)
	usecase := usecase.NewUserInteractor()
	controller := NewUserController(usecase)
	blder.AddCommand(controller.Handler())
}

func (s *controller) Handler() *cobra.Command {

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
	fmt.Printf("%#v\n", s.blder)
	//subs := s.blder.GetCommands()
	//cmd.AddCommand(subs...)
	return cmd
}

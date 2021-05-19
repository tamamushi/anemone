/* vim: set ts=4 sw=4: */

package controllers

import (
	"fmt"
	"github.com/spf13/cobra"
	"anemone/application/usecase"
)

type UserController interface {
	Handler() *cobra.Command
}

type userController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(u usecase.UserInteractor) UserController {
	return &UserController {u}
}

func (s *userController) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "test2",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test2 called")
		},
	}
	return cmd
}

func init() {
	NewUserController(
	rootCmd.AddCommand(test2Cmd)
}

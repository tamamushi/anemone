/* vim: set ts=4 sw=4: */

package controllers

import (
	"fmt"
	"github.com/spf13/cobra"

	"anemone/adapter/helper"
	"anemone/application/usecase"
)

type TestController interface {
	Handler() *cobra.Command
}

type testController struct {
	Interactor usecase.IUserUseCase
}

func newTestController(u usecase.IUserUseCase) TestController {
	return &testController{u}
}

func (s *testController) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "test1",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test1 called")
		},
	}
	return cmd
}

func init() {
	blder, _ := helper.GetBuilderInstance("root")

	usecase := usecase.NewUserInteractor()
	controller := newTestController(usecase)
	blder.AddCommand(controller.Handler())
}

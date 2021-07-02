/* vim: set ts=4 sw=4: */

package test

import (
	"anemone/adapter/controllers"
	"github.com/spf13/cobra"
)

type controllerMock struct {
	controllers.IUserController
}

func NewUserControllerMock() controllers.IUserController {
	return &controllerMock{}
}

func (c *controllerMock) Handler() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "User handle command group (mocking)",
	}

	return cmd
}

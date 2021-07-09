/* vim: set ts=4 sw=4: */

package test

import (
	"anemone/adapter/controllers"
	"anemone/adapter/gateway"
	"github.com/spf13/cobra"
)

type GatewayMock struct {
	gateway.IGateway
	// 本来はanemone/gatewayにある
	// テストの際の引数を満たす上で便宜上
	// controllerMockに書いてしまう方が早いため
	// ここで定義する
}

func (g *GatewayMock) SetResponse(i interface{}) {
	_ = i
}

type controllerMock struct {
	controllers.IUserController
}

func NewUserControllerMock() controllers.IUserController {
	return &controllerMock{}
}

func (c *controllerMock) Handler(g gateway.IGateway) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "User handle command group (mocking)",
	}
	_ = g
	return cmd
}

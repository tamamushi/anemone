/* vim: set ts=4 sw=4: */

package test

import (
	"bytes"

	"github.com/spf13/cobra"

	"anemone/adapter/controllers"
	"anemone/adapter/gateway"
)

// 本来はanemone/gatewayにあるテストの際の引数を満たす上で便宜上
// controllerMockに書いてしまう方が早いためここで定義する
type GatewayMock struct {
	gateway.IGateway
	buffer *bytes.Buffer
}

func (g *GatewayMock) SetResponse(i interface{}) {
	g.buffer.Write([]byte(i.(string)))
}

func (g *GatewayMock) SetOut(b *bytes.Buffer) {
	g.buffer = b
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

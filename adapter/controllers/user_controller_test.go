/* vim: set ts=4 sw=4: */

package controllers_test

import (
	"bytes"
	"fmt"
	"testing"

	"anemone/test"
	"github.com/spf13/cobra"

	"anemone/adapter/controllers"
	"anemone/codes"
)

func SetupUserControllerTest(
	b *bytes.Buffer,
	t *testing.T,
	tt *test.TCase,
) *cobra.Command {

	g := &test.GatewayMock{}
	g.SetOut(b)

	// ユースケースの準備
	interactor := test.NewUserUseCaseMock()

	// コントローラーの準備
	controller := controllers.NewUserController(interactor)

	// Rootcmdの構築と取得
	cmd, args := test.SetupRootCMD(tt)

	// Rootcmdへコマンドコントローラーを登録
	cmd.AddCommand(controller.Handler(g))
	cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
	return cmd
}

func TestUserControllerCalled_Handler(t *testing.T) {
	title := fmt.Sprintf("[Just called Behavior]")
	fmt.Printf("\n")
	cases := []*test.TCase{
		// パラメータが足りない場合はNot Enough Argument
		// Userがない場合、User XXXXがない場合どちらもエラー
		test.Case(
			title,
			"Test not enough argument.",
			test.SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			test.SetCommand(""),
			test.SetExpectCode(codes.NotEnoughArgument),
		),
		// User XXXXの「XXXX」がサポートされてないコマンドの場合、UnSupported Method
		// User unknown_methodはサポートされてないのでエラー
		test.Case(
			title,
			"Test unsupported Method.",
			test.SetExpectCodeMsg("Return code expected UnSupportedMethod."),
			test.SetCommand("unknown_method"),
			test.SetExpectCode(codes.UnSupportedMethod),
		),
	}
	test.TestRun(t, cases, SetupUserControllerTest, "UserController")
}

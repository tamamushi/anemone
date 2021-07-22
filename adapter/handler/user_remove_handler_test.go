// +build user

/* vim: set ts=4 sw=4: */

package handler_test

import (
	"bytes"
	"fmt"
	"testing"

	"anemone/test"
	"github.com/spf13/cobra"

	"anemone/adapter/handler"
	"anemone/codes"
	"anemone/errors"
)

func SetupRemoveHandlerTest(
	b *bytes.Buffer,
	t *testing.T,
	tt *test.TCase,
) *cobra.Command {

	g := &test.GatewayMock{}
	g.SetOut(b)

	// ユースケースの準備
	usecase := test.PrepareUseCaseMock(tt)

	// コントローラーの準備
	controller := test.NewUserControllerMock()

	// パーサーの準備
	format := test.PrepareParserMock(tt)
	// Parserは薄いのでここでは実体を使う
	parser := handler.NewUserRemoveParser(format)
	// 各ハンドラ毎にパーサーのラッパーを用意。これは、
	// 各ハンドラで必要なメソッドの違いを吸収し、同一メソッド名
	// でハンドラによる処理の違いを実現する為に存在する

	// ハンドラの準備
	userHandler := handler.NewUserRemoveHandler(parser)
	userHandler.SetGateway(g)

	handle := userHandler.GetHandle()
	for k, v := range handle.GetSetters() {
		switch k {
		case "UserUseCase":
			v(usecase)
		}
	}

	// Rootcmdの構築と取得
	cmd, args := test.SetupRootCMD(tt)

	// UserControllerにUserRemoveハンドラを登録
	controllerCmd := controller.Handler(g)
	controllerCmd.AddCommand(userHandler.Handle())

	// Rootcmdへコマンドコントローラーを登録
	cmd.AddCommand(controllerCmd)
	cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
	return cmd
}

// Test User Remove Handler
func TestUserRemoveHandlerCalled_Handle(t *testing.T) {
	title := fmt.Sprintf("[Validation Behavior]")
	fmt.Println()
	cases := []*test.TCase{
		// 引数がない場合はエラー
		test.Case(
			title,
			"Test not enough argument.",
			test.SetCommand("remove"),
			test.SetExpectCode(codes.NotEnoughArgument),
			test.SetExpectCodeMsg("Return code expected NotEnoughArgument."),
		),
		// 引数のidが指定されたフォーマットじゃない場合はエラー
		test.Case(
			title,
			"Test argument statement incorrect format.",
			test.SetCommand("remove"),
			test.SetArgument("id", "12"),
			test.SetParser(
				test.GetParserMethodStruct().
					SetTryParse(func(id string, _ interface{}) error {
						return errors.New(
							codes.InvalidArgument,
							"Mocking Dymmy invalid argument Error",
						)
					},
					),
			),
			test.SetExpectCode(codes.InvalidArgument),
			test.SetExpectCodeMsg("Return code expected InvalidArgument."),
		),
	}
	title = fmt.Sprintf("[Processing Behavior]")
	cases = append(cases, []*test.TCase{
		// 削除処理が正常終了
		test.Case(
			title,
			"Test normaly correct operation.",
			test.SetCommand("remove"),
			test.SetArgument("id", "1234-A78B-12ID-6789"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetRemove(func(id string) error {
						return nil
					},
					),
			),
			test.SetExpectedNil(),
			test.SetExpectedMsg("Expected operation Successfully"),
		),
		// 削除処理が異常終了
		test.Case(
			title,
			"Test return internal server error.",
			test.SetCommand("remove"),
			test.SetArgument("id", "1234-A78B-12ID-6789"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetRemove(func(id string) error {
						return fmt.Errorf("Mocking Dummy Error")
					},
					),
			),
			test.SetExpectCode(codes.InternalServerError),
			test.SetExpectCodeMsg("Return code expected InternalServerError."),
		),
	}...)
	test.TestRun(t, cases, SetupRemoveHandlerTest, "UserRemoveHandler")
}

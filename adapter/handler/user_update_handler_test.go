/* vim: set ts=4 sw=4: */

package handler_test

import (
	"fmt"
	"testing"

	"anemone/model"
	"anemone/test"
	"github.com/spf13/cobra"

	"anemone/adapter/handler"
	"anemone/codes"
)

func SetupUpdateHandlerTest(t *testing.T, tt *test.TCase) *cobra.Command {

	// ユースケースの準備
	usecase := test.NewUserUseCaseMock()
	method, _ := tt.GetMethod()
	inter1, ok1 := method.(*test.UserUseCaseMethod)
	if ok1 {
		usecase.MockUpdate = inter1.Update
	} else {
		f := func(interface{}) (*model.User, error) { return nil, nil }
		usecase.MockUpdate = f
	}

	// コントローラーの準備
	controller := test.NewUserControllerMock()

	// パーサーの準備
	parser := test.NewParserMock(model.User{})
	parser_method, _ := tt.GetParser()
	inter2, ok2 := parser_method.(*test.ParserMethod)

	parser.MockTryParseFormat = func(s string) error { return nil }
	parser.MockInput = func(s string) interface{} { return nil }
	if ok2 {
		if parser.MockTryParseFormat != nil {
			parser.MockTryParseFormat = inter2.TryParseFormat
		}
		if parser.MockInput != nil {
			parser.MockInput = inter2.Input
		}
	}

	// ハンドラの準備
	userHandler := handler.NewUserUpdateHandler(parser)
	userHandler.SetGateway(&test.GatewayMock{})

	handle := userHandler.GetHandle()
	for k, v := range handle.GetSetters() {
		switch k {
		case "UserUseCase":
			v(usecase)
		}
	}

	// Rootcmdの構築と取得
	cmd, args := test.SetupRootCMD(tt)

	// UserControllerにUserCreareハンドラを登録
	controllerCmd := controller.Handler(&test.GatewayMock{})
	controllerCmd.AddCommand(userHandler.Handle())

	// Rootcmdへコマンドコントローラーを登録
	cmd.AddCommand(controllerCmd)
	cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
	return cmd
}

// Test User Update Handler
func TestUserUpdateHandlerCalled_Handle(t *testing.T) {
	title := fmt.Sprintf("[Validation Behavior]")
	fmt.Println()
	cases := []*test.TCase{
		// 引数がない場合はエラー
		test.Case(
			title,
			"Test not enough argument.",
			test.SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			test.SetCommand("update"),
			test.SetExpectCode(codes.NotEnoughArgument),
		),
		// 引数のフォーマットが不正な場合はエラー
		test.Case(
			title,
			"Test argument statement incorrect format.",
			test.SetCommand("update"),
			test.SetArgument("data", "{Invalid JSON format}"),
			test.SetParser(
				test.GetParserMethodStruct().
					SetTryParseFormat(func(string) error {
						return fmt.Errorf("Mocking Dummy Error")
					},
					),
			),
			test.SetExpectCode(codes.InvalidArgument),
			test.SetExpectCodeMsg("Return code expected InvalidArgument."),
		),
	}
	title = fmt.Sprintf("[Processing Behavior]")
	cases = append(cases, []*test.TCase{
		// 更新処理が正常終了
		test.Case(
			title,
			"Test normaly correct operation.",
			test.SetCommand("update"),
			test.SetArgument("data", "hogehoge"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetUpdate(func(interface{}) (*model.User, error) {
						fmt.Fprintf(
							test.Buffer,
							"テストは通るが仕様が確定していない為本来はNG ",
						)
						return nil, nil
					},
					),
			),
			test.SetExpectedNil(),
			test.SetExpectedMsg("Expected operation Successfully."),
		),
		// 更新処理が異常終了
		test.Case(
			title,
			"Test return internal server error.",
			test.SetCommand("update"),
			test.SetArgument("data", "hogehoge"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetUpdate(func(interface{}) (*model.User, error) {
						return nil, fmt.Errorf("Mocking Dummy Error")
					},
					),
			),
			test.SetExpectCode(codes.InternalServerError),
			test.SetExpectCodeMsg("Return code expected InternalServerError."),
		),
	}...)
	test.TestRun(t, cases, SetupUpdateHandlerTest, "UserUpdateHandler")
}

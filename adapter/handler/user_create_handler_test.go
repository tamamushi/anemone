// +build -user

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

func SetupCreateHandlerTest(t *testing.T, tt *test.TCase) *cobra.Command {

	// ユースケースの準備
	usecase := test.NewUserUseCaseMock()
	usecase_method, _ := tt.GetMethod()
	inter1, ok1 := usecase_method.(*test.UserUseCaseMethod)
	if ok1 {
		usecase.MockCreate = inter1.Create
	} else {
		f := func(_ interface{}) (*model.User, error) { return nil, nil }
		usecase.MockCreate = f
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
	createHandler := handler.NewUserCreateHandler(parser)
	createHandler.SetGateway(&test.GatewayMock{})

	handle := createHandler.GetHandle()
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
	controllerCmd.AddCommand(createHandler.Handle())

	// Rootcmdへコマンドコントローラーを登録
	cmd.AddCommand(controllerCmd)
	cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
	return cmd
}

// Test User Create Handler
func TestUserCreateHandlerCalled_Handle(t *testing.T) {
	title := fmt.Sprintf("[Validation Behavior]")
	fmt.Println()
	cases := []*test.TCase{
		// 引数がない場合はエラー
		test.Case(
			title,
			"Test not enough argument.",
			test.SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			test.SetCommand("create"),
			test.SetExpectCode(codes.NotEnoughArgument),
		),
		// 引数のフォーマットが不正な場合はエラー
		test.Case(
			title,
			"Test argument statement incorrect format.",
			test.SetCommand("create"),
			test.SetArgument("data", "{Inavalid JSON format}"),
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
		// 作成処理が正常終了
		test.Case(
			title,
			"Test normaly correct operation.",
			test.SetCommand("create"),
			test.SetArgument("data", "hogehoge"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetCreate(func(interface{}) (*model.User, error) {
						fmt.Fprint(
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
		// 作成処理が異常終了
		test.Case(
			title,
			"Test return internal server error.",
			test.SetCommand("create"),
			test.SetArgument("data", "hogehoge"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetCreate(func(interface{}) (*model.User, error) {
						return nil, fmt.Errorf("Mocking Dummy Error")
					},
					),
			),
			test.SetExpectCode(codes.InternalServerError),
			test.SetExpectCodeMsg("Return code expected InternalServerError."),
		),
	}...)
	test.TestRun(t, cases, SetupCreateHandlerTest, "UserCreateHandler")
}

// +build user

/* vim: set ts=4 sw=4: */

package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"anemone/model"
	"anemone/test"
	"github.com/spf13/cobra"

	"anemone/adapter/handler"
	"anemone/codes"
)

func SetupCreateHandlerTest(
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
	input := test.PrepareInputMock(tt)
	output := test.PrepareOutputMock(tt)
	parser := handler.NewUserCreateParser(format, input, output)

	// ハンドラの準備
	createHandler := handler.NewUserCreateHandler(parser)
	createHandler.SetGateway(g)

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
					SetTryParse(func(string, interface{}) error {
						return fmt.Errorf("Mocking Dummy Error")
					},
					),
			),
			test.SetExpectCode(codes.InvalidArgument),
			test.SetExpectCodeMsg("Return code expected InvalidArgument."),
		),
	}
	title = fmt.Sprintf("[Processing Behavior]")
	result, _ := json.Marshal(user)
	parser := test.GetParserMethodStruct()
	cases = append(cases, []*test.TCase{
		// 作成処理が正常終了
		test.Case(
			title,
			"Test normaly correct operation.",
			test.SetCommand("create"),
			test.SetArgument("data", "{hogehoge}"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetCreate(func(interface{}) (*model.User, error) {
						return user, nil
					},
					),
			),
			test.SetParser(
				parser.SetInput(func(s string, m interface{}) interface{} {
					_ = json.Unmarshal([]byte(s), m)
					return m
				},
				),
			),
			test.SetParser(
				parser.SetOutput(func(m interface{}) string {
					model, _ := json.Marshal(m)
					return string(model)
				},
				),
			),
			test.SetExpected(string(result)),
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

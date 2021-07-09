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

func SetupFindsHandlerTest(t *testing.T, tt *test.TCase) *cobra.Command {

	// ユースケースの準備
	usecase := test.NewUserUseCaseMock()
	method, _ := tt.GetMethod()
	inter, ok := method.(*test.UserUseCaseMethod)
	if ok {
		usecase.MockFinds = inter.Finds
	} else {
		f := func() ([]*model.User, error) { return nil, nil }
		usecase.MockFinds = f
	}

	// コントローラーの準備
	controller := test.NewUserControllerMock()

	// ハンドラの準備
	findsHandler := handler.NewUserFindsHandler()

	handle := findsHandler.GetHandle()
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
	controllerCmd.AddCommand(findsHandler.Handle())

	// Rootcmdへコマンドコントローラーを登録
	cmd.AddCommand(controllerCmd)
	cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
	return cmd
}

// Test User Finds Handler
func TestUserFindsHandlerCalled_Handle(t *testing.T) {
	title := fmt.Sprintf("[Validation Behavior]")
	command := "finds"
	fmt.Println()
	cases := []*test.TCase{
		// 引数がない場合はエラー
		test.Case(
			title,
			"Test not enough argument.",
			test.SetCommand(command),
			test.SetExpectCode(codes.NotEnoughArgument),
			test.SetExpectCodeMsg("Return code expected NotEnoughArgument."),
		),
		// 引数のidが指定されたフォーマットじゃない場合はエラー
		test.Case(
			title,
			"Test argument statement incorrect format.",
			test.SetCommand(command),
			test.SetArgument("id", "12"),
			test.SetExpectCode(codes.InvalidArgument),
			test.SetExpectCodeMsg("Return code expected InvalidArgument."),
		),
		// 引数のidが指定されたキャラクタセットじゃなければエラー
		test.Case(
			title,
			"Test argument statement use incorrect character.",
			test.SetCommand(command),
			test.SetArgument("id", "1234#-567%8-1..5-6**90"),
			test.SetExpectCode(codes.InvalidArgument),
			test.SetExpectCodeMsg("Return code expected InvalidArgument."),
		),
	}
	title = fmt.Sprintf("[Processing Behavior]")
	cases = append(cases, []*test.TCase{
		// 検索処理が正常終了
		test.Case(
			title,
			"Test normaly correct operation.",
			test.SetCommand(command),
			test.SetArgument("id", "1234-A78B-12ID-6789"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetFinds(func() ([]*model.User, error) {
						fmt.Fprintf(
							test.Buffer,
							"テストは通るが仕様が確定していない為本来はNG ",
						)
						return nil, nil
					},
					),
			),
			test.SetExpectedNil(),
			test.SetExpectedMsg("Expected operation Successfully"),
		),
		// 検索処理が異常終了
		test.Case(
			title,
			"Test return internal server error.",
			test.SetCommand(command),
			test.SetArgument("id", "1234-A78B-12ID-6789"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetFinds(func() ([]*model.User, error) {
						return nil, fmt.Errorf("Mocking Dummy Error")
					},
					),
			),
			test.SetExpectCode(codes.InternalServerError),
			test.SetExpectCodeMsg("Return code expected InternalServerError."),
		),
	}...)
	test.TestRun(t, cases, SetupFindsHandlerTest, "UserFindsHandler")
}

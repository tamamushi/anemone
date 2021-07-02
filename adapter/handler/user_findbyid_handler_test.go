/* vim: set ts=4 sw=4: */

package handler_test

import (
	"fmt"
	"testing"

	"anemone/test"
	"github.com/spf13/cobra"

	"anemone/adapter/handler"
	"anemone/codes"
)

func SetupFindByIdHandlerTest(t *testing.T, tt *test.TCase) *cobra.Command {

	// ユースケースの準備
	usecase := test.NewUserUseCaseMock()
	method, _ := tt.GetMethod()
	inter, ok := method.(*test.UserUseCaseMethod)
	if ok {
		usecase.MockFindById = inter.FindById
	} else {
		usecase.MockFindById = func(id string) error { return nil }
	}

	// コントローラーの準備
	controller := test.NewUserControllerMock()

	// ハンドラの準備
	userHandler := handler.NewUserFindByIdHandler()

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
	controllerCmd := controller.Handler()
	controllerCmd.AddCommand(userHandler.Handle())

	// Rootcmdへコマンドコントローラーを登録
	cmd.AddCommand(controllerCmd)
	cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
	return cmd
}

// Test User FindById Handler
func TestUserFindByIdHandlerCalled_Handle(t *testing.T) {
	title := fmt.Sprintf("[Validation Behavior]")
	command := "findbyid"
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
					SetFindById(func(id string) error {
						fmt.Fprintf(
							test.Buffer,
							"テストは通るが仕様が確定していない為本来はNG ",
						)
						return nil
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
					SetFindById(func(id string) error {
						return fmt.Errorf("Mocking Dummy Error")
					},
					),
			),
			test.SetExpectCode(codes.InternalServerError),
			test.SetExpectCodeMsg("Return code expected InternalServerError."),
		),
	}...)
	test.TestRun(t, cases, SetupFindByIdHandlerTest, "UserFindByIdHandler")
}

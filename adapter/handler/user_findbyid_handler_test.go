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
	"anemone/errors"
)

func SetupFindByIdHandlerTest(
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
	output := test.PrepareOutputMock(tt)
	// Parserは薄いのでここでは実体を使う
	parser := handler.NewUserFindByIdParser(format, output)
	// 各ハンドラ毎にパーサーのラッパーを用意。これは、
	// 各ハンドラで必要なメソッドの違いを吸収し、同一メソッド名
	// でハンドラによる処理の違いを実現する為に存在する

	// ハンドラの準備
	userHandler := handler.NewUserFindByIdHandler(parser)
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

	// UserControllerにUserFindByIdハンドラを登録
	controllerCmd := controller.Handler(g)
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
	result, _ := json.Marshal(user)
	cases = append(cases, []*test.TCase{
		// 検索処理が正常終了
		test.Case(
			title,
			"Test normaly correct operation.",
			test.SetCommand(command),
			test.SetArgument("id", "1234-A78B-12ID-6789"),
			// Controllerで呼び出されたUseCaseにより、modelが返される。
			// ここでは、&model.User{"test", "test"}
			// 返されたmodel.UserはParser.Outputに渡される。
			// Outputは渡されたmodelをそのままJson.Marshalしstringに変換
			// 変換されたStringはControllerでgateway.SetResponseされ、
			// gatewayから抽出可能となる。
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetFindById(func(id string) (*model.User, error) {
						return user, nil
					},
					),
			),
			test.SetParser(
				test.GetParserMethodStruct().
					SetOutput(func(m interface{}) string {
						model, _ := json.Marshal(m)
						return string(model)
					},
					),
			),
			// 期待値がnilの場合。つまり処理が正常に終了しreturnがnilの場合
			// 正常終了判定とするフラグメソッド
			//test.SetExpectedNil(),
			test.SetExpected(string(result)),
			test.SetExpectedMsg("Expected operation Successfully."),
		),
		// 検索処理が異常終了
		test.Case(
			title,
			"Test return internal server error.",
			test.SetCommand(command),
			test.SetArgument("id", "1234-A78B-12ID-6789"),
			test.SetMethod(
				test.GetUserUseCaseMethodStruct().
					SetFindById(func(id string) (*model.User, error) {
						return nil, fmt.Errorf("Mocking Dummy Error")
					},
					),
			),
			test.SetExpectCode(codes.InternalServerError),
			test.SetExpectCodeMsg("Return code expected InternalServerError."),
		),
	}...)
	test.TestRun(t, cases, SetupFindByIdHandlerTest, "UserFindByIdHandler")
}

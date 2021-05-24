/* vim: set ts=4 sw=4: */

package handler_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"anemone/test"
	"github.com/spf13/cobra"

	"anemone/adapter/handler"

	"anemone/adapter/helper"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
	//	"anemone/model"
)

func Example() {
	// 初期化サンプル
	// init()の伝播の中でコマンド構築を行う

	blder, err := helper.GetBuilderInstance("user")
	if err != nil {
		msg := "Failed to building User command group (%s)"
		fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, err))
		os.Exit(1)
	}
	usecase := usecase.NewUserInteractor()
	handler := handler.NewUserCreateHandler(usecase)
	blder.AddCommand(handler.Handle())
}

var index = 0

func SetupTest(t *testing.T, tt *test.TCase) *cobra.Command {

	// ユースケースの準備
	usecase := test.NewUserUseCaseMock()
	method, _ := tt.GetMethod()
	inter, ok := method.(*test.UserUseCaseMethod)
	fmt.Printf("%#v", ok)
	if ok {
		usecase.MockCreate = inter.Create
		fmt.Println("ok")
	}

	// コントローラーの準備
	controller := test.NewUserControllerMock()

	// ハンドラの準備
	userHandler := handler.NewUserCreateHandler(usecase)

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

func testRun(t *testing.T, cases []*test.TCase) {
	for _, tt := range cases {
		t.Run(tt.GetCaseName(), func(t *testing.T) {
			index++
			//タイトルをBOLD設定
			t.Logf(test.Title(index, tt.GetTestName()))

			cmd := SetupTest(t, tt)
			buf := new(bytes.Buffer)
			cmd.SetOutput(buf)
			err := cmd.Execute()
			//t.Logf("　　\x1b[1mResult:\x1b[0m %s", )

			if err, ok := err.(errors.Errors); ok {
				c := errors.Code(err)
				if c == tt.GetExpectCode() {
					//ExpectCodeが返却されれば期待通り
					t.Logf(test.Expected(tt.GetExpectCodeMsg()))
				} else {
					//ExpectCode以外が返却されたらおかしい
					t.Logf(test.UnExpected(tt.GetExpectCodeMsg(), c))
				}
			} else {
				c := buf.String()
				if c != tt.GetExpected() || c == "" {
					word := fmt.Sprintf("\x1b[31mNG\x1b[0m")
					t.Errorf("　　[%s] %s hoge returned %s", word, tt.GetExpectedMsg(), c)
				} else {
					word := fmt.Sprintf("\x1b[32mOK\x1b[0m")
					t.Logf("　　[%s] %s", word, tt.GetExpectedMsg())
				}
			}
		})
	}
}

// Test User Create Handler
func TestUserCreateHandlerCalled_Handle(t *testing.T) {
	title := fmt.Sprintf("[Validation Behavior]")
	fmt.Printf("\n")
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
			"Test unsupported Method.",
			test.SetExpectCodeMsg("Return code expected UnSupportedMethod."),
			test.SetExpectCode(codes.UnSupportedMethod),
		),
	}
	title = fmt.Sprintf("[Processing Behavior]")
	usecase := test.GetUserUseCaseMethodStruct()
	cases = append(cases, []*test.TCase{
		// 作成処理が正常終了
		test.Case(
			title,
			"Test unsupported Method.",
			test.SetExpectCodeMsg("Return code expected UnSupportedMethod."),
			test.SetMethod(usecase.SetCreate(func() {
				fmt.Printf("HogeHo")
				return
			})),
			test.SetExpectCode(codes.UnSupportedMethod),
		),
		// 作成処理が異常終了
		test.Case(
			title,
			"Test unsupported Method.",
			test.SetExpectCodeMsg("Return code expected UnSupportedMethod."),
			test.SetCommand("UnsupportedMethod"),
			test.SetExpectCode(codes.UnSupportedMethod),
		),
	}...)
	testRun(t, cases)
}

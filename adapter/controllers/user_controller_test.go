/* vim: set ts=4 sw=4: */

package controllers_test

import (
	"bytes"
	"fmt"
	"testing"

	"anemone/test"
	"github.com/spf13/cobra"

	"anemone/adapter/controllers"
	//	"anemone/adapter/helper"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
	//	"anemone/model"
)

type userUseCaseMock struct {
	usecase.IUserUseCase
}

type tCase struct {
	caseName   string
	testName   string
	subCommand string
	tArgument  *test.TestArgument
	//	tMethod       tMethod
	expectCode    codes.Code
	expectCodeMsg string
	expected      string
	expectedMsg   string
}

type option func(*tCase)

func Case(title string, testName string, opts ...option) *tCase {
	tcase := &tCase{
		caseName:  title,
		testName:  testName,
		tArgument: nil,
	}
	for _, opt := range opts {
		opt(tcase)
	}
	return tcase
}
func SetArgument(t *test.TestArgument) option {
	return func(tc *tCase) {
		tc.tArgument = t
	}
}
func SetExpectCodeMsg(s string) option {
	return func(tc *tCase) {
		tc.expectCodeMsg = s
	}
}
func SetExpectCode(c codes.Code) option {
	return func(tc *tCase) {
		tc.expectCode = c
	}
}
func SetExpectedMsg(s string) option {
	return func(tc *tCase) {
		tc.expectedMsg = s
	}
}
func SetExpected(t string) option {
	return func(tc *tCase) {
		tc.expected = t
	}
}
func SetCommand(s string) option {
	return func(tc *tCase) {
		tc.subCommand = s
	}
}

var index = 0

func SetupUserControllerTest(t *testing.T, tt *tCase) *cobra.Command {

	// ユースケースの準備
	usecase := &userUseCaseMock{}

	// コントローラーの準備
	controller := controllers.NewUserController(usecase)

	cmd, args := test.SetupRootCMD(tt.subCommand, tt.tArgument)

	// rootcmdへコマンドコントローラーを登録
	cmd.AddCommand(controller.Handler())
	cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
	return cmd
}

func testRun(t *testing.T, cases []*tCase) {
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			index++
			t.Logf(fmt.Sprintf("[%03d]\x1b[1m%s\x1b[0m", index, tt.testName))
			cmd := SetupUserControllerTest(t, tt)
			buf := new(bytes.Buffer)
			cmd.SetOutput(buf)
			err := cmd.Execute()
			//t.Logf("　　\x1b[1mResult:\x1b[0m %s", )

			if err, ok := err.(errors.Errors); ok {
				c := errors.Code(err)
				if c == tt.expectCode {
					//NotEnoughArgumentが返却されれば期待通り
					word := fmt.Sprintf("\x1b[32mOK\x1b[0m")
					t.Logf("　　[%s] %s", word, tt.expectCodeMsg)
				} else {
					//NotEnoughArgument以外が返却されたらおかしい
					word := fmt.Sprintf("\x1b[31mNG\x1b[0m")
					t.Errorf("　　[%s] %s But returned %s", word, tt.expectCodeMsg, c)
				}
			} else {
				c := buf.String()
				if c != tt.expected || c == "" {
					word := fmt.Sprintf("\x1b[31mNG\x1b[0m")
					t.Errorf("　　[%s] %s hoge returned %s", word, tt.expectedMsg, c)
				} else {
					word := fmt.Sprintf("\x1b[32mOK\x1b[0m")
					t.Logf("　　[%s] %s", word, tt.expectedMsg)
				}
			}
		})
	}
}

func TestUserControllerCalledHandler(t *testing.T) {
	title := fmt.Sprintf("[Just called Behavior]")
	fmt.Printf("\n")
	cases := []*tCase{
		Case(
			title,
			"Test not enough argument.",
			SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			SetCommand(""),
			SetExpectCode(codes.NotEnoughArgument),
		),
		Case(
			title,
			"Test unsupported Method.",
			SetExpectCodeMsg("Return code expected UnSupportedMethod."),
			SetCommand("UnsupportedMethod"),
			SetExpectCode(codes.UnSupportedMethod),
		),
	}
	testRun(t, cases)
}

/*
func TestUserControllerCalledCreate(t *testing.T) {
	title := fmt.Sprintf("[Called Create Method Behavior]")
	fmt.Printf("\n")
	cases := []*tCase{
		Case(
			title,
			" Test method works properly.",
			SetCommand("create"),
			SetExpectedMsg(""),
			SetExpected(""),
		),
		Case(
			title,
			" Test not enough argument.",
			SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			SetCommand("create"),
			SetExpectCode(codes.NotEnoughArgument),
		),
		Case(
			title,
			" Return code expected InvalidArgument.",
			SetExpectCodeMsg("Return code expected InvalidArgument."),
			SetCommand("create"),
			SetExpectCode(codes.InvalidArgument),
		),
		Case(
			title,
			" Return code expected InternalServerError.",
			SetExpectCodeMsg("Return code expected InternalServerError."),
			SetCommand("create"),
			SetExpectCode(codes.InternalServerError),
		),
	}
	testRun(t, cases)
}

func TestUserControllerCalledRemove(t *testing.T) {
	title := fmt.Sprintf("[Called Remove Method Behavior]")
	fmt.Printf("\n")
	cases := []*tCase{
		Case(
			title,
			" Test method works properly.",
			SetCommand("remove"),
			SetExpectedMsg("Return value expected 1234-5678-9123-4567."),
			SetArgument(&tArgument{Name: "id", Value: "1234-5678-9123-4567"}),
			SetMockUseCase(tMethod{Remove: func(id string) error { return nil }}),
			SetExpected("1234-5678-9123-4567"),
		),
		Case(
			title,
			" Test not enough argument.",
			SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			SetCommand("remove"),
			SetExpectCode(codes.NotEnoughArgument),
		),
		Case(
			title,
			" Return code expected InvalidArgument.",
			SetExpectCodeMsg("Return code expected InvalidArgument."),
			SetCommand("remove"),
			SetExpectCode(codes.InvalidArgument),
		),
		Case(
			title,
			" Return code expected InternalServerError.",
			SetExpectCodeMsg("Return code expected InternalServerError."),
			SetCommand("remove"),
			SetArgument(&tArgument{Name: "id", Value: "xxxxx-xxxxx-xxxxx-xxxxx"}),
			SetMockUseCase(tMethod{
				Remove: func(id string) error {
					return fmt.Errorf("usecase remove method error")
				},
			}),
			SetExpectCode(codes.InternalServerError),
		),
	}
	testRun(t, cases)
}

func TestUserControllerCalledUpdate(t *testing.T) {
	title := fmt.Sprintf("[Called Update Method Behavior]")
	fmt.Printf("\n")
	cases := []*tCase{
		Case(
			title,
			" Test method works properly.",
			SetCommand("update"),
			SetExpectedMsg(""),
			SetExpected(""),
		),
		Case(
			title,
			" Test not enough argument.",
			SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			SetCommand("update"),
			SetExpectCode(codes.NotEnoughArgument),
		),
		Case(
			title,
			" Return code expected InvalidArgument.",
			SetExpectCodeMsg("Return code expected InvalidArgument."),
			SetCommand("update"),
			SetExpectCode(codes.InvalidArgument),
		),
		Case(
			title,
			" Return code expected InternalServerError.",
			SetExpectCodeMsg("Return code expected InternalServerError."),
			SetCommand("update"),
			SetExpectCode(codes.InternalServerError),
		),
	}
	testRun(t, cases)
}

func TestUserControllerCalledFindById(t *testing.T) {
	title := fmt.Sprintf("[Called Findbyid Method Behavior]")
	fmt.Printf("\n")
	cases := []*tCase{
		Case(
			title,
			" Test method works properly.",
			SetCommand("findbyid"),
			SetExpectedMsg("Return value expected "),
			SetArgument(&tArgument{Name: "id", Value: "1234-5678-9123-4567"}),
			SetMockUseCase(tMethod{
				FindById: func(id string) (*model.User, error) {
					user := model.User{}
					return &user, nil
				},
			}),
			SetExpected(""),
		),
		Case(
			title,
			" Test not enough argument.",
			SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			SetCommand("findbyid"),
			SetExpectCode(codes.NotEnoughArgument),
		),
		Case(
			title,
			" Return code expected InvalidArgument.",
			SetExpectCodeMsg("Return code expected InvalidArgument."),
			SetCommand("findbyid"),
			SetExpectCode(codes.InvalidArgument),
		),
		Case(
			title,
			" Return code expected InternalServerError.",
			SetExpectCodeMsg("Return code expected InternalServerError."),
			SetCommand("findbyid"),
			SetArgument(&tArgument{Name: "id", Value: "xxxxx-xxxxx-xxxxx-xxxxx"}),
			SetMockUseCase(tMethod{
				FindById: func(id string) (*model.User, error) {
					return nil, fmt.Errorf("usecase findbyid method error")
				},
			}),
			SetExpectCode(codes.InternalServerError),
		),
	}
	testRun(t, cases)
}

func TestUserControllerCalledFinds(t *testing.T) {
	title := fmt.Sprintf("[Called Finds Method Behavior]")
	fmt.Printf("\n")
	cases := []*tCase{
		Case(
			title,
			" Test method works properly.",
			SetCommand("finds"),
			SetExpectedMsg(""),
			SetExpected(""),
		),
		Case(
			title,
			" Test not enough argument.",
			SetExpectCodeMsg("Return code expected NotEnoughArgument."),
			SetCommand("finds"),
			SetExpectCode(codes.NotEnoughArgument),
		),
		Case(
			title,
			" Return code expected InvalidArgument.",
			SetExpectCodeMsg("Return code expected InvalidArgument."),
			SetCommand("finds"),
			SetExpectCode(codes.InvalidArgument),
		),
		Case(
			title,
			" Return code expected InternalServerError.",
			SetExpectCodeMsg("Return code expected InternalServerError."),
			SetCommand("finds"),
			SetExpectCode(codes.InternalServerError),
		),
	}
	testRun(t, cases)
}
*/

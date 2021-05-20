/* vim: set ts=4 sw=4: */

package controllers_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"

	"anemone/adapter/controllers"
	"anemone/adapter/helper"
	"anemone/application/usecase"
	"anemone/codes"
	"anemone/errors"
)

type userUseCaseMock struct {
	usecase.IUserUseCase
	MockCreate func()
}

func (m *userUseCaseMock) Create() {
	m.MockCreate()
}

type tMethod struct {
	Test func()
}

type tCase struct {
	caseName   string
	ExpectMsg  string
	UnExpecMsg string
	subCommand string
	tMethod    tMethod
	arg1       string
	arg2       string
	Evaluate   func(interface{}) bool
}

func SetupUserControllerTest(t *testing.T, tt *tCase) (*cobra.Command, *bytes.Buffer) {

	// ユースケースの準備
	usecase := &userUseCaseMock{}
	usecase.MockCreate = tt.tMethod.Test

	// コントローラーの準備
	controller := controllers.NewUserController(usecase)

	cmd := &cobra.Command{
		Use:   "anemone",
		Short: "A brief description of your application",

		// Usageは出さない
		SilenceUsage: true,
	}
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buf := new(bytes.Buffer)
	cmd.SetOutput(buf)

	// rootcmdへコマンドコントローラーを登録
	cmd.AddCommand(controller.Handler())

	args := helper.NewArgumentBuilder()
	if len(tt.subCommand) != 0 {
		args.AddCommand(tt.subCommand)
	}

	cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
	return cmd, buf
}

func testRun(t *testing.T, cases []tCase) {
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {

			// rootcmdの準備
			cmd, _ := SetupUserControllerTest(t, &tt)
			err := cmd.Execute()
			if err != nil {
				return
			}
		})
	}
}

func TestUserControllerCalledCommand(t *testing.T) {
	t.Run("Called Create Behavior", func(t *testing.T) {

		t.Logf("Called command check.")
		expect := fmt.Sprintf("Return code expected NotEnoughArgument.")
		cmd, _ := SetupUserControllerTest(
			t,
			&tCase{
				tMethod:    tMethod{Test: func() { return }},
				subCommand: "create",
			},
		)
		//t.Logf("log: %s", buf)

		if err := cmd.Execute(); err != nil {
			t.Errorf("　[NG] %s But returned %s", expect, errors.Code(err))
		}
	})
}

func TestUserControllerCheckHandler(t *testing.T) {

	t.Run("Called Behavior", func(t *testing.T) {

		t.Logf("Called Not enough argumnt.")
		expect := fmt.Sprintf("Return code expected NotEnoughArgument.")
		cmd, _ := SetupUserControllerTest(
			t,
			&tCase{
				tMethod:    tMethod{Test: func() { return }},
				subCommand: "",
			},
		)
		err := cmd.Execute()

		if err, ok := err.(error); ok {
			c := errors.Code(err)
			if c == codes.NotEnoughArgument {
				//NotEnoughArgumentが返却されれば期待通り
				t.Logf("　[OK] %s", expect)
			} else {
				//NotEnoughArgument以外が返却されたらおかしい
				t.Errorf("　[NG] %s But returned %s", expect, errors.Code(err))
			}
		}

		t.Logf("Called unsupported Method.")
		cmd, _ = SetupUserControllerTest(
			t,
			&tCase{
				tMethod:    tMethod{Test: func() { return }},
				subCommand: "UnsupportedMethod",
			},
		)
		err = cmd.Execute()

		expect = fmt.Sprintf("Return code expected UnSupportedMethod.")
		if err, ok := err.(error); ok {
			c := errors.Code(err)
			if c == codes.UnSupportedMethod {
				//NotEnoughArgumentが返却されれば期待通り
				t.Logf("　[OK] %s", expect)
			} else {
				//NotEnoughArgument以外が返却されたらおかしい
				t.Errorf("　[NG] %s But returned %s", expect, errors.Code(err))
			}
		}
	})
}

/*
func TestUserController_CalledCreate(t *testing.T) {
	cases := []tCase{
		{"create user successfully",
			"create",
			tMethod{
				Test: func() {
					fmt.Printf("called usecase.FindById(): user found\n\n")
				},
			},
			"", "",
			true,
		},
	}
	testRun(t, cases)
}

func TestUserController_CalledRemove(t *testing.T) {
	cases := []testCase{
		{"user remove successfully",
			"remove",
			testMethod{
				Test: func() {
					fmt.Println("user remove")
				},
			},
			"", "",
			true,
		},
		{"user remove faild",
			"remove",
			testMethod{
				Test: func() {
					fmt.Println("user remove faild")
				},
			},
			"", "",
			false,
		},
	}
	testRun(t, cases)
}

func TestUserController_CalledFindById(t *testing.T) {
	cases := []testCase{
		{"find user successfully",
			"findbyid",
			testMethod{
				Test: func() {
					fmt.Printf("called usecase.FindById(): user found\n\n")
				},
			},
			"--id", "3",
			true,
		},
		{"find user faild",
			"findbyid",
			testMethod{
				Test: func() {
					fmt.Printf("called usecase.FindById(): user not found\n\n")
				},
			},
			"", "",
			false,
		},
	}
	testRun(t, cases)
}
*/

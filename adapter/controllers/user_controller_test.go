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
)

type userUseCaseMock struct {
	usecase.IUserUseCase
	MockCreate func()
}

func (m *userUseCaseMock) Create() {
	m.MockCreate()
}

type testMethod struct {
	Test func()
}

type testCase struct {
	caseName   string
	subCommand string
	testMethod testMethod
	arg1       string
	arg2       string
	want       bool
}

func SetupUserControllerTest(t *testing.T) (*cobra.Command, *bytes.Buffer) {

	cmd := &cobra.Command{
		Use:   "anemone",
		Short: "A brief description of your application",

		// Usageは出さない
		SilenceUsage: true,
	}
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buf := new(bytes.Buffer)
	cmd.SetOutput(buf)
	return cmd, buf
}

func testRun(t *testing.T, cases []testCase) {
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {

			// ユースケースの準備
			usecase := &userUseCaseMock{}
			usecase.MockCreate = tt.testMethod.Test

			// コントローラーの準備
			controller := controllers.NewUserController(usecase)

			// rootcmdの準備
			cmd, stdOut := SetupUserControllerTest(t)

			// rootcmdへコマンドコントローラーを登録
			cmd.AddCommand(controller.Handler())

			args := helper.NewArgumentBuilder()
			if len(tt.subCommand) != 0 {
				args.AddCommand(tt.subCommand)
			}
			//fmt.Printf("argument %s\n", tt.arg1)
			/*
				if (len(tt.arg1) > 0) && len(tt.arg2) > 0 {
					args.AddArgs(tt.arg1, tt.arg2)
				}
			*/
			cmd.SetArgs(append([]string{"user"}, args.GetArgString()...))
			err := cmd.Execute()
			if err != nil {
				// errの中にはエラーコードが入ってる状態にする。
				// errコードを見て、エラーメッセージを出力する。
				t.Fatal(stdOut.String(), err)
			}
		})
	}
}

func TestUserController_RequiredCommand(t *testing.T) {
	cases := []testCase{
		{"No target command?",
			"",
			testMethod{
				Test: func() { return },
			},
			"", "",
			true,
		},
	}
	testRun(t, cases)
}

func TestUserController_CalledCreate(t *testing.T) {
	cases := []testCase{
		{"create user successfully",
			"create",
			testMethod{
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

/*
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

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
	subCommand string
	tMethod    tMethod
	arg1       string
	arg2       string
	want       bool
}

type Buf = bytes.Buffer
type Cobra = cobra.Command

func SetupUserControllerTest(t *testing.T, tt *tCase) (*Cobra, *Buf) {

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

			err := cmd.Execute.(errors.)
			fmt.Printf("%#v", err)
			t.Logf("Expected faild")
			//fmt.Printf("%v", err.Code())
			fmt.Printf(": %v", err.Code())
			if err != nil {
				t.Logf("%s", err)
			}
		})
	}
}

func TestUserController_RequiredCommand(t *testing.T) {
	cases := []tCase{
		{"No target command?",
			"",
			tMethod{
				Test: func() { return },
			},
			"", "",
			true,
		},
	}
	testRun(t, cases)
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

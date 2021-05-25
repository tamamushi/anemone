/* vim: set ts=4 sw=4: */

package controllers_test

import (
	"fmt"
	"testing"

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
	testMethod testMethod
	subCommand string
	arg1       string
	arg2       string
	want       bool
}

type Option func(*testCase)

func Argkey(title string) Option {
	return func(t *testCase) { t.arg1 = title }
}

func ArgData(data string) Option {
	return func(t *testCase) { t.arg2 = data }
}

func NewTestCase(title string, test testMethod, command string, options ...Option) *testCase {
	testcase := testCase{title, testMethod{test}, command "", "", false}

	for _, option := range options {
		option(&testcase)
	}
	return &testcase
}

func testRun(t *testing.T, cases []testCase) {
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			usecase := &userUseCaseMock{}
			usecase.MockCreate = tt.testMethod.Test

			controller := controllers.NewUserController(usecase)
			cmd := controller.Handler()

			args := helper.NewArgumentBuilder()
			args.AddCommand(tt.subCommand)

			fmt.Println("%s", tt.arg1)
			/*
				if (len(tt.arg1) > 0) && len(tt.arg2) > 0 {
					args.AddArgs(tt.arg1, tt.arg2)
				}
			*/

			cmd.SetArgs(args.GetArgString())
			cmd.Execute()
		})
	}
}

func TestUserController_CalledCreate(t *testing.T) {
	cases := []testCase{
		NewTestCase("user create successfully",
			testMethod{
				Test: func() {
					fmt.Printf("called usecase.Create(): user create\n\n")
				},
			},
			"create",
			true,
		},
	}
	testRun(t, cases)
}

/*
func TestUserController_CalledRemove(t *testing.T) {
	cases := []testCase{
		{"user remove successfully",
			testMethod{
				Test: func() {
					fmt.Println("user remove")
				},
			},
			"remove",
			true,
		},
		{"user remove faild",
			testMethod{
				Test: func() {
					fmt.Println("user remove faild")
				},
			},
			"remove",
			false,
		},
	}
	testRun(t, cases)
}

func TestUserController_CalledFindById(t *testing.T) {
	cases := []testCase{
		{"find user successfully",
			testMethod{
				Test: func() {
					fmt.Printf("called usecase.FindById(): user found\n\n")
				},
			},
			"findbyid",
			true,
		},
		{"find user faild",
			testMethod{
				Test: func() {
					fmt.Printf("called usecase.FindById(): user not found\n\n")
				},
			},
			"findbyid",
			false,
		},
	}
	testRun(t, cases)
}
*/

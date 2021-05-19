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
	subCommand string
	testMethod testMethod
	arg1       string
	arg2       string
	want       bool
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

			fmt.Printf("argument %s\n", tt.arg1)
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

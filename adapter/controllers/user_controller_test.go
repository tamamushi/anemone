/* vim: set ts=4 sw=4: */

package controllers_test

import (
	"fmt"
	"testing"

	"anemone/adapter/controllers"
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
	watn       bool
}

func TestUserController_CalledCreate(t *testing.T) {

	cases := []testCase{
		{"user create successfully",
			testMethod{
				Test: func() {
					fmt.Println("user create")
				},
			},
			"create",
			true,
		},
		{"user create faild",
			testMethod{
				Test: func() {
					fmt.Println("user create faild")
				},
			},
			"create",
			false,
		},
	}
	testRun(t, cases)
}

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

func testRun(t *testing.T, cases []testCase) {
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			usecase := &userUseCaseMock{}
			usecase.MockCreate = tt.testMethod.Test

			controller := controllers.NewUserController(usecase)
			cmd := controller.Handler()
			cmd.SetArgs([]string{tt.subCommand})
			cmd.Execute()
		})
	}
}

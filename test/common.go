/* vim: set ts=4 sw=4: */

package test

import (
	"fmt"

	"github.com/spf13/cobra"

	"anemone/adapter/helper"
	"anemone/codes"
)

type TestArgument struct {
	Name  string
	Value string
}

func SetupRootCMD(tt *TCase) (*cobra.Command, helper.ArgumentBuilder) {

	cmd := &cobra.Command{
		Use:   "anemone",
		Short: "A brief description of your application",

		// Usageは出さない
		SilenceUsage: true,
	}
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	args := helper.NewArgumentBuilder()
	if len(tt.GetCommand()) != 0 {
		args.AddCommand(tt.GetCommand())
	}
	a := tt.GetArgument()
	if a != nil {
		args.AddArgs(a.Name, a.Value)
	}
	return cmd, args
}

func Title(i int, s string) string {
	return fmt.Sprintf("[%03d]\x1b[1m%s\x1b[0m", i, s)
}

func Expected(expect string) string {
	return fmt.Sprintf("　　[\x1b[32mOK\x1b[0m] %s", expect)
}

func UnExpected(expect string, ret codes.Code) string {
	return fmt.Sprintf("　　[\x1b[31mNG\x1b[0m] %s But returned %s", expect, ret)
}

type TCase struct {
	caseName      string
	testName      string
	subCommand    string
	tArgument     *TestArgument
	expectCode    codes.Code
	expectCodeMsg string
	expected      string
	expectedMsg   string
	tMethod       interface{}
}

type option func(*TCase)

func Case(title string, testName string, opts ...option) *TCase {
	tcase := &TCase{
		caseName:  title,
		testName:  testName,
		tArgument: nil,
		tMethod:   nil,
	}
	for _, opt := range opts {
		opt(tcase)
	}
	return tcase
}

func (t *TCase) GetTestName() string {
	return t.testName
}
func (t *TCase) GetCaseName() string {
	return t.caseName
}

// Set/Get Method
func (t *TCase) GetMethod() (interface{}, bool) {
	if t.tMethod != nil {
		return t.tMethod, true
	}
	return nil, false
}
func SetMethod(t interface{}) option {
	return func(tc *TCase) {
		tc.tMethod = t
	}
}

// Set/Get Command
func (t *TCase) GetCommand() string {
	return t.subCommand
}
func SetCommand(s string) option {
	return func(tc *TCase) {
		tc.subCommand = s
	}
}

// Set/Get Argument
func (t *TCase) GetArgument() *TestArgument {
	return t.tArgument
}
func SetArgument(t *TestArgument) option {
	return func(tc *TCase) {
		tc.tArgument = t
	}
}

// Set/Get ExpectCode
func (t *TCase) GetExpectCode() codes.Code {
	return t.expectCode
}
func SetExpectCode(c codes.Code) option {
	return func(tc *TCase) {
		tc.expectCode = c
	}
}

// Set/Get ExpectCodeMsg
func (t *TCase) GetExpectCodeMsg() string {
	return t.expectCodeMsg
}
func SetExpectCodeMsg(s string) option {
	return func(tc *TCase) {
		tc.expectCodeMsg = s
	}
}

// Set/Get ExpectedMsg
func (t TCase) GetExpectedMsg() string {
	return t.expectedMsg
}
func SetExpectedMsg(s string) option {
	return func(tc *TCase) {
		tc.expectedMsg = s
	}
}

// Set/Get Expected
func (t *TCase) GetExpected() string {
	return t.expected
}
func SetExpected(t string) option {
	return func(tc *TCase) {
		tc.expected = t
	}
}

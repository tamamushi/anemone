/* vim: set ts=4 sw=4: */

package test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"

	"anemone/adapter/helper"
	"anemone/codes"
	"anemone/errors"
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

func OK() string {
	return "\x1b[32mOK\x1b[0m"
}
func NG() string {
	return "\x1b[31mNG\x1b[0m"
}

func ExpectedError(expect string) string {
	return fmt.Sprintf("　　[\x1b[32mOK\x1b[0m] %s", expect)
}

func UnExpectedError(expect string, ret codes.Code) string {
	return fmt.Sprintf("　　[\x1b[31mNG\x1b[0m] %s But returned %s", expect, ret)
}

func ExpectedOperation(expect string) string {
	return fmt.Sprintf("　 [\x1b[32mOK\x1b[0m] %s", expect)
}

func UnExpectedOperation(expect string, ret string) string {
	return fmt.Sprintf("　 [\x1b[31mNG\x1b[0m] %s But returned %s", expect, ret)
}

var index = 0
var Buffer *bytes.Buffer

func TestRun(
	t *testing.T,
	cases []*TCase,
	setupfunc func(t *testing.T, tt *TCase) *cobra.Command,
	testName string,
) {
	fmt.Printf("\x1b[1mTesting [%s] \x1b[0m\n", testName)
	for _, tt := range cases {
		t.Run(tt.GetCaseName(), func(t *testing.T) {
			index++
			//タイトルをBOLD設定
			t.Logf(Title(index, tt.GetTestName()))

			cmd := setupfunc(t, tt)
			//buf := new(bytes.Buffer)
			Buffer = new(bytes.Buffer)
			cmd.SetOut(Buffer)
			cmd.SetErr(Buffer)
			err := cmd.Execute()
			result := "NOOP"

			if tt.GetExpectCodeMsg() != "" {
				if err, ok := err.(errors.Errors); ok {
					// エラーの場合、エラーコードにより正しさを判定
					c := errors.Code(err)
					if c == tt.GetExpectCode() {
						//ExpectCodeが返却されれば期待通り
						t.Logf(ExpectedError(tt.GetExpectCodeMsg()))
						result = OK()
					} else {
						//ExpectCode以外が返却されたらおかしい
						t.Errorf(UnExpectedError(tt.GetExpectCodeMsg(), c))
						result = NG()
					}
				} else {
					//ExpectCodeMsgのテストなのにNilが返却されたらおかしい
					t.Errorf(UnExpectedError(tt.GetExpectCodeMsg(), "Nil"))
					result = NG()
				}
			} else {
				// nil が返却された場合
				if err == nil && tt.GetExpected() == "OK" {
					// 正常動作完了
					t.Logf(ExpectedOperation(tt.GetExpectedMsg()))
					fmt.Fprintf(Buffer, "正常動作\n")
					result = OK()
				} else {
					c := Buffer.String()
					t.Errorf(UnExpectedOperation(tt.GetExpectedMsg(), c))
					result = NG()
				}
			}
			fmt.Printf("[%03d]:[%s]", index, result)
			fmt.Printf(" %s", Buffer)
			if Buffer.String() == "" {
				fmt.Printf("\n")
			}
		})
	}
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
func SetArgument(k string, v string) option {
	t := &TestArgument{k, v}
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
func SetExpectedNil() option {
	return func(tc *TCase) {
		tc.expected = "OK"
	}
}

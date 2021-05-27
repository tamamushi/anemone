/* vim: set ts=4 sw=4: */

package codes_test

import (
	"fmt"
	"testing"

	"anemone/codes"
)

func TestCodes(t *testing.T) {
	for _, tt := range []struct {
		title string
		code  codes.Code
		want  string
	}{
		{
			title: "Checking OK Code as expected",
			code:  codes.OK,
			want:  "OK",
		},
		{
			title: "Checking NotEnoughArgument Code as expected",
			code:  codes.NotEnoughArgument,
			want:  "NotEnoughArgument",
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			codeStr := fmt.Sprintf("%s", tt.code)

			explain := fmt.Sprintf("Normaliy: code == want (%s) expected true", tt.want)
			if codeStr != tt.want {
				t.Fatal(fmt.Sprintf("return %s:, Not %s", tt.want, tt.want))
				t.Logf("[Faild] %s", explain)
			} else {
				t.Logf("[OK] %s", explain)
			}

			explain = fmt.Sprintf("UnNormaliy: code == code.ForTestcode expected false")
			if tt.code == codes.ForTestCode {
				t.Fatal(fmt.Sprintf("%s is not %s", tt.code, codes.ForTestCode))
				t.Logf("[Faild] %s", explain)
			} else {
				t.Logf("[OK] %s", explain)
			}
		})
	}
}

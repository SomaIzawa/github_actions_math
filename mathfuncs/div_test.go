package mathfuncs_test

import (
	"github_actions_math/mathfuncs"
	"testing"
)

func Test_Div_1(t *testing.T) {
	resp, _ := mathfuncs.Div(10, 5)
	if resp != 2 {
		t.Fatal("error: unexpected result: [10, 5]")
	}
	resp, _ = mathfuncs.Div(0, 5)
	if resp != 0 {
		t.Fatal("error: unexpected result: [0, 5]")
	}
	_, err := mathfuncs.Div(0, 0)
	if err == nil {
		t.Fatal("error: unexpected result: [0, 0]")
	}
}
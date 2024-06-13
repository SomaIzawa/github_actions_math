package mathfuncs_test

import (
	"github_actions_math/mathfuncs"
	"testing"
)

func Test_Add_1(t *testing.T) {
	if mathfuncs.Add(1, 1) != 2 {
		t.Fatal("failed test")
	}
}

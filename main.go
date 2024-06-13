package main

import (
	"fmt"
	"github_actions_math/mathfuncs"
)

func main() {
	// 100 + 2 = ?
	fmt.Println("100 + 2 =", mathfuncs.Add(100, 2))
	// 100 ÷ 2 = ?
	resp, _ := mathfuncs.Div(100, 2)
	fmt.Println("100 ÷ 2 =", resp)
}

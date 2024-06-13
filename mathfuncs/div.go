package mathfuncs

import "fmt"

func Div(num1, num2 int) (int, error)  {
	if (num2 == 0){
		return 0, fmt.Errorf("error: denominator is zero")
	}
	return num1 / num2, nil
}
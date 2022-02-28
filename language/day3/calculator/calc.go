package calculator

import "errors"

func Add(a,b int) int {
	return a+b
}

func Sub(a,b int) int {
	return a-b
}

func Mul(a,b int) int {
	return a*b
}

func Mod(a,b int) int,error {
	if b==0{
		return 0,errors.New("second operand cannot be zero")
	}
	return a%b,nil
}
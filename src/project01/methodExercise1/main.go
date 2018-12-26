package main

import (
	"fmt"
)

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("输入有误")		
	}
	return res
}
func main(){
	calcuator := Calcuator{10.2,3.2}
	res := calcuator.getRes('*')
	fmt.Printf("res=%.2f",res)
}
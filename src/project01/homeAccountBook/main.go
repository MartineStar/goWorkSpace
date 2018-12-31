package main

import (
	"fmt"
)
func main(){
	key :=""
	loop := true
	for {
		fmt.Println("----------------------家庭收支记账软件--------------------")
		fmt.Println()
		fmt.Println("                        1.收支明细")
		fmt.Println("                        2.登记收入")
		fmt.Println("                        3.登记支出") 
		fmt.Println("                        4.退   出")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key{
		case "1":
			fmt.Println("----------------------当前收支明细纪录--------------------")
			fmt.Println()
		case "2":
			// fmt.Println("----------------------当前收支明细纪录--------------------")
		
		case "3":
			// fmt.Println("----------------------当前收支明细纪录--------------------")
		
		case "4":
			loop = false	//此处，break只能退出switch，而不能退出for循环
		default:
			fmt.Println("请输入正确的选项")		
		}
		if !loop{
			break
		}

	}
	fmt.Println("您已退出家庭收支软件，拜拜~")
}
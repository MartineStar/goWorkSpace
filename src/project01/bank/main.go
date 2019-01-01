package main

import (
	"fmt"
)


//账户结构体
type account struct {
	user string
	password string
	balance float64
}
//存钱
func (accountdemo account) Deposit(user string,pwd string,money float64) {
	var count int = 3
	if pwd != accountdemo.password {
		for{
			count--			
			fmt.Printf("对不起，您输入的密码有误！请重新输入(剩余次数%d)：\n",count)
			fmt.Scanln(&pwd)
			if pwd == accountdemo.password{
				break
			}else if count==0 {
				fmt.Printf("登陆失败，即将退出...\n")
				return 
			}
		}
	}
	fmt.Println("登陆成功")	
	fmt.Printf("%v存入金额%v",user,money)
	if money <= 0 {
		fmt.Println("请存入大于0元的金额")
	}else{
		accountdemo.balance += money
		fmt.Println("存入成功",money,"元")
	}	
}
func main(){

	account1 := account{
		user : "biob",
		password : "666666",
		balance:100,
	}
	account1.Deposit("biob","123456",20.8)
}
package main
import (
	"fmt"
	"os"
	"chatroom/client/process"
)

//定义两个变量，一个表示用户id,一个表示用户密码
var userId int
var userPwd string
var userName string

func main(){
	//接收用户的选择
	var key int

	for {
		fmt.Println("---------------欢迎登陆多人聊天系统--------------------")
		fmt.Println("\t\t\t 1. 登陆聊天室")
		fmt.Println("\t\t\t 2. 注册用户")
		fmt.Println("\t\t\t 3. 退出系统")
		fmt.Print("\t\t\t 4. 请选择(1-3):")
		//注意此处的\n不能省，才能达到换行效果
		fmt.Scanf("%d\n",&key)
		switch key {
			case 1:
				fmt.Println("登陆聊天室")
				//用户要登陆
				fmt.Print("请输入用户id:")
				//注意\n要带上
				fmt.Scanf("%d\n",&userId)
				fmt.Print("请输入用户密码:")
				fmt.Scanf("%s\n",&userPwd)
				userProcess := &process.UserProcess{}
				//先把登陆的函数，写到另外一个文件，login.go
				userProcess.Login(userId,userPwd)
				
			case 2:
				fmt.Println("注册用户")
				fmt.Print("请输入用户id:")
				//注意\n要带上
				fmt.Scanf("%d\n",&userId)
				fmt.Print("请输入用户密码:")
				fmt.Scanf("%s\n",&userPwd)
				fmt.Print("请输入用户名:")
				fmt.Scanf("%s\n",&userName)
				userProcess := &process.UserProcess{}
				userProcess.Register(userId,userPwd,userName)
				os.Exit(0)

			case 3:
				fmt.Println("退出系统")
				os.Exit(0)
			default:
				fmt.Println("您的输入有误，请重新输入！")
		}
	}
	// //根据用户的输入，显示新的提示信息
	// if key ==1 {
	// 	//用户要登陆
	// 	fmt.Print("请输入用户id:")
	// 	//注意\n要带上
	// 	fmt.Scanf("%d\n",&userId)
	// 	fmt.Print("请输入用户密码:")
	// 	fmt.Scanf("%s\n",&userPwd)
	// 	//先把登陆的函数，写到另外一个文件，login.go
	// 	login(userId,userPwd)

	// } else if key ==2 {
	// 	fmt.Println("处理注册逻辑")
	// }
}
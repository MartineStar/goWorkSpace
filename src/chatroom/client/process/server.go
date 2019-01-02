package process

import (
	"fmt"
	"os"
	"net"
	"chatroom/client/utils"
)

//显示登陆成功后的界面..
func ShowMenu(){
	fmt.Println("----------恭喜xxx登陆成功------------")
	fmt.Println("--------1. 显示在线用户列表------------")
	fmt.Println("--------2. 发送消息       ------------")
	fmt.Println("--------3. 信息列表       ------------")
	fmt.Println("--------4. 退出系统       ------------")
	fmt.Printf("请选择(1-4):\t")
	var key int
	fmt.Scanf("%d\n",&key)
	switch key{
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("您选择了退出系统...")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误,请重新输入...")
	}

}

//和服务器端保持通讯

func serverProcessMes(Conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发送的消息
	//只要服务器端不断开链接，客户端不断开链接，此处读消息就能一直保持阻塞直到读取到消息
	//从而达到和服务器保持通讯的目的
	//声明一个transfer实例，不停的读取服务端发送的消息，
	fmt.Println("客户端保持与服务器端链接的协程开启")
	tf := &utils.Transfer{
		Conn:Conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息...")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg failed,error=",err)
			return
		}
		//如果读取到消息，进行下一步处理
		fmt.Printf("mes=%v\n",mes)
	}
}
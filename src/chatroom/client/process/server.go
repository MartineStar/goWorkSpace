package process

import (
	"fmt"
	"os"
	"net"
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
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
	var content string
	smsProcess := SmsProcess{} 
	fmt.Scanf("%d\n",&key)
	switch key{
		case 1:
			outputOnlineUser()
		case 2:
			// fmt.Println("发送消息")
			fmt.Println("您发送群消息： ")
			fmt.Scanf("%s\n",&content)
			//调用smsProcess的发送群聊的方法
			smsProcess.SendGroupMes(content)

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
		switch mes.Type {
			case message.NotifyUserStatusMesType://系统推送有人上线信息
				//1.取出NotifyUserStatusMes
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
				//2.把这个用户的信息，状态保存到客户map[int]User中
				updateUserStatus(&notifyUserStatusMes)
				//3.处理
			case message.SmsMesType://有人群发消息
				outputGroupMes(&mes)
			default:
				fmt.Println("服务器端返回了未知的消息类型")
			
		}
	}
}
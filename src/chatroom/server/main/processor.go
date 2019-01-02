package main

import (
	"fmt"
	"net"
	"chatroom/common/message"
	"chatroom/server/process"
	"chatroom/server/utils"
	"io"
)

//创建Processor结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcesssMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
//此处形参为什么mes传递的是指针？？
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType:
			//处理登陆逻辑
			userProcess := &process.UserProcess{
				Conn : this.Conn,
			}

			err = userProcess.ServerProcessLogin(mes)

		case message.RegisterMesType:
			//处理注册逻辑
			userProcess := &process.UserProcess{
				Conn : this.Conn,
			}

			err = userProcess.ServerProcessRegister(mes)
			
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return
}



func (this *Processor)handler() (err error){
	//循环等待客户端发送链接
	for {
		//创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		var mes message.Message
		mes,err = tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出！")
				return
			}
			fmt.Println("readPkg failed,error=",err)			
			return
		}
		
		err = this.serverProcessMes(&mes)
		if err != nil {
			return
		}
	}
	
}

//报错：err is shadowed during return,原因:返回值列表中已经有err,再函数内又定义了err,程序不明确返回哪个err,指明err即可
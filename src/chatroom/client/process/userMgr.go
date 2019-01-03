package process

import (
	"fmt"
	"chatroom/common/message"
	"chatroom/client/model"
)

//客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)
//表示当前用户，在用户登陆成功后完成初始化
var CurUser model.CurUser

//在客户端显示当前在线的用户
func outputOnlineUser(){
	//遍历onlineUsers
	fmt.Println("当前在线用户列表为：")
	for id,_ := range onlineUsers{
		fmt.Println("用户id:\t",id)
	}
}

//编写方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){
	//先判断是否onlineUser中是否有该用户，有则更新状态，没有则添加用户
	//这样做相比直接添加元素(效果一样)，更提升效率
	user,ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
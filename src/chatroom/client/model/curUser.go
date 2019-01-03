package model

import (
	"net"
	"chatroom/common/message"
)

//定义一个结构体，用于保存当前用户信息
//因为这个结构体不是共用的，不能放在common/message中
type CurUser struct {
	Conn net.Conn
	message.User
}

//因为在客户端，很多地方回使用到CurUser,在userMgr文件中将其设置为全局
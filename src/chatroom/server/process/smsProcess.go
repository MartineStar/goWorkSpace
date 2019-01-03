package process

import (
	 "fmt"
	 "encoding/json"
	 "chatroom/common/message"
	 "net"
	 "chatroom/server/utils"
)

type SmsProcess struct {
	//...
}

//转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	//取出mes里的内容smsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal smsMes failed,error=",err)
		return 
	}

	//因为时转发消息，完全可以不用涉及到里面的内容，把消息原封不动的转发给对方就可以了
	data,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal failed,error=",err)
		return
	}
	
	//遍历服务器的onlineUsers map[int]*UserProcess,转发消息	
	for id,userProcess := range userMgr.onlineUsers {
		//过滤掉自己
		if id == smsMes.UserId{
			continue
		}
		this.SendMesToEachOnlineUser(data,userProcess.Conn)
	}
}

//编写单个转发消息的函数
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte,conn net.Conn) {
	//创建transfer实例，发送data
	tf := &utils.Transfer{
		Conn:conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Printf("转发消息to %v 失败,error=%v",conn,err)
	}

}
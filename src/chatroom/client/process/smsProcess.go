package process
import (
	"fmt"
	"chatroom/client/utils"
	"encoding/json"
	"chatroom/common/message"
)

type SmsProcess struct {
	//...
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error){
	//创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content

	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	
	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal failed,error=",err)
		return
	}
	mes.Data = string(data)

	//序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal failed,error=",err)
		return
	}

	//发送消息
	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("sendGroup send failed,error=",err)
		return 
	}
	return
}
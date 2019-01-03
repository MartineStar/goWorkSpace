package process

import (
	"fmt"
	"encoding/json"
	"chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	//传进来的mes类型一定是SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed,error=",err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户[%v]说：%v\n",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
}
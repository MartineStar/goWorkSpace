package process

import (
	"fmt"
	"net"
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"	
	"encoding/json"
)

type UserProcess struct {
	Conn net.Conn
}

//编写也该serverProcessLogin函数，专门处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data,并直接反序列化成registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)

	if err != nil {
		fmt.Println("json.Unmarshal failed,error=",err)
		return
	}

	//2.先声明一个resMes,用于服务器返回消息给客户端
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//3.声明一个 LoginResMes，并完成赋值
	var registerResMes message.RegisterResMes

	//4.到redis数据库完成注册
	//1.使用model.MyUserDao 到redis验证
	err = model.MyUserDao.Register(&registerMes.User)	

	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXIST.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "register encounter unknow error"
		}
	}else {
		registerResMes.Code = 200
	}

	//5.将registerResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal registerResMes failed,error=",err)
		return
	}

	//6.将data赋值给resMes
	resMes.Data = string(data)

	//7.对resMes 进行序列化，准备发送给客户端
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal resMes failed,error",err)
		return
	}

	//8.发送到客户端
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//编写也该serverProcessLogin函数，专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	
	if err != nil {
		fmt.Println("json.Unmarshal failed,error=",err)
		return
	}

	//2.先声明一个resMes,用于服务器返回消息给客户端
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//3.声明一个 LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//4.到redis数据库完成验证
	//1.使用model.MyUserDao 到redis验证
	user, err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)	
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500 //500状态码，表示该用户不存在
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403 //500状态码，表示该用户不存在
			loginResMes.Error = err.Error()
		}else {
			loginResMes.Code = 505 //500状态码，表示该用户不存在
			loginResMes.Error = "服务器内部错误..."
		}
	}else {
		loginResMes.Code = 200
		fmt.Println(user,"登陆成功")
	}

	// //如果用户id=100,密码=123456,则认为合法，否则不合法
	// if loginMes.UserId ==100 && loginMes.UserPwd == "123456"{
	// 	//合法
	// 	loginResMes.Code = 200
	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500 //500状态码，表示该用户不存在
	// 	loginResMes.Error = "该用户不存在，请先注册再登陆！！！"
	// }

	//3.将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal loginResMes failed,error=",err)
		return
	}

	//将data赋值给resMes
	resMes.Data = string(data)

	//对resMes 进行序列化，准备发送给客户端
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal resMes failed,error",err)
		return
	}

	//发送data,因多次使用，将其封装成函数writePkg
	//因为使用了分层模式(mvc),先创建要给Transfer实例
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
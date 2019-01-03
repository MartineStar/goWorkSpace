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
	//增加一个字段，表示该Conn是哪一个用户
	UserId int
}

//编写通知所在线的用户的方法
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onlineUsers,然后每一个发送NotifyUserStatusMes
	for id,userProcess := range userMgr.onlineUsers{
		//过滤自己
		if id == userId{
			continue
		}
		//开始通知每一个人我的在线信息，单独写成一个方法
		userProcess.NotifyMeOnline(userId)
	}

}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data,err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal notifyUserStatusMes failed,error=",err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//对mes再次序列化，准备发送
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal mes failed,error=",err)
		return
	}

	//8.发送到客户端
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyUserStatusMes推送消息失败,error=",err)
	}
	return
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
	if err != nil {
		fmt.Println("返回注册结果给客户端失败,error=",err)
	}
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
		//这里，因为用户登陆成功，我们就把登陆成功的用户放入userMgr中
		//将登陆成功的用户的userId赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		// fmt.Println("在线用户：",userMgr.GetAllOnlineUser())

		//通知其他用户,本账号上线
		this.NotifyOthersOnlineUser(loginMes.UserId)

		//将当前在线用户的id放入loginResMes.UsersId
		for id,_ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId,id)
		}
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
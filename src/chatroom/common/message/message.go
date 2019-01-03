package message
//common包是专门存放公用的东西
//message是存放公用的消息结构体

//消息类型常量
const (
	LoginMesType			= "LoginMes"
	LoginResMesType		 	= "LoginResMes"
	RegisterMesType		 	= "RegisterMes"
	RegisterResMesType		= "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType				="SmsMes"
)

//定义用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus	
)

type Message struct{
	Type string		`json:"type"`	//消息类型,需要先定义一些常量，存放消息类型
	Data string 	`json:"data"`	//消息的内容
}

//配合服务器端推送用户状态变化的消息，服务器推送消息
type NotifyUserStatusMes struct {
	UserId int 		`json:"userId"`
	Status int 		`json:"status"`	//用户状态
}

//先定义两个消息，后面需要再增加
type LoginMes struct {
	UserId int		`json:"userId"`
	UserPwd string	`json:"userPwd"`
	UserName string `json:"userName"`	//用户名
}

type LoginResMes struct {
	Code int 		`json:"code"`	//状态码，500表示该用户未注册，200表示登陆成功，300表示其他错误
	UsersId []int	`json:"userIds"`//用于保存所有在线用户的id
	Error string 	`json:"error"`	//返回错误信息，如果没有错误，则为nil
}

//注册的消息结构体
type RegisterMes struct {
	//...
	User User	`json:"user"`
}

//服务器返回的注册信息
type RegisterResMes struct {
	Code int 		`json:"code"`	//状态码，400表示该用户未注册，200表示登陆成功
	Error string 	`json:"error"`	//返回错误信息，如果没有错误，则为nil
}

//增加Sms结构体，用于发送消息
type SmsMes struct {
	Content string 	`json:"content"`
	User	//匿名结构体，继承
}


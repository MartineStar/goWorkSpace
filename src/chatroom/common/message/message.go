package message
//common包是专门存放公用的东西
//message是存放公用的消息结构体

//消息类型常量
const (
	LoginMesType 	= "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct{
	Type string		`json:"type"`	//消息类型,需要先定义一些常量，存放消息类型
	Data string 	`json:"data"`	//消息的内容
}

//先定义两个消息，后面需要再增加
type LoginMes struct {
	UserId int		`json:"userId"`
	UserPwd string	`json:"userPwd"`
	UserName string `json:"userName"`	//用户名
}

type LoginResMes struct {
	Code int 		`json:"code"`	//状态码，500表示该用户未注册，200表示登陆成功，300表示其他错误
	Error string 	`json:"error"`	//返回错误信息，如果没有错误，则为nil
}

//注册的消息结构体
type RegisterMes struct {
	//...
}
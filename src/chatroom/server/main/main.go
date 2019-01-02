package main
import (
	"fmt"
	"net"
	"time"
	"chatroom/server/model"
)

// func writePkg(conn net.Conn,data []byte) (err error) {
// 	//先发送给一个长度给对方
// 	var pkgLen uint32 
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

// 	//发送长度
// 	_,err = conn.Write(buf[:4])
// 	if err != nil{
// 		fmt.Println("conn.Write(resMes's length) failed,error=",err)
// 		return
// 	}

// 	//发送data本身
// 	n, err := conn.Write(data)
// 	if n != int(pkgLen) || err !=nil {
// 		fmt.Println("conn.Write(resMes's content) failed,error=",err)
// 		return
// 	}
// 	return
// }

// func readPkg(conn net.Conn) (mes message.Message,err error) {
	
// 	buf := make([]byte,8096)
// 	fmt.Println("等待读取客户端发送的消息...")

// 	//conn.Read 在conn没有关闭的情况下才会阻塞
// 	//  如果客户端关闭了conn,则不会阻塞
// 	_,err = conn.Read(buf[:4])	//err再返回值列表中已经默认定义，不需要:=
// 	if err != nil {
// 		// 自定义错误并返回
// 		// err = errors.New("read pkg header<data's length> error")
// 		return
// 	}

// 	//根据buf[:4]转成一个unit32类型
// 	pkgLen := binary.BigEndian.Uint32(buf[0:4])

// 	//根据pkgLen 读取消息内容,将读到的内容放到buf中,n表示实际读取到的字节数
// 	n ,err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		err = errors.New("read pkg body<data's content> error")
// 		return
// 	}

// 	//把buf[:pkgLen]反序列化成 message.Message类型,因为mes是结构体，需要传地址
// 	err = json.Unmarshal(buf[:pkgLen],&mes)  
// 	if err != nil{
// 		fmt.Println("json.Unmarshal failed,error=",err)
// 		return 
// 	}
// 	return
// }

// //编写也该serverProcessLogin函数，专门处理登陆请求
// func serverProcessLogin(conn net.Conn,mes *message.Message) (err error) {
// 	//1.先从mes中取出mes.Data,并直接反序列化成LoginMes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	
// 	if err != nil {
// 		fmt.Println("json.Unmarshal failed,error=",err)
// 		return
// 	}


// 	//2.先声明一个resMes,用于服务器返回消息给客户端
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	//3.声明一个 LoginResMes，并完成赋值
// 	var loginResMes message.LoginResMes

// 	//如果用户id=100,密码=123456,则认为合法，否则不合法
// 	if loginMes.UserId ==100 && loginMes.UserPwd == "123456"{
// 		//合法
// 		loginResMes.Code = 200
// 	} else {
// 		//不合法
// 		loginResMes.Code = 500 //500状态码，表示该用户不存在
// 		loginResMes.Error = "该用户不存在，请先注册再登陆！！！"
// 	}

// 	//3.将loginResMes 序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal loginResMes failed,error=",err)
// 		return
// 	}

// 	//将data赋值给resMes
// 	resMes.Data = string(data)

// 	//对resMes 进行序列化，准备发送给客户端
// 	data,err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal resMes failed,error",err)
// 		return
// 	}

// 	//发送data,因多次使用，将其封装成函数writePkg
// 	err = writePkg(conn,data)
// 	return
// }

// //编写一个ServerProcesssMes函数
// //功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
// //此处形参为什么mes传递的是指针？？
// func serverProcessMes(conn net.Conn,mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		//处理登陆逻辑
// 		err = serverProcessLogin(conn,mes)
// 	case message.RegisterMesType:
// 		//处理注册逻辑
// 	default:
// 		fmt.Println("消息类型不存在，无法处理")
// 	}
// 	return
// }

//处理和客户端的通讯
func processHandler(conn net.Conn) {
	//延时关闭
	defer conn.Close()
	
	//这里调用总控，创建要给Processor
	processor := &Processor{
		Conn : conn,
	}

	err := processor.handler()
	if err != nil {
		fmt.Println("客户端和服务器端通讯协程出现错误,error=",err)
		return
	}

}



//编写一个函数，完成对userDao的初始化任务
func initUserDao(){
	//这里的pool 本身就是main包下的全局变量
	//注意初始化顺序问题：initPool()  -->initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func init(){
	//程序启动时，开始连接池初始化
	initPool("localhost:6379",16,0,300 * time.Second)
	//初始化userDao
	initUserDao()
}

func main(){

	//提示信息
	fmt.Println("服务器在8889端口监听....")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen failed,error=",err)
		return
	}

	defer listen.Close()

	//监听成功，等待客户端链接服务器
	for {
		fmt.Println("等待客户端连接...")
		conn,err := listen.Accept()
		fmt.Println("Accept() connect from ip:",conn.RemoteAddr().String())
		if err != nil{
			fmt.Println("listen.Accept failed,error=",err)
		}
		//一旦连接成功，则启动过一个协程与客户端保持通讯
		go processHandler(conn)
	}
}
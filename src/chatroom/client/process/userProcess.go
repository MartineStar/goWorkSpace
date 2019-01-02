package process

import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"chatroom/common/message"
	"chatroom/client/utils"
)

type UserProcess struct {
	//暂时不需要字段..
}

func (this *UserProcess) Register(userId int,userPwd string,userName string) (err error){
	//1.连接到服务器
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial failed,error=",err)
		return
	}
	//延时关闭,必须加，不然之后会有莫名其妙的错误,都找不到来源
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建要给RegisterMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId  = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registernMes 序列化
	data,err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal failed,error=",err)
		return
	}

	//5.将data切片转化成string并赋给mes.Data字段
	mes.Data = string(data)
	
	//6.将mes 序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal failed,error=",err)
		return
	}	

	//7.发送消息
	//创建Transfer实例
	tf := utils.Transfer{
		Conn : conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(conn) failed,error=",err)
		return
	}

	//8.接收消息
	mes,err = tf.ReadPkg()	
	
	if err != nil {
		fmt.Println("readPkg(conn) failed,error=",err)
		return
	}

	//解析消息，显示对应的注册结果
	var registerResMes message.RegisterResMes
	//将mes的Data部分序列化成RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed,error=",err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("恭喜您，注册成功，请重新登陆")
		// os.Exit(0)
	}else{
		fmt.Println(registerResMes.Error)
		// os.Exit(0)
		// fmt.Println("注册失败")

	}
	return



}


//关联一个用户登陆的方法
//登陆函数
//返回信息最好是error类型而不要为bool,error能更好的描述错误信息
func (this *UserProcess) Login(userId int,userPwd string) (err error){
	//定协议...

	//1.连接到服务器
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial failed,error=",err)
		return
	}
	//延时关闭,必须加，不然之后会有莫名其妙的错误,都找不到来源
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建要给LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId  = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes 序列化
	data,err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal failed,error=",err)
		return
	}

	//5.将data切片转化成string并赋给mes.Data字段
	mes.Data = string(data)
	
	//6.将mes 序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal failed,error=",err)
		return
	}	

	//7.此时，data就是我们要发送给服务器的消息
	//7.1先发送data的长度
	//注意：encoding/binary下的ByteOrder规定了如何将字节序列和16、32或64比特的无符号整数进行互相转换
	// binary里面的 bigEndian大段字节序类型，它存在ByteOrder里面的所有方法
	//uint32：最大能表示的数：（2**32-1）/1024/1024 约为4G,足够
	var pkgLen uint32 
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4],pkgLen)

	//发送长度
	n,err := conn.Write(buf[:4])
	if n != 4 || err != nil{
		fmt.Println("conn.Write(data's length) failed,error=",err)
		return
	}


	fmt.Println(string(data))
	_,err = conn.Write(data)
	if err != nil{
		fmt.Println("conn.Write(data's content) failed,error=",err)
		return
	}

	//创建Transfer实例
	tf := utils.Transfer{
		Conn : conn,
	}
	// //休眠20s
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20s....")
	//处理服务器返回的消息
	mes,err = tf.ReadPkg()	
	
	if err != nil {
		fmt.Println("readPkg(conn) failed,error=",err)
		return
	}

	var loginResMes message.LoginResMes
	//将mes的Data部分序列化成LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed,error=",err)
		return
	}
	if loginResMes.Code == 200 {
		// fmt.Println("登陆成功了")
		//开启一个协程保持与客户端的链接
		go serverProcessMes(conn)
		//1.循环显示登陆成功后的二级菜单
		for {
			ShowMenu()
		}

	}else{
		fmt.Println(loginResMes.Error)
	}
	return
}
package main	
//两个文件同处于main包下,那么main函数可以直接调用main包下另一个文件的方法或变量
//注意：前提是需要将client这个文件夹进行go build，然后再调用生成的文件,
//通过go run main.go运行会报错说（从在同一个mian包下的另一个文件的)方法/变量未定义
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"chatroom/common/message"

)
//登陆函数
//返回信息最好是error类型而不要为bool,error能更好的描述错误信息
func login(userId int,userPwd string) (err error){
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

	_,err = conn.Write(data)
	if err != nil{
		fmt.Println("conn.Write(data's content) failed,error=",err)
		return
	}
	// //休眠20s
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20s....")
	//处理服务器返回的消息
	mes,err = readPkg(conn)	
	
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
		fmt.Println("登陆成功了")
	}else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
package main
import (
	"fmt"
	"net"//网络socket开发时，net包含我们开发所需的所有方法和函数
)
//协程函数，用于处理可以客户端的请求
func process(conn net.Conn){
	//循环接收客户端发送的数据
	defer conn.Close()	//关闭连接
	for {
		//创建一个新的切片
		buf := make([]byte,1024)
		
		//1.等待客户端通过conn发送消息
		//2.如果客户端没有write发送消息，协程就会在此阻塞
		n,err :=conn.Read(buf)
		//底层会自动的维护这个连接，通过互相丢包达到效果
		//一旦连接断开，就会报错，还有其他错误也会报错
		if err != nil{
			fmt.Println("客户端已退出，连接断开...协程即将关闭",err)
			return			
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Printf("[%v]>> %v",conn.RemoteAddr().String(),string(buf[:n]))
	}
}
func main(){
	fmt.Println("服务器开始监听8888端口...")
	//tcp表示使用的网络协议是tcp,
	//0.0.0.0:8888 表示在本地监听8888端口
	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil{
		fmt.Println("监听失败...")
		return
	}
	//main函数退出就关闭监听端口，即延时关闭listen
	defer listen.Close()
	//循环等待客户端来连接
	for{
		//等待客户端连接
		fmt.Println("等待客户端连接...")
		conn,err :=listen.Accept()
		if err != nil{
			fmt.Println("Accept() err=",err)
		}else {
			fmt.Println("Accept() connect from ip:",conn.RemoteAddr().String())
		} 
		go process(conn)
	}
	
}
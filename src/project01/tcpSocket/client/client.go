package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)
func main(){
	conn ,err := net.Dial("tcp","192.168.86.1:8888")
	if err != nil{
		fmt.Println("client dial err=",err)
		return 
	}
	fmt.Println("连接成功：",conn)

	reader := bufio.NewReader(os.Stdin)	//os.Stdin代表标准输入[终端]

	for {
		//从终端读取一行用户的输入，准备发送服务器
		//读取直到读到\n,返回读取到的数据和\n,或\r\n
		fmt.Printf("[%v]>>",conn.LocalAddr().String())
		line,err := reader.ReadString('\n')	
		if err != nil {
			fmt.Println("读取终端内容失败:",err)
		}
		//如果用户输入的是 exit就退出
		line = strings.Trim(line," \r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		//将line发送给服务器
		//报错：no new variables on left side of :=,将:=换成=
		
		_,err = conn.Write([]byte(line+"\n"))
		if err != nil {
			fmt.Println("发送消息失败：",err)
		}		
	}	
}
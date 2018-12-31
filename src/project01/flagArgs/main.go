package main

import (
	"fmt"
	"flag"
)
func main(){
	//定义变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int
	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	// "u" ,就是 -u 指定参数
	// "" , 默认值，当用户没有输入该参数值时使用的默认值
	// "用户名，默认为空",参数说明
	flag.StringVar(&user,"u","","用户名，默认为空")
	flag.StringVar(&pwd,"pwd","123456","密码，默认123456")
	flag.StringVar(&host,"h","localhost","主机名，默认localhost")
	flag.IntVar(&port,"p",3306,"端口号,默认3306")

	//这里有一个非常重要的操作，转换，必须在参数注册后,使用前使用Parse
	flag.Parse()
	//输出结果
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v",user,pwd,host,port)

}
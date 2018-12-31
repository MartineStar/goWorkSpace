package main

import (
	"fmt"
	"strconv"
	"time"
)
//在主线程，开启一个goroutine,该协程每隔一秒输出"hello world"
//在主线程也每隔一秒输出"hello golang",输出10次后，退出程序
//要求主线程和goroutine同时执行

//goroutine一般是以函数为单位开启的
func test(){
	for i:=1;i<=10;i++{
		fmt.Println("<test> hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main(){
	//开启一个协程
	go test()

	for i:=1;i<=10;i++{
		fmt.Println("<main> hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
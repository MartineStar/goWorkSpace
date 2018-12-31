package main

import "fmt"

func main(){
	intChan := make(chan int ,3)
	intChan <- 100
	intChan <- 99
	close(intChan)
	//这时不能再写入数据
	// intChan <- 999	//panic: send on closed channel
	a := <- intChan	//关闭后读取，没有问题
	fmt.Println("ok",a)

	intChan2 := make(chan int,100)
	for i :=0;i < 100 ;i++{
		intChan2 <- i*2
	}
	//遍历，不能使用普通的for循环
	//注意：管道按顺序只返回一个值，管道不存在下标或者键的概念
	//遍历数据之前必须先关闭管道，不然渠道最后一个数据之后还会等待，直到deadlock错误
	close(intChan2)
	for v := range intChan2{
		fmt.Println("v=",v)
	}


}
package main

import (
	"fmt"
)



//write data
func writeData(intChan chan int){
	for i:=0;i<50;i++{
		//放入数据
		intChan <- i
	}
	//写完数据之后可以直接关闭，当读取的时候就不会报错
	close(intChan)
}

func readData(intChan chan int,exitChan chan bool){
	for {
		v, ok := <- intChan	//读取到值返回true,读取不到时(读完的时候)返回false
		if !ok {
			break
		}
		fmt.Println("readData 读到数据=",v)
	}
	//读取完数据后，向exitChan写入数据并立即关闭exitChan
	exitChan <- true
	close(exitChan)
}
func main(){
	//管道时引用类型，不同协程操作的同名管道是同一个管道
	//创建读写管道和退出管道
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)
	go writeData(intChan)
	go readData(intChan,exitChan)
	//读取退出管道
	for {
		_,ok := <- exitChan
		if ok {
			break
		}
	}



}
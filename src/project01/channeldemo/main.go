package main
import (
	"fmt"
)
func main(){
	//创建一个可以存放3个int类型的管道
	var intChan chan int
	//初始化
	intChan = make(chan int,3)

	//打印intChan,得到的是地址，说明是引用类型
	fmt.Println("intChan=",intChan)//intChan= 0xc042070080
	//向管道写入数据
	intChan<- 10
	num :=211
	intChan<- num
	intChan<- 99	
	//输出管道长度和容量,长度指的是目前管道中的元素个数
	//注意：channel的容量是在make的时候就确定的，不能动态增长
	//     当我们向管道写入数据时，不能超过其容量，否则报错deadlock
	fmt.Println("写入后:\nchannel len=",len(intChan),",cap=",cap(intChan))
	//从管道中读取数据，注意:
	//	在没有使用协程的情况下，如果我们的管道数据已经全部取出，
	//	再取就会报错：deadlock
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=",num2)
	fmt.Println("取出后:\nchannel len=",len(intChan),",cap=",cap(intChan))

	var mapChan chan map[string]string
	mapChan = make(chan map[string]string,10)
	m1 := make(map[string]string,20)
	m1["city1"] = "北京"
	m1["city2"] = "上海"

	m2 := make(map[string]string,20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"
	//写入管道
	mapChan <- m1
	mapChan <- m2
	m3 := <- mapChan
	fmt.Println("m3=",m3) 
	
	//定义并初始化存放任意数据类型的管道
	var allChan chan interface{} = make(chan interface{},10)
	allChan <- 10
	allChan <- "jack"
	allChan <- true
	allChan <- 3.8
	allChan <- []int{1,2,3}

	var receiver interface{}
	fmt.Println(len(allChan))
	// for i:=0;i<int(len(allChan));i++{//有疑问：通过len(allChan)数据取不全
	for i:=0;i<5;i++{
		receiver = <- allChan
		fmt.Println(i,receiver)
	}


	
	




}
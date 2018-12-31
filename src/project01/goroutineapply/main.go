package main

import (
	"fmt"

)

func putNum(intChan chan int){
	for i :=2;i <= 80000;i++ {
		intChan <- i
	}
	close(intChan)
}

func PrimeNum(intChan chan int,primeChan chan int,exitChan chan bool) {
	var flag bool
	for {
		num,ok := <-intChan
		if !ok{	//取不到数据
			break
		}
		flag = true		
		for i:=2 ;i< num ;i++{	//判断是否为素数
			if num % i ==0{	//不是素数
				flag =false
				break
			}
		}
		if flag{
			//将这个数放入primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个协程取完数据，退出")
	exitChan <- true
}
func main(){
	intChan := make(chan int,1000)
	primeChan := make(chan int,2000)
	//标识退出的管道
	exitChan := make(chan bool,4)
	//开启协程，向intChan 中写如数据
	go putNum(intChan)

	//开启4个协程，从intChan取出数据，并判断是否为素数
	for i :=0 ;i < 4; i++{
		go PrimeNum(intChan,primeChan,exitChan)
	}	
	//当我们从exitChan 取出数据之后，说明primeChan数据已经放完了
	//此时可以开一个协程关闭primeChan
	go func(){
		//主线程进行判断协程是否都执行完毕，当取出四次数据后，退出循环
		for i := 0;i < 4;i++{
			<- exitChan	//阻塞等待取出数据			
		}
		close(primeChan)	
	}()

	//遍历primeNum,取出结果
	for {
		res,ok := <- primeChan
		if !ok {
			break
		}
		fmt.Println("素数=",res)
	}
	fmt.Println("主线程退出")

}
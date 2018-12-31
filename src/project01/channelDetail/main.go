package main
import (
	"fmt"
)
//只写
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0;i < 10;i++{
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}
//只读
func recv(ch <-chan int, exitChan chan struct{}){
	for {
		v,ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}

func main(){
	var ch chan int
	ch = make(chan int,10)
	exitChan := make(chan struct{},2)
	//将双向管道赋值给只写管道，该管道在该函数中只可写入
	go send(ch,exitChan)
	//将双向管道赋值给只读管道，该管道在该函数中只可读取
	go recv(ch,exitChan)
	var total = 0
	
	for _ = range exitChan{
		total++
		if total ==2{
			break
		}
	}
	fmt.Println("over...")
}
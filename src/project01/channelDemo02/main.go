package main
import (
	"fmt"
)

func main(){

	//使用select解决从管道取数据的阻塞问题
	intChan :=make(chan int,10)
	for i :=0 ;i < 10 ;i++{
		intChan <- i
	}
	stringChan := make(chan string,5)
	for i:= 0;i < 5;i++{
		stringChan <- "hello"+fmt.Sprint("%d",i)
	}
	//传统的方法在遍历管道时，如果不关闭会阻塞而deadlock
	//但是我们在开发过程中，当项目比较大时，往往不确定什么时候关闭管道
	//使用select方式解决
	label:
	for {
		select {
			//注意：这里如果intChan一直没有关闭，不会一直阻塞而deadlock
			//		会自动到下一个case匹配
		case v := <- intChan:
			fmt.Println("从intChan读取的数据：",v)
		case v := <- stringChan:
			fmt.Println("从stringChan读取的数据：",v)
		default:
			//...此处程序员可以加入自己的逻辑
			fmt.Println("数据取完了，不玩了，拜拜")
			break label  
			//推荐使用return退出循环，此处单纯用break只能跳出switch
			//如果实在需要使用break而不想退出main程序，只能配合标签跳出
		}
		

	}

}
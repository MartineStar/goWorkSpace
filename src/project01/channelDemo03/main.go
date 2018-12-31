package main
import (
	"fmt"
	"time"
)
func sayHello(){
	for i := 0;i < 10;i++{
		time.Sleep(time.Second)
		fmt.Println("hello,muzukix")
	}
}

func bridge(funcPtr interface{},args ...interface()){
	// ...业务逻辑
}

func test(){
	//defer recover捕获异常
	defer func(){
		//捕获test抛出的panic
		if err := recover();err !=nil{
			fmt.Println("test() 发生错误：",err)
		}
	}()
	var myMap map[int]string
	//此处并没有初始化就赋值，报错
	//panic: assignment to entry in nil map
	myMap[0] = "golang"
}
func main(){
	go sayHello()
	go test()
	for i :=0;i<10;i++{
		
		fmt.Println("hello,muzukix~~~~")
		time.Sleep(time.Second)
	}
}
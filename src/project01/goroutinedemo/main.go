package main
import (
	"fmt"
	"time"
	"runtime"
	"sync"
)
//思路：
	// 1.编写一个函数，计算各个数的阶乘，并放入到map中
	// 2.我们启动的协程多个，统计的结果将放到map中
	// 3.map因该是一个全局的

var (
	myMap = make(map[int]int,10)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁,sync(synchornized)同步，Mutex互斥
	lock sync.Mutex
)
func test(n int){
	res :=1
	for i :=1;i <= n;i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	//将结果放入myMap
	myMap[n] = res
	//解锁
	lock.Unlock()
}
func main(){
	runtime.GOMAXPROCS(4)
	//这里开启多个协程完成这个任务
	for i:=1; i<=200;i++{
		go test(i)
	}
	time.Sleep(10* time.Second)
	//多个cpu同时往一个map写入东西(报错：concurrent map write)，
	//注意：写不能同时写，必须存在写保护，读可以同时读
	//遍历myMap,输出结果,此时因为10s之后协程可能并没有执行完毕
	//所以在读的过程也需要加锁
	lock.Lock()
	for i,v := range myMap{
		fmt.Printf("map[%v] = %v\n",i,v)
	}	
	lock.Unlock()
}
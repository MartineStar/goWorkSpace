package main
import (
	"fmt"
	"runtime"

)

func main(){
	//查看操作系统的cpu数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=",cpuNum)

	//设置使用cpu的个数
	runtime.GOMAXPROCS(cpuNum -1)
	fmt.Println("ok")
}
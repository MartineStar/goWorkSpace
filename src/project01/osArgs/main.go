package main
import (
	"fmt"
	"os"
)
func main(){
	fmt.Println("命令行的参数有：",len(os.Args),"个,如下：")
	//遍历Args
	for i,v := range os.Args{
		fmt.Printf("Args[%v] = %v\n",i,v)
	}
	//使用Args可以先将go程序打包：go build -o xxx.exe main.go
	//调用：./xxx.exe abc 5343 d:/kkk.txt 哈哈
	//也可以不打包，直接调用go run main.go abc 5343 d:/kkk.txt
	//输出：
	//Args[0] = F:\codernote\golang\goWorkSpace\src\project01\osArgs\test.exe
	// Args[1] = abc
	// Args[2] = 123
	// Args[3] = d:/abc.txt
	// Args[4] = 哈哈
}
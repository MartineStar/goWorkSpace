package main

import (
	"fmt"
	"io/ioutil"
)
func main(){
	//使用ioutil.ReadFile一次性将文件读取到位
	//如果文件很大的情况下，效率很低
	file := "F:/codernote/golang/hello.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err = %v",err)
	}
	//把读取到的内容输入到终端
	fmt.Printf("content=%v",string(content))	//[]byte

	//因为没有显式的Open文件，因此也不需要显式的Close
	//因为，文件的Open和Close被封装到ReadFile函数内部

}
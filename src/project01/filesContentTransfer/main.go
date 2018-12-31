package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	//将F:/abc.txt的文件内容导入到E:/kkk.txt
	file1Path := "F:/abc.txt"
	file2Path := "e:/kkk.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil{
		//读取文件失败
		fmt.Println("读取文件失败：",err)
		return
	}
	err = ioutil.WriteFile(file2Path,data,0666)
	if err != nil {
		fmt.Printf("write file error=%v\n",err)
	}
	//没有显式的打开文件，也不需要显式关闭，打开关闭文件已经封装在方法中
}
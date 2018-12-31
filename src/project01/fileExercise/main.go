package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)
func main(){
	//1.向文件中写入5句"hello world"
	filePath := "F:/abc.txt"
	//以读写的方式打开，并且当写入的时候向原文件追加内容,即：
	//		O_WRONLY 模式和 O_CREATE的组合，用 |
	file,err := os.OpenFile(filePath,os.O_RDWR | os.O_APPEND,0666)
	if err != nil {
		fmt.Println("打开文件失败：",err)
		return
	} else {
		//读取文件
		reader := bufio.NewReader(file)
		for {
			str ,err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Print(str)
		}
	}
	// 及时关闭file句柄
	defer file.Close()

	str := "hello,taylor\r\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i:= 0;i < 5; i++ {
		//WriterString写入是不带\n的，如果需要换行，自行加上
		writer.WriteString(str)
	}
	//因为writer时带缓存的，因此在使用WriterString方法时，
	//需要将调用Flush将缓存的数据写入到文件中
	writer.Flush()
}
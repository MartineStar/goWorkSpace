package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
)
//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount int 	//记录英文个数
	NumCount int	//记录数字个数
	SpaceCount int	//记录空格个数
	OtherCount int 	//记录其他字符个数
}

func main(){
	fileName := "e:/abc.txt"
	file,err := os.Open(fileName)
	if err != nil{
		fmt.Println("open file err=",err)
		return
	}
	defer file.Close()
	//定义CharCount实例,int类型字段默认值为0
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)
	//开始循环读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err ==io.EOF{
			break
		}
		str1 := []rune(str)
		//遍历str，进行统计
		for _,v := range str1{
			//v为一个字符，而case得到的时bool值，会报错
			// switch v{
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <='9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数=%v,数字的个数=%v,空格的个数=%v,其他字符个数=%v",
	count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)


}
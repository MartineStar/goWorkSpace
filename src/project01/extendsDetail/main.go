package main

import (
	"fmt"
)
type Goods struct{
	Name string
}
type TV struct {
	Goods
	int
	n int
}
func main(){
	//当匿名字段为基本数据类型时，该怎么访问
	var tv TV 
	tv.Name = "中央电视台"
	//直接使用结构体名.基本数据类型即可
	tv.int = 20
	tv.n = 100
	fmt.Println(tv)	//{{中央电视台} 20 100}

}
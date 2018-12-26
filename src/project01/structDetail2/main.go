package main

import (
	"fmt"
	"encoding/json"
)
//所以如果想将结构体的字段再json序列化的时候变成小写
//需要用struct tag,即：`json:"name"`
type Monster struct{
	Name string `json:"name"`
	Age int	`json:"age"`
	Skill string	`json:"skill"`
}
func main(){
	//1.创建Monster变量
	monster := Monster{"牛魔王",500,"芭蕉扇"}

	//2.将monster变量序列化成json格式字符串
	//json.Marshal是再json包下面访问monster,
	//如果monster的字段为小写，则访问不到，所以必须大写
	//json.Marshal函数使用了反射
	jsonStr,err :=json.Marshal(monster)
	if err != nil {
		fmt.Println("转化失败",err)
	}
	fmt.Println("jsonStr:",string(jsonStr))
	
	
}
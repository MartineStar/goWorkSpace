package main

import (
	"fmt"
)

func main(){
	//声明map切片
	var monsters []map[string]string
	//初始化切片
	monsters = make([]map[string]string,2)
	//
	if monsters[0] == nil {
		//初始化map
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "玉兔精"
		monsters[0]["age"] = "400"
	}
	
	if monsters[1] == nil {
		//初始化map
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "牛魔王"
		monsters[1]["age"] = "1200"
	}
	//动态增加monsters,通过切片的append函数
	newMonster := map[string]string{
		"name" : "火云邪神",
		"age" : "298",
	}
	monsters = append(monsters,newMonster)
	fmt.Println(monsters)
}
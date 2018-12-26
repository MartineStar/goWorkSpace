package main

import (
	"fmt"
	"sort"
)

type Stu struct{
	Name string
	Age int
	Address string
}
func main(){
	// 演示map的值为struct类型
	students := make(map[string]Stu,10)
	stu1 := Stu{Name:"tom",Age:18,Address:"北京"}
	stu2 := Stu{Name:"mary",Age:28,Address:"上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)
	//遍历
	for k,v :=range students{
		fmt.Printf("学生的编号是：%v\n",k)
		fmt.Printf("学生的名字是：%v\n",v.Name)
		fmt.Printf("学生的年龄是：%v\n",v.Age)
		fmt.Printf("学生的地址是：%v\n",v.Address)
		fmt.Println()
	}




	map1 := map[int]string{
		1 : "唐僧",
		29 : "孙悟空",
		13 : "猪八戒",
		42 : "沙悟净",
		12 : "牛魔王",
	}

	var map2 map[int]map[string]string
	fmt.Println("map2[32]",map2[1])
	if map2[32] == nil{
		fmt.Println("map1[32]为空")
	}
		
	
	//定义一个切片，不同初始化，直接通过append对它进行操作
	var keys []int
	for i,_ := range map1{
		keys = append(keys,i)
	}
	//对切片进行排序
	sort.Ints(keys)
	fmt.Println(keys)
	
	for _, v := range keys{
		fmt.Printf("map1[%v] = %v\n\n",v,map1[v])
	}


}
package main

import (
	"fmt"
)


type Person struct{
	Name string
	Age int
	Scores [5]float64
	// 引用类型存的是地址，当没有为其赋值时，没有指向任何地址，为nil
	ptr *int
	slice []int
	map1 map[string]string

}

//定义结构体
type Cat struct {
	Name string
	Age int
}

func main(){
	//方式1
	var cat1 Cat

	//方式2(常用)
	cat2 := Cat{"tom",2}

	//方式3:new函数返回的是一个指针
	cat3 := new(Cat)
	//(*cat3).Name ="jack"也可以写成cat3.Name="jack"
	//原因：go的设计者为了程序员使用方便，在底层有对cat3.Name="jack"进行处理
	//底层会给cat3加上取值运算 (*cat3).Name = "jack"
	(*cat3).Name = "jack"	//标准写法
	cat3.Age = 23			//也支持这种写法

	//方式4:可以直接给字段赋初值
	var cat4 *Cat = &Cat{"sco",12}
	//因为cat4是一个指针，因此标准的访问字段的方式为：(*cat4).Name="scott"
	//go设计者为了程序员使用方便，也可以cat4.Name = "scott"，原因同上
	(*cat4).Name = "scott"
	cat4.Age = 10




	var p1 Person
	if p1.ptr == nil{
		fmt.Println("p1.ptr=nil,表现为：",p1.ptr)
	}
	if p1.slice == nil{
		fmt.Println("p1.slice=nil,表现为：",p1.slice)
	}
	if p1.map1 == nil{
		fmt.Println("p1.map1=nil,表现为：",p1.map1)
	}
	//如果需要使用结构体中的引用类型map,切片或指针，需要现make分配空间
	p1.slice = make([]int,10)
	p1.slice[0] = 100
	//不同结构体变量是独立的，一个结构体变量的字段改变不会影响另一个
	var monster1 Cat
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1	//结构体为值类型，默认值拷贝
	monster2.Name = "青青"
	fmt.Println(monster1,monster1)//{牛魔王 500} {青青 500}








	//当声明一个结构体变量时，空间就已经分配了，默认值为零值
	//结构体为值类型
	// var cat1 Cat

	// //为结构体变量中元素赋值
	// cat1.Name = "小白"
	// cat1.Age = 3
	// cat1.Color = "白色"
	// cat1.Hobby = "吃鱼"


	// fmt.Println("猫猫信息如下：")
	// fmt.Println("name：",cat1.Name)
	// fmt.Println("猫猫信息如下：",cat1.Age)
	// fmt.Println("猫猫信息如下：",cat1.Color)

	// <


	
}
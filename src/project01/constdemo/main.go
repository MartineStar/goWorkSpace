package main

import (
	"fmt"
)
func main(){
	// var num = 9
	// //1.声明常量的时候必须赋值
	// const tax int = 0
	// //2.常量是不能修改的
	// tax = 10

	// //3.常量只能修饰bool,数值类型(int,float系列),string等基本类型
	// //常量不能修饰变量,因为变量可变，而常量不可变。
	// const b = num / 3

	// //4.所有在编译阶段不能确定的值,都不能赋值给常量，如函数的返回值
	// slice := []int{1,2,3}
	// const c = len(slice)	//len函数的返回值编译不能确定

	// const d = getVal()  	//getVal函数返回值运行时才能确定，报错

	// //5.比较简洁的写法
	// const (
	// 	a = 1
	// 	b = 2
	// )
	// //6.比较专业的写法
	// const (
	// 	a = iota	//表示给a赋值0，b在a的基础上+1,c在b的基础上+1
	// 	b			//1
	// 	c			//2
	// )
	// fmt.Println(a,b,c,d,e,f)

	//面试题
	const (
		//注意：
		//1.iota是以行为递增条件的，是常量计数器，只能在常量表达式中使用
		//2.写了第一个iota,后面的常量写不写都会增1
		//3.iota只在本个常量组起效果，重开一个常量组，iota从0开始，
		//		即遇到一个const就从0开始计数
		//4.iota可以进行表达式的加减乘除操作
		a = iota *10+1		//0
		b 					// 11
		c,d = iota,iota	//2,2	
	)	
	const e = iota	//0
	const f = iota	//0
	fmt.Println(a,b,c,d,e,f)
	
}
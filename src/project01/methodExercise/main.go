package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}
//声明一个方法area和Circle绑定，可以放回面积
func (c Circle) area() float64 {	
	return 3.14 * c.radius * c.radius
}
//为了提高效率，通常我们会将方法和结构体的指针类型绑定
//只需要拷贝地址，而不需要拷贝值，当值比较大时更明显
func (c *Circle) area2() float64 {
	//修改结构体变量的值
	//因为c为指针，因此标准的访问其字段方式为(*c).radius
	//因为go编译器底层做了优化，也支持c.radius调用字段
	c.radius = 10	
	// return 3.14 * (*c).radius * (*c).radius	
	return 3.14 * c.radius * c.radius
}
func main(){
	//创建Circle变量
	var c Circle
	c.radius = 7.0
	//标准的调用方式	
	// res := (&c).area2()
	
	//编译器底层做了优化，(&c).area2()等价于c.area2()
	//编译器底层会自动加上&
	res := c.area2()
	fmt.Println(res)
}
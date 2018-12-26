package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
}

type MethodUtils struct {}

//给*Student实现string方法
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
	return str
}

func main(){
	var stu = Student{
		Name : "tom",
		Age : 20,
	}
	//当进行传值的方式打印学生信息时，会调用默认的fmt.Println
	fmt.Println(stu)	//{tom 20}
	//当进行指针传递时，会调用结构体实现的String()方法
	fmt.Println(&stu)	//Name=[tom] Age=[20]
}
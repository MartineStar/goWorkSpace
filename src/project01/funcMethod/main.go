package main

import (
	"fmt"
)

// //函数值传递
// func test01(p Person) {
// 	fmt.Println("test01()=",p.Name)
// }
// //函数指针传递
// func test02(p *Person) {
// 	fmt.Println("test02()=",(*p).Name)
// }
//方法引用传递
// func (p *Person) test04(){
// 	p.Name = "jack~"
// 	fmt.Println("test04()=",p.Name)
// }

type Person struct {
	Name string
}

//方法值传递
func (p Person) test03(){
	p.Name = "jack"
	fmt.Println("test03()=",p.Name)
}

func main(){
	person :=Person{"tom"}
	//通过地址调用值传递方法test03,调用成功
	//总结：从形式上是传入地址，但本质仍然是值拷贝
	(&person).test03()	//test03()= jack

	//打印person结构体变量的值，没有改变，说明是拷贝了值，而非地址
	//所以底层是将引用类型改变成值类型，不过是语法上支持地址调用值类型方法
	fmt.Println("main()=",person.Name)	//main()= tom




	// //通过值调用形式依然可以调用引用传递方法
	// //编译器底层有将值转换成地址类型，即添加&
	// person.test04()	//test04()= jack~
	// //打印person结构体变量的值，变成了jack~，说明是进行了地址拷贝
	// fmt.Println("main()=",person.Name)	//main()= jack~















	// fmt.Println()
	// test01(person)	//函数：形参为值传递就必须是传值而不能是指针
	// test02(&person)	//函数：形参为引用传递就必须是引用传值而不能是值

}
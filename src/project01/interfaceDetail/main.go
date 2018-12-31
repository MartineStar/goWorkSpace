package main
import  "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	//继承了BInterface,和CInterface
	BInterface
	CInterface
	//AInterface自己也有方法
	test03()
}
//如果需要实现AInterface，就必须将BInterface,CInterface也实现
type Stu struct {}

//方法或函数的方法体内也可以没有内容
func (stu Stu) test01(){}
func (stu Stu) test02(){}
func (stu Stu) test03(){}

type T interface{}

func main(){
	var stu Stu
	var t1 T = stu
	//空接口可以直接在函数中写，空接口也可以看成是一种数据类型
	var t2 interface{} = stu
	//将float64类型赋值给空接口t2,任意类型的值都可以赋值给空接口
	t2 = 9.3

	fmt.Println(t1)
	fmt.Println(t2)
}
package main
import (
	"fmt"
)
//声明一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {}
//让Phone结构体实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

//让相机Camera 实现Usb接口方法
type Camera struct {}
//让Phone结构体实现Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

//循环判断输入的参数类型
func TypeJudge(items ...interface{}){
	//items表示参数名,是一个数组，...表示任意多个，interface{}为空接口类型
	fmt.Println(items)
	for i,x := range items {
		switch x.(type) {	//x.(type)可以判断x的类型
		case bool:
			fmt.Printf("第%v个值为bool类型，值为：%v\n",i,x)
		case string:
			fmt.Printf("第%v个值为string类型，值为：%v\n",i,x)
		case float64,float32:
			fmt.Printf("第%v个值为浮点数类型，值为：%v\n",i,x)
		case int,int32,int64:
			fmt.Printf("第%v个值为整数类型，值为：%v\n",i,x)
		default:
			fmt.Println("无法判断类型")
		}
	}
} 

func main(){
	//多态数组
	var usbArr [3] Usb
	fmt.Println("赋值前：",usbArr)
	//赋值前： [<nil> <nil> <nil>]
	usbArr[0] = Phone{}
	usbArr[1] = Camera{}
	usbArr[2] = Phone{}
	//赋值后： [{} {} {}]
	fmt.Println("赋值后：",usbArr)

	//带检测的类型断言
	var x interface{}
	var b float32 = 1.4
	x = b

	if y,ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型：%T,y的值：%v",y,y)	
	}else {
		fmt.Println("convert fail")
	}
	
	fmt.Println("继续执行...")	
	//y的类型：float32,y的值：1.4


	// //创建结构体变量
	// computer := Computer{}
	// phone := Phone{}
	// camera := Camera{}

	// //关键
	// computer.Working(phone)
	// computer.Working(camera)
	slice := []int{1,2,3}
	TypeJudge(1,2.0,3,4,5,6,slice,"helo","gogo")

	

}
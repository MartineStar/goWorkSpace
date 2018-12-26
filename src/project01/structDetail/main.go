package main

import (
	"fmt"
)

type Point struct{
	x ,y int
}
type Rect struct{
	leftUp ,rightUp Point
}

type Rect2 struct{
	leftUp ,rightUp *Point
}
func main(){
	//r2有两个*Point类型，这两个*Point类型的本身地址也是连续的
	//但是他们指向的地址不一定连续
	r2 := Rect2{&Point{10,20},&Point{30,40}}
	//打印本身地址
	fmt.Printf("r2.leftUp 本身的地址=%p\n",&r2.leftUp)
	fmt.Printf("r2.rightUp 本身的地址=%p\n",&r2.rightUp)
	//打印指向地址
	fmt.Printf("r2.leftUp 指向的地址=%p\n",r2.leftUp)
	fmt.Printf("r2.rightUp 指向的地址=%p\n",r2.rightUp)
	//输出：
	// r2.leftUp 本身的地址=0xc04203e1b0
	// r2.rightUp 本身的地址=0xc04203e1b8
	// r2.leftUp 指向的地址=0xc04204a070
	// r2.rightUp 指向的地址=0xc04204a080

	r1 := Rect{Point{1,2},Point{3,4}}
	fmt.Printf("r1.leftUp.x 的地址=%p\n",&r1.leftUp.x)
	fmt.Printf("r1.leftUp.y 的地址=%p\n",&r1.leftUp.y)
	fmt.Printf("r1.rightUp.x 的地址=%p\n",&r1.rightUp.x)
	fmt.Printf("r1.rightUp.y 的地址=%p\n",&r1.rightUp.y)
	// 输出：相邻地址相差8个字节，即int的字节数
	// r1.leftUp.x 的地址=0xc0420480a0
	// r1.leftUp.y 的地址=0xc0420480a8
	// r1.rightUp.x 的地址=0xc0420480b0
	// r1.rightUp.y 的地址=0xc0420480b8


}
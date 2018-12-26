package main
import (
	"fmt"
	"project01/factory/model"
)
func main(){
	// 通过工厂模式访问model包中的student结构体
	var stu = model.NewStudent("bibo",99.8)
	//得到的是地址即 &{...}
	fmt.Printf("stu类型为：%T,stu的值为：%v\n",stu,stu)
	
	fmt.Println("name=",stu.Name,"score=",stu.GetScore())
	//输出：name= bibo score= 99.8

}
package main

import (
	"fmt"
	"math/rand"
	 "time"
)
func main(){

	var intArr3 [5]int
	//rand的seed是用来初始化生成器到一个确定的状态，
	//如果不设定，就会使用默认值，此时seed是固定的，
	//产生的随机数也是固定的，固定的随机数就是[81 87 47 59 81]
	rand.Seed(time.Now().UnixNano())
	for i := 0;i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)


	// arr := [3]int{1,2,3}
	// test02(&arr)
	// fmt.Println(arr)

	// var hens [6]float64
	// // hens以及该被定义了，所以hens[0]:=3.0是错误的写法，hen[3]=3.0才正确
	// hens[0] = 3.0
	// hens[1] = 5.0
	// hens[2] = 3.4
	// hens[3] = 1.0
	// hens[4] = 2.0
	// hens[5] = 50.0
	// totalWeight := 0.0
	// for i :=0;i < len(hens);i++ {
	// 	totalWeight += hens[i]
	// }
	// //------------------------------------------------------------
	// // 注意：
	// // 	float64 / int 语法错误，必须同种类型才能运算
	// // 	如果此处是 totalWeight / 6则不会报错，因为6是一个常量值，类型是没有固定的，
	// // 	当你把一个常量值给一个变量的时候，这个变量的类型就固定了		
	// //--------------------------------------------------------------
	// averWeight := totalWeight / float64(len(hens))
	// fmt.Printf("鸡的总体重为%.2f,平均体重为%.2f",totalWeight,averWeight)
	
	//演示for-range
	// var heroes [3]string = [3]string{"松江","吴用","卢俊义"}
	// for index,value :=range heroes {
	// 	fmt.Printf("index = %v,value = %v\n",index,value)
	// }
	
	

	

}
package main

import (
	"fmt"
)

func main(){
	
	// 切片的拷贝
	var slice4 []int = []int{1,2,34,56,7}
	var slice5 = make([]int,10)
	//将后面的切片拷贝到前面的切片
	copy(slice5,slice4)
	fmt.Println("slice4",slice4)
	//slice4 [1 2 34 56 7]
	fmt.Println("slice5",slice5)
	
	var chars []byte = []byte{'a','b','c'}
	str := string(chars)
	fmt.Println("str",str)
	//slice5 [1 2 34 56 7 0 0 0 0 0]

	

	// // 用append内置函数，可以对切片进行动态追加
	// var slice3 []int = []int{11,22,33}
	// // 通过append直接给slice追加具体的元素
	// slice = append(slice,444,55,6)
	// fmt.Printf("slice3=%v\n",slice3)

	// // 通过append直接给slice追加另外切片
	// // 追加的切片后面需要带...,这是固定写法
	// slice3 = append(slice3,slice...)
	// fmt.Printf("slice3=%v\n",slice3)

	// //常规for循环进行切片遍历
	// var arr [5]int = [...]int{1,2,3,4,5}
	// slice := arr[1:4]
	// for i :=0;i< len(slice);i++ {
	// 	fmt.Printf("slice[%d] = %v\n",i,slice[i])
	// }

	// fmt.Println("==========================")
	// //for range进行切片遍历
	// for i ,v := range slice {
	// 	fmt.Printf("i = %v,v = %v\n",i,v)
	// }

	
	// //直接指定具体数组，原理类似make
	// var slice []string = []string{"tom","kink","jacky"}
	// fmt.Printf("slice = %v\n",slice)
	// fmt.Printf("slice len = %v\n",len(slice))
	// //当cap没有指定时，和len相等
	// fmt.Printf("slice cap = %v\n",cap(slice))


	// var slice []float64 = make([]float64,5,10)
	// slice[1] = 10
	// slice[3] = 99
	
	// fmt.Printf("slice = %v\n",slice)
	// fmt.Printf("slice len = %v\n",len(slice))
	// fmt.Printf("slice cap = %v\n",cap(slice))

	//演示切片的基本使用
	// var intArr [5]int = [...]int{1,22,34,66,99}

	// //声明一个切片,引用的是intArr的第2到第4个元素，不包含4
	// slice := intArr[1:3]

	// fmt.Printf("type of slice=%T\n",slice)
	// fmt.Printf("value of slice=%v\n",slice)
	// fmt.Printf("len of slice=%d\n",len(slice))
	// fmt.Printf("capacity of slice=%d",cap(slice))
	// fmt.Printf("slice[0]=%v",slice[0])
	// 输出
	// type of slice=[]int
	// value of slice=[22 33]
	// len of slice=2
	// capacity of slice=5

 }

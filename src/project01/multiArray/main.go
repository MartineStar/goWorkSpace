package main

import "fmt"
func main(){
	// // 二维数组其实就是一维数组里面的元素类型为一维数组
	// var arr [4][6]int
	// arr[1][2] = 1
	// arr[2][3] = 2
	// arr[3][4] = 3
	// length := len(arr)
	// inner_length := len(arr[0])
	// for i := 0; i < length;i++ {
	// 	for j :=0;j < inner_length;j++ {
	// 		fmt.Print(arr[i][j]," ")
	// 	}
	// 	fmt.Println()
	// }
	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)
	fmt.Printf("arr2的地址是：%p\n",&arr2)
	fmt.Printf("arr[0]的地址是%p,arr[0][0]的地址是%p\n",&arr2[0],&arr2[0][0])
	fmt.Println()
	fmt.Printf("arr[1]的地址是%p,arr[1][0]的地址是%p\n\n",&arr2[1],&arr2[1][0])


}

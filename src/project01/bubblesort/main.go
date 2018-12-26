package main

import "fmt"
//指针就是在原来的数据类型前加*
func Bubble(arr *[5]int) {
	fmt.Println("排序之前：",arr)
	for i :=1 ;i < len(arr);i++ {
		for j :=0;j < len(arr)-i;j++{
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j],(*arr)[j+1] = (*arr)[j+1],(*arr)[j]
			}
		}
	}
	fmt.Println("排序之后：",arr)
}

func main(){
	var bubbleArr [5]int = [5]int{23,42,12,2,5}
	Bubble(&bubbleArr)


}
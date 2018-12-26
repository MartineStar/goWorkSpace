package main

import "fmt"

//此处对数组只是进行查找而不会改变原数组，所以：
//进行引用拷贝而不进行值拷贝，效率更高，当数据量很大的时候更明显
func BinaryFind(arr *[6]int,leftIndex int,rightIndex int,findVal int) {
	//递归条件：判断leftIndex 是否大于rightIndex
	if leftIndex > rightIndex{
		fmt.Println("找不到")
		return
	}	
	//先找到中间的下标
	middle := (leftIndex + rightIndex) /2
	if (*arr)[middle] > findVal {
		BinaryFind(arr,leftIndex,middle-1,findVal)
	}else if (*arr)[middle] < findVal{
		BinaryFind(arr,middle+1,rightIndex,findVal)
	}else{
		fmt.Println("找到了,下标为",middle)
	}
}
func main(){
	//二分查找的数组必须是已经排序了的
	arr := [6]int{3,4,34,109,208,299}
	var findVal int
	fmt.Println("请输入需要查找的数：")
	fmt.Scanln(&findVal)
	BinaryFind(&arr,0,len(arr)-1,findVal)
}
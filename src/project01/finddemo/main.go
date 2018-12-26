package main

import "fmt"

//顺序查找方式1
func ArrFind(arr [5]string) {
	var name string
	fmt.Println("请输入需要查找的姓名：")
	fmt.Scanln(&name)
	for i := 0 ;i < len(arr);i++ {
		if name == arr[i] {
			fmt.Printf("找到%v ,下标为%v",name,i)
			break
		} else if i == len(arr) -1{
			fmt.Printf("未找到%v",name)
		}
	}
}
//顺序查找方式2(推荐)
func ArrFind2(arr [5]string) {
	var name string
	fmt.Println("请输入需要查找的姓名：")
	fmt.Scanln(&name)
	index := -1
	for i := 0 ;i < len(arr);i++ {
		if name == arr[i] {
			index = i
		} 
	}
	if index == -1 {
		fmt.Println("未找到",name)
	}else {
		fmt.Printf("找到 %v~,下标为 %v~",name,index)
	}
}
func main(){
	arr := [...]string{"古天乐","马云","刘强东","刘亦菲","杨幂"}
	ArrFind2(arr)
}
package main

import "fmt"
func main(){
	var arr3 = [...][3]int{{1,2,3},{4,5,6}}
	
	//for range遍历
	for i,v := range arr3{
		for index,value := range v{
			fmt.Printf("arr3[%v][%v] = %v\t",i,index,value)
		}
		fmt.Println()
	}

		//for循环遍历
		for i:= 0;i < len(arr3);i++{
			for j:=0;j < len(arr3[i]);j++{
				fmt.Printf("%v\t",arr3[i][j])
			}
			fmt.Println()
		}
		fmt.Println("--------------------------")

}
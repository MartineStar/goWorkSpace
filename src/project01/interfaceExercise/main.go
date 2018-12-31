package main
import (
	"fmt"
	"sort"
	"math/rand"
)
//声明Hero结构体
type Hero struct{
	Name string
	Age int 
}
//声明一个Hero结构体切片类型
type HeroSlice []Hero

//实现Sort的interface接口
func (hs HeroSlice) Len() int{
	return len(hs)
} 
//Less方法就是决定使用什么标准进行排序
func (hs HeroSlice) Less(i,j int) bool{
	return hs[i].Age < hs[j].Age
} 

//Swap交换
func (hs HeroSlice) Swap(i,j int){
	hs[i],hs[j] = hs[j],hs[i]
} 

func main(){
	var heroes HeroSlice
	for i := 0; i < 10; i++{
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d的",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heroes = append(heroes,hero)
	}

	//排序前的顺序
	for _,v := range heroes {
		fmt.Println(v)
	}
	fmt.Println("排序后")
	//调用sort.Sort
	sort.Sort(heroes)
	for _,v := range heroes {
		fmt.Println(v)
	}

}
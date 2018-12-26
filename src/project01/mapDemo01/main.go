package main

import (
	"fmt"
)

func main(){
	//方式一：先声明，再make,后使用
	var a map[string]string	
	a = make(map[string]string,10)
	fmt.Println(a)	
	a["nol"] = "宋江"
	a["nol2"] = "无用"	
	a["nol"] = "武松"
	fmt.Println(a)

	//方式三:声明并直接初始化
	heroes := map[string]string{
		"hero1" : "宋江",
		"hero2" : "卢俊义",
		"hero3" : "吴用",	//最后的键值对逗号也要带上
	}
	heroes["hero4"] = "武松"
	fmt.Println(heroes)

	//方式二：声明并make,make可以不指定长度
	cities :=make(map[string]string)
	cities["city1"] = "北京"
	cities["city2"] = "上海"
	cities["city3"] = "天津"
	fmt.Println(cities)

	// 简单map遍历
	for k,v := range cities{
		fmt.Printf("k = %v,v = %v",k,v)
	}

	//更新，添加操作
	cities["city5"] = "USA"
	cities["city2"] = "shanghai"
	fmt.Println("更新后cities,",cities)	

	// delete操作
	delete(cities,"city1")
	fmt.Println("删除后cities,",cities)

	// map元素查找操作
	val, ok := cities["cityl"]
	if ok {
		fmt.Println("有city1 key这个值",val)
	}else{
		fmt.Println("没有city1 key这个值")
	}






	//案例,map里面嵌套map,注意一旦使用map就要初始化make
	studentMap := make(map[string]map[string]string)
	
	//里面的map也需要先make初始化，才能使用
	studentMap["stu01"] = make(map[string]string,3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "广州白云区新市墟"

	studentMap["stu02"] = make(map[string]string,3)
	studentMap["stu02"]["name"] = "muzuki"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["address"] = "北京市海淀区百度创意园"

	studentMap["stu03"] = make(map[string]string,3)
	studentMap["stu03"]["name"] = "amy"
	studentMap["stu03"]["sex"] = "女"
	studentMap["stu03"]["address"] = "武汉市武昌区蔡甸"
	fmt.Println(studentMap)

	//复杂结构的map遍历
	for k1,v1 := range studentMap {
		fmt.Printf("k1=%v\n",k1)
		for k2,v2 :=range v1 {
			fmt.Printf("\t k2=%v,v2=%v\n",k2,v2)
		}
	}


}
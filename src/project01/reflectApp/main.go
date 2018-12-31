package main
import (
	"fmt"
	"reflect"
)
//定义一个Monster结构体
type Monster struct{
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32
	Sex string
}
//方法，显示s的值
func (s Monster) Print(){
	fmt.Println("---start----")
	fmt.Println(s)
	fmt.Println("----end-----")
}

func TestStruct(a interface{}){
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	//如果传入的不是struct,就退出
	//因为是在reflect包下定义的常量，所以是reflect.Struct
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取到该结构体的几个字段
	num := val.NumField()
	//遍历结构体的所有字段
	for i:= 0;i < num ;i++{
		//通过reflect.Value的Field方法可以获取到具体的字段值
		fmt.Printf("field %d:值为=%v\n",i,val.Field(i))
		//通过reflect.Type的Field方法可以获取到tag信息,两者的Field方法很不一样
		
		//Get方法传入的参数为标签的键名，即`json:"name"`中的json,当然这个名字可以自定义
		//在json序列化时，键名是固定的，因为其内部源码有将键名固定为json
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段存在tag标签名字为json的，就打印出它的值
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n",i,tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Println("struct 有",numOfMethod,"个方法")

	//调用指定的方法
	//Val.Method(1) 表示获取到第二个方法
	//方法默认是按照名字进行排序，即字符串的ascii码排序规则
	//Call(nil)表示调用这个方法，并不传入任何参数
	val.Method(1).Call(nil)	

	//调用结构体的第1个方法，需要传入参数
	var params []reflect.Value //声明reflect.Value类型切片
	params = append(params,reflect.ValueOf(10),reflect.ValueOf(10))
	//Call方法要求传入的参数必须时reflect.Value切片类型
	//	返回的结果也是reflect.Value切片类型,存放返回来的值(reflect.Value类型)
	res := val.Method(0).Call(params)
	fmt.Println("res=",res[0].Int())

}

//方法2：返回两个数的和
func (s Monster) GetSum(n1,n2 int) int{
	return n1 + n2
}
//方法3：接收4个值，给s赋值
func (s Monster) Set(name string,age int, score float32,sex string){
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func main(){
	var a Monster = Monster{
		Name : "黄鼠狼精",
		Age : 400,
		Score : 34.6,
		Sex : "male",
	}
	TestStruct(&a)
}
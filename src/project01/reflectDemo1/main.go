package main
import (
	"fmt"
	"reflect"
)
//通过反射，修改num int的值，修改student的值
func reflect01(b interface{}) {
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rval Kind=",rVal.Kind()) 
	//输出：rval Kind= ptr,指针类型

	//3.SetInt方法的对象必须不能是指针类型，需要经过Elem方法转化
	// rVal.SetInt(20)	//报错
	rVal.Elem().SetInt(20)	//正确的写法
}

func main(){
	var num int = 10
	//注意值类型传递地址才能改变原来的值
	reflect01(&num)
	fmt.Println("改变后num的值为：",num)	//20
}
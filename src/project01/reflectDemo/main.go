package main
import (
	"fmt"
	"reflect"
)


type Student struct{
	Name string
	Age int
}
//对结构体的反射
func reflectTest(b interface{}) {
	//1.获取到reflect.Value
	rVal := reflect.ValueOf(b)

	//2.将reflect.Value 转换成interface{}
	iVal := rVal.Interface()
	fmt.Printf("iVal=%v,iVal类型=%T\n",iVal,iVal)
	// iVal={tom 12},iVal类型=main.Student

	//3.获取变量对应的Kind
	//(1) 通过reflect.Value.Kind() 获取Kind
	fmt.Println("(reflect Value) kind=",rVal.Kind())

	//(2) 通过reflect.Type.Kind() 获取Kind
	rType := reflect.TypeOf(b)
	fmt.Println("(reflect Type) kind=",rType.Kind())

	//4.将结构体中的Name打印出来
	// fmt.Println("Name = ",iVal.Name)
	//编译报错：
	//因为反射的本质是在运行的时候才能确定某个值的具体类型，
	//	但是在编译阶段是确定不了的，编译器会报错，必须类型断言
	//	类型断言可以通过switch进行更健壮的判断
	student,ok := iVal.(Student)
	if !ok{
		fmt.Println("it is not ok for type Student")
		return 
	}
	fmt.Println("Name = ",student.Name)
}

func main(){
	stu := Student{"tom",12}
	reflectTest(stu)
		
}
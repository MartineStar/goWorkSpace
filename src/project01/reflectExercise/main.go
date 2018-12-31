package main

import (
	"fmt"
	"reflect"
)
func main(){
	var str = "hello world"

	//此处对于值数据类型来说必须传地址，否则执行SetXxx时，报错
	iVal := reflect.ValueOf(str)	

	fmt.Printf("%T",iVal)

	//reflect.Valuepanic:reflect.Value.SetString using unaddressable value
	iVal.SetString("tom")

	//reflect.Valuepanic: call of reflect.Value.Elem on string Value
	iVal.Elem().SetString("tom")
}

package monster
import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)
type Monster struct {
	Name string
	Age int
	Skill string
}

//给Monster绑定一个方法Store，可以将Monster变量(对象),序列化并保存
func (this *Monster) Store() bool{
	//先序列化
	data ,err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =",err)
		return false
	}
	// 保存文件
	filePath :="f:/monster.ser"
	err = ioutil.WriteFile(filePath,data,0666)
	if err != nil {
		fmt.Println("write file err=",err)
		return false
	}
	return true
}

//给Monster绑定ReStore,可以将一个序列化的Monster从文件中读取，
//并反序列化为Monster对象，
func (this *Monster) ReStore() bool{
	//从文件中读取序列化的字符串
	filePath := "f:/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err !=nil{
		fmt.Println("reade file err=",err)
		return false
	}
	//shi用读取到的data []byte,反序列化
	json.Unmarshal(data,this)
	if err != nil{
		fmt.Println("unmarshal fail err=",err)
		return false
	}
	return true
}
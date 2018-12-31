package main
import (
	"fmt"
)
//将Pupil和Graduate共有的属性抽象出来，放在Student结构体中
type Student struct {
	Name string
	Age int
	Score int
}
//将Pupil和Graduate共有的方法也绑定到Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v, 年龄=%v, 成绩=%v",stu.Name,stu.Age,stu.Score)
}

//小学生
type Pupil struct {
	Student	//嵌入Student匿名结构体(没有名称，只有类型)
}
//大学生
type Graduate struct {
	Student
}

//Pupil结构体特有的方法，保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试......")
}
//Graduate结构体特有的方法，保留
func (g *Graduate) testing() {
	fmt.Println("大学生正在考试......")
}

func (stu *Student) SetScore(score int) {
	//业务判断
	stu.Score = score
}



func main(){

	//对结构体潜入了匿名结构体之后，用法会发生改变
	pupil := &Pupil{}
	pupil.Student.Name = "tom"	
	//上述语句也可以写成：pupil.Name = "tom"
	//底层执行原理：就近原则
		//1.编译器会先看pupil对应的类型Pupil有没有Name字段，
		//	如果有，则直接调用Pupil类型的Name字段
		//  如果没有，则去嵌入的匿名结构体Student中看有没有Name字段,
			//	如果有则调用，没有则继续查找下一个匿名结构体，
				//如果第一层的匿名结构体都没有找到，就逐层往下查找,依旧没有就报错
				
	pupil.Student.Age = 8	//此处等价于 pupil.Age = 8
	pupil.testing()	
	pupil.Student.ShowInfo()	//此处等价于pupil.ShowInfo()
}
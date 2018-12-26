package model
//定义一个私有结构体，包外访问不到
type student struct {
	Name string
	score float64
}

//工厂模式解决包外访问不到student结构体的问题
func NewStudent(n string,s float64) *student {
	//直接返回student结构体的地址给调用者
	return &student{
		Name : n,
		score : s,
	}
}

//如果score字段首字母小写，则在其他包不能直接访问，我们可以提供一个方法
func (s *student) GetScore() float64{
	return s.score	//包内可以访问score
}
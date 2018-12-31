package monster		//同一个包下
import (
	"testing"	//自动化测试包
)
//测试用例，测试Store方法
func TestStore(t *testing.T) {
	//先创建一个Monster实例
	monster :=Monster{
		Name :"红孩儿",
		Age : 10,
		Skill : "三味真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望为=%v,实际为=%v",true,res)
	}
	t.Logf("monster.Store() 测试成功")
}

//测试用例，测试ReStore方法
func TestReStore(t *testing.T) {
	//先创建一个Monster实例
	var monster Monster

	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误，希望为=%v,实际为=%v",true,res)
	}
	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误，希望monster.Name=%v,实际=%v",
		"红孩儿",monster.Name)
	}
	t.Logf("monster.Store() 测试成功")
}
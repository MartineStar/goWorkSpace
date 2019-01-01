package main
import (
	"fmt"
	 "github.com/garyburd/redigo/redis"

)
type Monster struct{
	name string
	age int
	skill string
}
func main(){
	//连接redis
	conn,err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println("connect to redis failed,err =",err)
		return
	}else{
		fmt.Println("connect to redis successful,conn=",conn)
	}
	//声明一个monster变量
	monster := Monster{}

	//循环输入monster信息并逐个写入redis	
	for monsterNo :=1;monsterNo<=3;monsterNo++{			
		fmt.Printf("请输入monster%v的信息(name,age,skill),并以空格隔开：",monsterNo)
		fmt.Scanf("%s %d %s",&monster.name,&monster.age,&monster.skill)
		fmt.Println("monster[name]:",monster.name,"monster[age]:",monster.age,"monster[skill]:",monster.skill)

		monsterName := "monster"+ fmt.Sprintf("%d",monsterNo)
		fmt.Println(monsterName) 
		_,err = conn.Do("HMSet",monsterName,"name","jhon","age",29)

		if err != nil{
			fmt.Println("hmset key-value failed,err=",err)
			return
		} 
	}
}

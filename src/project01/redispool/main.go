package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
//定义全局的pool
var pool *redis.Pool

//启动程序时，初始化连接池
//用的较多的连接池初始化配置
func init(){
	pool = &redis.Pool {
		MaxIdle:8,
		MaxActive:0,
		IdleTimeout:300,
		Dial:func() (redis.Conn,error) {
			return redis.Dial("tcp","localhost:6379")
		},		
	}
}

func main(){
	//先从pool取出一个链接
	conn := pool.Get()
	//用完关闭链接，链接又回到空闲状态
	defer conn.Close()
	_,err := conn.Do("Set","name2","小白")
	if err != nil{
		fmt.Println("conn.Do set failed,err=",err)
	}
	//取出
	res,err := redis.String(conn.Do("Get","name2"))
	if err != nil{
		fmt.Println("conn.Do get failed,err=",err)
		return
	}
	fmt.Println("res=",res)

	//一旦关闭链接池，再取链接,就会指向一个无效的地址，
	//当执行Do的时候就会报错get on closed pool
	pool.Close()
}
package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	
)
func main(){
	//通过go向redis写入数据和读取数据
	//连接redis服务器,c表示连接
	conn,err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println("connect to redis failed,err =",err)
		return
	}else{
		fmt.Println("connect to redis successful,conn=",conn)
	}
	
	//及时关闭连接
	defer conn.Close()
	//Do表示执行指令
	_,err = conn.Do("Set","key1",998)
	if err != nil{
		fmt.Println("set key-value failed,err=",err)
		return
	}
	//将get获取的key1的值断言成Int
	r,err := redis.Int(conn.Do("Get","key1"))
	if err != nil{
		fmt.Println("get key1 failed,err=",err)
		return
	}
	fmt.Println(r)

}
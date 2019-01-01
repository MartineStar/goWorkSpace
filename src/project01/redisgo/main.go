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

	//批量set key-value,即MSet
	_,err = conn.Do("lpush","heroList","nol1:宋江",30,"no2:武松",30)
	if err != nil{
		fmt.Println("mset key-value failed,err=",err)
		return
	} 

	//批量get key-value,即MGet
	r,err := redis.String(conn.Do("rpop","heroList"))
	if err != nil{
		fmt.Println("mget key-value failed,err=",err)
		return
	}
	fmt.Println(r)
	// //遍历mget获取到的切片
	// for _,v := range r{
	// 	fmt.Println(v)
	// }

	// //设置name 10s过期
	// _,err = conn.Do("expire","name",10)
	// if err != nil {
	// 	fmt.Println("set expire error,err =",err)
	// 	return
	// }
	// //及时关闭连接
	// defer conn.Close()
	// //HMset操作
	// _,err = conn.Do("HMSet","user02","name","jhon","age",29)
	// if err != nil{
	// 	fmt.Println("hmset key-value failed,err=",err)
	// 	return
	// } 

	// //将hmget获取的值断言成Strings,因为获取到的是一个切片
	// r,err := redis.Strings(conn.Do("HMGet","user02","name","age"))
	// if err != nil{
	// 	fmt.Println("hget key1 failed,err=",err)
	// 	return
	// }

	// //遍历取出r中的元素
	// for i,v := range r{
	// 	fmt.Printf("r[%v]=%v\n",i,v)
	// }
}
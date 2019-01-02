package main
import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//定义一个全局pool
var pool *redis.Pool

//初始化redis连接池函数,因为程序可能存在很多的初始化程序
//所以分开写函数，而不是直接放到一个init函数中
func initPool(address string,maxIdle,maxActive int,idleTime time.Duration){
	pool = &redis.Pool {
		MaxIdle:maxIdle,	//最大空闲链接数
		MaxActive:maxActive,	//和数据库的最多链接数，0表示没有限制
		IdleTimeout:idleTime,//最长空闲时间
		Dial:func() (redis.Conn,error) {//初始化链接的代码
			return redis.Dial("tcp",address)
		},		
	}
}
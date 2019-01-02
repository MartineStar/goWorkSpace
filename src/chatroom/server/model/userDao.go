package model

import (
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"chatroom/common/message"
)

//在服务器启动后，就初始化一个userDao实例
//把它做成全局的变量，在需要和reids交互时，就直接使用,提升效率
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}
//链接池程序一般放在程序开始初始化的时候，即服务端的main函数
//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool:pool,
	}
	return
}
//UserDao应该提供哪些方法：
//1.根据用户id返回一个User的实例+err
func (this *UserDao) getUserById(conn redis.Conn,id int) (user *User,err error){

	//通过id去 redis查询对应的用户
	res, err := redis.String(conn.Do("HGet","users",id))
	if err != nil {
		if err == redis.ErrNil {//表示在user 哈希中，没有找到对应id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	
	//这里我们需要把res反序列化成User实例
	err = json.Unmarshal([]byte(res),&user)	//次数需要传递的时引用类型&user
	if err != nil {
		fmt.Println("json.Unmarshal failed,error=",err)
		return
	}
	return

}

//注册
func (this *UserDao) Register(user *message.User) (err error) {
	//先从UserDao 的链接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	_,err = this.getUserById(conn,user.UserId)
	if err != nil {
		return 
	}

	//此时，说明id在redis不存在
	data ,err := json.Marshal(user)
	if err != nil {
		return
	}

	//入库 ,此处的"users"最好不要写死了，写成一个全局的常量最好
	_,err = conn.Do("HSet","users",user.UserId,string(data))
	if err != nil {
		fmt.Println("保存注册用户错误,error=",err)
		return
	}
	return
}

//完成登陆的校验
//1.Login完成对用户的验证
//2.如果用户的id和pwd都正确，则返回个user的实例
//3.如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int,userPwd string) (user *User,err error){
	//先从UserDao 的链接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil {
		return 
	}
	//这是证明这个用户在数据库是存在的
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
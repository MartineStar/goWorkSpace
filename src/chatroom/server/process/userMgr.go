package process
import (
	"fmt"
)

//因为UserMgr实例在服务器端有且只有一个
//并且很多地方都会使用到，因此将其定义为一个全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init(){
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess,1024),
	}
}

//完成对onlineUsers添加,无则添加，有则修改
func (this *UserMgr) AddOnlineUser(userProcess *UserProcess){
	this.onlineUsers[userProcess.UserId] = userProcess
}

//删除
func (this *UserMgr) DelOnlineUser(userId int){
	delete(this.onlineUsers,userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess{ 
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (userProcess *UserProcess,err error){
	userProcess,ok := this.onlineUsers[userId]
	if !ok {//此用户 不在线或不存在该账号
		err = fmt.Errorf("用户%d 不在在",userId)	//格式化一个错误
		return
	}
	return
}
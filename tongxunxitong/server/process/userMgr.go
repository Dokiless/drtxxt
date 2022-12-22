package process2

import "fmt"

//定义一个全局变量userMgr

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对userMgr的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUser的添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// 返回当前在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据用户ID返回对应状态
func (this *UserMgr) GetOnlineUserById(userid int) (up *UserProcess, err error) {
	//从map中取一个值
	up, ok := this.onlineUsers[userid]
	if !ok { //说明用户不在线
		err = fmt.Errorf("该用户当前离线", userid)
		return
	}
	return
}

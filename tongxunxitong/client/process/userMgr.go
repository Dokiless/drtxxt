package process

import (
	"fmt"
	"go_code/project01/main/tongxunxitong/client/modle"
	message "go_code/project01/main/tongxunxitong/common/massage"
)

// 客户端维护用的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var curUser modle.CurUser

// 在客户端显示当前在线用户
func outputOnlineUsers() {
	fmt.Println("当前在线用户：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)

	}
}

// 编写一个方法,处理返回的NotifyUserStatusMes
func updateUserStatusMes(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //说明用户不在线
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UseStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUsers()
}

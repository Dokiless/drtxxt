package message

// 定义一个用户结构体
type User struct {
	UserId    int    `json:"userId"`
	UserPwd   string `json:"userPwd"`
	UserName  string `json:"userName"`
	UseStatus int    `json:"useStatus"` //用户在线状态
}

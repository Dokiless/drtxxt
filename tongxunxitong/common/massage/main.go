package message

import "go_code/project01/main/tongxunxitong/server/model"

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 定义用于表示用户在线状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatys
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

// 用户登录消息类型
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code    int    `json:"code"` //反馈状态码，状态码为200定义为表示登陆成功，状态码为500定义为表示登陆失败
	UserIds []int  //保存用户id切片
	Error   string `json:"error"` //错误信息
}

// 用户注册消息类型
type RegisterMes struct {
	User model.User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"` //反馈状态码，状态码为200定义为表示注册成功，状态码为400定义为表示该用户已被注册
	Error string `json:"error"`
}

// 定义一个配合服务器端推送用户在线状态的消息结构体
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

// 定义一个消息结构体
type SmsMes struct {
	Content string `json:"content"` //消息内容
	User           //匿名结构体,继承
}

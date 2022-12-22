package process2

import (
	"encoding/json"
	"fmt"
	message "go_code/project01/main/tongxunxitong/common/massage"
	"go_code/project01/main/tongxunxitong/server/model"
	"go_code/project01/main/tongxunxitong/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表明是哪个用户的链接
	UserId int
}

// 编写一个通知其他用户有人上线的方法
// userid通知其他用户自己上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onlineUsers一个个发送上线通知
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		//开始通知，另外写一个方法
		up.NotifyMeOnline(userId)
	}
}

// 推送消息
func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//发送，创建Transfer实例
	tf := &utils.Tsansfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err", err)
	}
}

// 编写一个ServerProcessRegisterMes函数,专门处理注册消息
func (this *UserProcess) ServerProcessRegisterMes(mes message.Message) (err error) {
	//从mes中取出mes.Data,再将其反序列化
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//2.再声明一个LoginResMes并完成赋值
	var registerResMes message.RegisterResMes
	//需要到redis数据库完成注册
	//1.使用model.MyUserDao到redis进行验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXITS {
			registerResMes.Code = 400
			registerResMes.Error = model.ERROR_USER_EXITS.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "注册时发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	//3.将loginResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.将resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//6.发送消息,将代码封装到writePkg函数中
	tf := &utils.Tsansfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 编写一个ServerProcessloginMes函数,专门处理登录消息
func (this *UserProcess) ServerProcessloginMes(mes message.Message) (err error) {
	//核心代码
	//从mes中取出mes.Data,再将其反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2.再声明一个LoginResMes并完成赋值
	var loginResMes message.LoginResMes

	//需要到redis数据库验证用户信息是否正确
	//1.使用model.MyUserDao到redis进行验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	fmt.Println(user)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500 //状态码为500定义为表示用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 300 //状态码为300表示密码错误
			loginResMes.Error = err.Error()
		}
		//loginResMes.Code = 500 //状态码为500定义为表示登陆失败
		//loginResMes.Error = "您输入的账号或者密码不正确，请重新输入"
	} else {
		loginResMes.Code = 200 //状态码为200定义为表示登陆成功
		//这里因为用户登陆成功，要把对应的用户加入到UserMgr中
		//将登陆成功的userId赋值给this
		this.UserId = loginMes.UserId
		this.NotifyOthersOnlineUser(loginMes.UserId)
		userMgr.AddOnlineUser(this)
		//将当前在线用户的ID放入到LoginResMes.UserIds中
		//先遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, "登陆成功")
	}

	//先假定用户ID为111，密码为abc,进行判断
	/*if loginmes.UserId == 111 && loginmes.UserPwd == "abc" {
		loginResMes.Code = 200 //状态码为200定义为表示登陆成功
	} else {
		loginResMes.Code = 500 //状态码为500定义为表示登陆失败
		loginResMes.Error = "您输入的账号或者密码不正确，请重新输入"
	}*/

	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.将resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//6.发送消息,将代码封装到writePkg函数中
	tf := &utils.Tsansfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

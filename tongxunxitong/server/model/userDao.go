package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 服务器启动后，初始化一个userDao实例
// 定义一个全局变量，需要链接redis是直接使用即可
var (
	MyUserDao *UserDao
)

// 定义一个userDao结构体
type UserDao struct {
	pool redis.Pool
}

// 使用工厂模式，创建一个userDao实例
func NewUserDao(pool redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示在users哈希中没有此id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{} //将res反序列化成user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 创建一个Login方法验证登录信息 1.如果用户账号密码都正确，则返回一个user实例 2.如果发生错误，则返回错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从连接池取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *User) (err error) {
	//先从连接池取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXITS
		return
	}

	//现在就能完成注册
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	_, err = conn.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存用户注册信息失败")
		return
	}
	return
}

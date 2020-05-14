package model

import (
	"chatroom/server/model/github.com/gomodule/redigov1.8.1/redis"
	"encoding/json"
	"fmt"
)

var MyUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

//工厂模式
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserIdById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res), user),err=", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	user, err = this.getUserIdById(conn, userId)
	if err != nil {
		//fmt.Println("this.getUserIdById(conn, userId),err=",err)
		return
	}
	///获得到了user
	if user.UserPwd != userPwd {
		//密码不一致
		err = ERROR_USER_PWD

		return
	}
	return
}

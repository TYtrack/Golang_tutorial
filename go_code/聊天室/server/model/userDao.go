/*
 * @Author: your name
 * @Date: 2021-12-02 01:26:43
 * @LastEditTime: 2021-12-05 02:19:03
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/server/model/userDao.go
 */

package model

import (
	"encoding/json"
	"fmt"
	"go_code/聊天室/common"

	"github.com/garyburd/redigo/redis"
)

var MyUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

func InitUserDao(temp_pool *redis.Pool) *UserDao {
	MyUserDao = &UserDao{
		pool: temp_pool,
	}

	return MyUserDao
}

func (this *UserDao) GetUserByUserName(username string) (user common.User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("HGET", "users", username))

	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXISTS
		}
		return
	}
	registerMes := common.RegisterMes{}
	json.Unmarshal([]byte(res), &registerMes)

	user = registerMes.User

	if err != nil {
		return
	}
	return user, err

}

func (this *UserDao) CreateUser(info_string string) (user common.User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	var registerMes common.RegisterMes
	err = json.Unmarshal([]byte(info_string), &registerMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	user = registerMes.User

	_, err = this.GetUserByUserName(user.Username)

	if err == nil || err != ERROR_USER_NOT_EXISTS {
		err = ERROR_SERVER
		return
	}

	res, err := redis.Int(conn.Do("HSETNX", "users", user.Username, info_string))
	fmt.Println("userDao ", res, "   ", err)
	if err != nil {
		return
	}
	return
}

func (this *UserDao) Login(username string, password string) (user common.User, err error) {
	user, err = this.GetUserByUserName(username)
	if err != nil {
		return
	}
	if password != user.Password {
		err = ERROR_USER_PWD
	}
	return
}

func (this *UserDao) Register(info_string string) (user common.User, err error) {

	user, err = this.CreateUser(info_string)

	return
}

func (this *UserDao) GetAllUser() (users []string, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	users, err = redis.Strings(conn.Do("HKEYS", "users"))
	return
}

func (this *UserDao) PutOfflineMsg(to_username string, bytes string) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	res, err := redis.Int(conn.Do("RPUSH", "offline_msg_"+to_username, bytes))
	if err != nil {
		return
	}
	if res != 1 {
		err = ERROR_REDIS
	}
	return
}

func (this *UserDao) PutOfflineChat(to_username string, bytes string) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	res, err := redis.Int(conn.Do("RPUSH", "offline_chat_"+to_username, bytes))
	if err != nil {
		return
	}
	if res != 1 {
		err = ERROR_REDIS
	}
	return
}

func (this *UserDao) GetOfflineChat(to_username string) (res []string, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	res, err = redis.Strings(conn.Do("LRANGE", "offline_chat_"+to_username, 0, -1))
	return
}

func (this *UserDao) GetOfflineMsg(to_username string) (res []string, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	res, err = redis.Strings(conn.Do("LRANGE", "offline_msg_"+to_username, 0, -1))
	return
}

func (this *UserDao) RemOfflineChat(to_username string) {
	conn := this.pool.Get()
	defer conn.Close()
	redis.Int(conn.Do("DEL", "offline_chat_"+to_username))
}

func (this *UserDao) RemOfflineMsg(to_username string) {
	conn := this.pool.Get()
	defer conn.Close()
	redis.Int(conn.Do("DEL", "offline_msg_"+to_username))
}

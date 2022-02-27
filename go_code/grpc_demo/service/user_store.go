/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-14 02:49:43
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-14 12:52:05
 * @FilePath: /grpc_demo/service/user_store.go
 */

package service

import "sync"

type UserStore interface {
	Save(user *User) error
	Find(username string) (*User, error)
}

type InMemoryUserStore struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*User),
	}
}

func (userStore *InMemoryUserStore) Save(user *User) error {
	userStore.mu.Lock()
	defer userStore.mu.Unlock()

	if userStore.users[user.Username] != nil {
		return ErrAlreadyExists
	}

	userStore.users[user.Username] = user.Clone()
	return nil
}

func (userStore *InMemoryUserStore) Find(username string) (*User, error) {
	userStore.mu.RLock()
	defer userStore.mu.RUnlock()

	user := userStore.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil

}

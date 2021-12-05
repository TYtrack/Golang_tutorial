/*
 * @Author: your name
 * @Date: 2021-11-23 00:35:58
 * @LastEditTime: 2021-11-23 00:44:46
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/oop/model/account.go
 */
package model

import "fmt"

type Account struct {
	name     string
	balcance float64
	password string
}

func NewAccount(n string, p string, b float64) (a *Account) {
	return &Account{
		name:     n,
		password: p,
		balcance: b,
	}
}

func (a *Account) SetBalance(tt float64) {
	(*a).balcance = tt
}

func (a *Account) SetPassword(tt string) {
	if len(tt) != 6 {
		fmt.Println("密码设置错误")
	} else {
		(*a).password = tt
	}

}

func (a *Account) GetBalance() float64 {
	return (*a).balcance

}

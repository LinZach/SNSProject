package model

import (
	"SNSProject/Crypt"
	"SNSProject/DB"
	"SNSProject/session"
	"errors"
	"fmt"
	b "github.com/orca-zhang/borm"
	"strings"
	"time"
)

type User struct {
	Uid  int16 `borm:"uid"`
	Username string `borm:"username"`
	Account  string `borm:"account"`
	Password string `borm:"password"`
	Avatar string	`borm:"avatar"`
	Slogan string	`borm:"slogan"`
	Gender int	`borm:"gender"`
}

//插入用户
func Insert(user User) error {
	t := b.Table(DB.DB, "user").Debug()

	//密码加密存储
	enCode := Crypt.EnCrypt(user.Password)
	user.Password = enCode

	_, err := t.Insert(&user)
	if err != nil {
		if strings.Contains(err.Error(),"Error 1366") {
			return errors.New("用户已存在")
		}
		return err
	}

	return nil
}

//查询用户是否存在
func Query(user User) int {
	t := b.Table(DB.DB, "user").Debug()

	count, err := t.Select(&user, b.Where("account = ?", user.Account))

	if err != nil {
		fmt.Print(err)
	}

	return count
}

//查询用户是否存在 使用uid
func QueryWithUid(user User) int {
	t := b.Table(DB.DB, "user").Debug()

	count, err := t.Select(&user, b.Where("uid = ?", user.Uid))

	if err != nil {
		fmt.Print(err)
	}

	return count
}

//登录验证
func ValidataUser(user *User) bool {
	t := b.Table(DB.DB, "user").Debug()

	var localUser User
	_, err := t.Select(&localUser, b.Where("account = ?", user.Account))
	if err != nil {
		return false
	}

	//验证密码
	if Crypt.VerifyCrypt(user.Password, localUser.Password) {
		user.Uid = localUser.Uid
		return true
	}

	return false
}

//修改用户资料
func UpdateUserProfile(user User) error {
	table := b.Table(DB.DB, "user")
	_, err := table.Update(b.Where("username = ? & Avatar = ? & slogan = ? & gender = ?", user.Username, user.Avatar, user.Slogan, user.Gender))
	return err
}

//插入token
func SetToken(uid int16, token string) error {
	err := session.Set(string(uid), token, time.Hour * 24)
	return err
}


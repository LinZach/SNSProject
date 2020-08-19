package Crypt

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var Slat = "focus"

func SlatCode(code string) string {
	data := []byte(code + Slat)
	has := md5.Sum(data)
	md5Encode := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5Encode
}

func EnCrypt(code string) string {
	//这里假设客户端先md加密一次，得到编码后再加盐md5一次，并加密为数据库唯一密码存入
	hash, err := bcrypt.GenerateFromPassword([]byte(SlatCode(code)), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print(err.Error())
	}

	return string(hash)
}

func VerifyCrypt(code, localCode string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(localCode), []byte(SlatCode(code)))
	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	return true
}
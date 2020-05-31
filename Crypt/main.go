package Crypt

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
)

const Salt = "0123456789012345"

//aes解密
func Decrtpt(str string) (string, error) {
	decode := mdDecode(str)
	if !strings.Contains(decode, Salt) {
		return "", errors.New("密码解析失败")
	}

	//分割盐
	decode = strings.Split(decode, Salt)[0]
	//二次解密
	decode = mdDecode(str)
	////base64 解码一遍
	//baseDecode := base64.StdEncoding.EncodeToString([]byte(aesDecode))
	return decode, nil
}

//md5 对比
func MDCheck(content, encrypted string) bool {
	return strings.EqualFold(mdDecode(content), encrypted)
}

func mdDecode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
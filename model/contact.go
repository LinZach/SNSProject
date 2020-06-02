package model

import (
	"SNSProject/DB"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/goinggo/mapstructure"
	b "github.com/orca-zhang/borm"
)

type Contact struct {
	Uid  int32 `borm:"uid"`
	body string `borm:"contact"`
}

type ContactBody struct {
	Uid      int32  `json:"uid"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}

//添加到联系人
func AddContact(uid int32, body ContactBody) error {
	con := Contact{
		Uid: uid,
	}

	var count = 0
	var err error
	var localCon Contact
	localCon, count, err = queryContactMember(uid)

	if err != nil && count != 0 {
		return errors.New("查找对应目录失败: " + err.Error())
	}

	var isSwap bool
	isSwap, err = validateIsSwapFirend(uid, body.Uid)
	if isSwap {
		return errors.New("已经是好友")
	}

	var mapResult map[string]interface{}
	if count > 0 {//如果存在列表则获取列表
		con = localCon
		err = json.Unmarshal([]byte(con.body), &mapResult)
	}else {//没有则创建存入
		mapResult = make(map[string]interface{})
	}

	var conJson []byte
	conJson, err = json.Marshal(body)//获取添加联系人json
	if err != nil {
		return err
	}

	var mapBody map[string]interface{}
	err = json.Unmarshal(conJson, &mapBody)//将json映射到map

	keyID := strconv.FormatInt(int64(body.Uid), 10)
	mapResult[keyID] = mapBody//创建以uid为主键mapbody为value的集合

	var bodyStr []byte
	bodyStr, err = json.Marshal(mapResult)//将集合转换成json，存入
	con.body = string(bodyStr)

	table := b.Table(DB.DB, "user_contact")

	if count > 0 {//存在列表则更新列表
		_, err = table.Update(&con,b.Where(b.Eq("uid", uid)))
	}else {//不存在则创建列表
		_, err = table.Insert(&con)
	}
	return err
}

//获取联系人列表
func QueryContactList(uid int32) ([]ContactBody, error) {
	var contact Contact
	table := b.Table(DB.DB, "user_contact")
	_, err := table.Select(&contact, b.Where("uid = ?", uid))

	if err != nil {
		return nil, err
	}

	//取出联系人json，写入集合
	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(contact.body), &mapResult)
	if err != nil {
		return nil, err
	}

	//遍历集合 取出所有value 以数组返回联系人列表
	var formatContacts []ContactBody

	for _, v := range mapResult {
		var body ContactBody
		err = mapstructure.Decode(v, &body)
		if err == nil {
			formatContacts = append(formatContacts, body)
		}
	}

	return formatContacts, nil
}

//查询联系人
func queryContactMember(uid int32) (Contact, int, error) {
	var con Contact

	table := b.Table(DB.DB, "user_contact")
	count, err := table.Select(&con, b.Where("uid = ?", uid))

	return con, count, err
}

func validateIsSwapFirend(uid, fid int32) (bool, error) {
	var con Contact

	table := b.Table(DB.DB, "user_contact")
	count, err := table.Select(&con, b.Where("uid = ?", uid))

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, errors.New("不存在的用户")
	}
	//取出用户联系人对象，将联系人列表json写入集合
	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(con.body), &mapResult)
	if err != nil {
		return false, err
	}

	//验证是否添加了该用户
	fidStr := strconv.FormatInt(int64(fid), 10)
	if _, ok := mapResult[fidStr]; ok {
		return true, nil
	}
	return false, nil
}
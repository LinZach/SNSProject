package model

import (
	"SNSProject/DB"
	"encoding/json"
	"errors"

	"github.com/goinggo/mapstructure"
	"github.com/orca-zhang/borm"
)

type Contact struct {
	Uid  int32
	body string
}

type ContactBody struct {
	Uid      int32  `json:"uid"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}

//添加到联系人
func AddContact(uid int32, body ContactBody) error {
	con := Contact{
		Uid: body.Uid,
	}

	var count = 0
	var err error
	con, count, err = queryContactMember(uid)

	if err != nil {
		return errors.New("查找对应目录失败: " + err.Error())
	}

	var mapResult map[string]interface{}
	if count > 0 {
		err = json.Unmarshal([]byte(con.body), &mapResult)
	}

	var conJson []byte
	conJson, err = json.Marshal(body)
	if err != nil {
		return err
	}

	var mapBody map[string]interface{}
	err = json.Unmarshal(conJson, &mapBody)

	mapResult[string(body.Uid)] = mapBody

	var bodyStr []byte
	bodyStr, err = json.Marshal(mapResult)
	con.body = string(bodyStr)

	table := borm.Table(DB.DB, "user_contact")
	_, err = table.Insert(&con)
	return err
}

//获取联系人列表
func QueryContactList(uid int32) ([]ContactBody, error) {
	var contact Contact
	table := borm.Table(DB.DB, "user_contact")
	_, err := table.Select(&contact, borm.Where("uid = ?", uid))

	var formatContacts []ContactBody

	if err != nil {
		return nil, err
	}

	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(contact.body), &mapResult)
	if err != nil {
		return nil, err
	}

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

	table := borm.Table(DB.DB, "user_contact")
	count, err := table.Select(&con, borm.Where("uid = ?", uid))

	return con, count, err
}

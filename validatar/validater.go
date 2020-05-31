package validatar

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const(
	TYPE = "type"
	NOT_EMPTY = "NotEmpty"
	INT_MIN = "int-min"
	INT_MAX = "int-max"
	STR_LENTH = "str-len"
	STR_MAX_LENGTH = "str-max-len"
	STR_MIN_LENGTH = "str-min-len"
)

func StructValidate(bean interface{}) error {
	fields := reflect.ValueOf(bean).Elem()
	for i := 0; i < fields.NumField(); i++ {
		filed := fields.Type().Field(i)
		valid := filed.Tag.Get("valid")
		if valid == "" {
			continue
		}

		value := fields.FieldByName(filed.Name)
		err := fieldValidate(filed.Name, valid, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func fieldValidate(fileName, valid string, value reflect.Value) error {
	valids := strings.Split(valid, " ")
	for _, valid := range valids {

		if strings.Index(valid, TYPE) != -1 {
			v := value.Type().Name()
			split := strings.Split(valid, "=")
			t := split[1]
			if v != t {
				return errors.New(fileName + " type must is " + t)
			}
		}

		if strings.Index(valid, NOT_EMPTY) != -1 {
			str := value.String()
			if str == "" {
				return errors.New(fileName + "value not empty")
			}
		}
		
		if strings.Index(valid, INT_MIN) != -1 {
			v := value.Int()
			split := strings.Split(valid, "=")
			rule, err := strconv.Atoi(split[1])
			if err != nil {
				return errors.New(fileName + "验证规则有误")
			}

			if int(v) <= rule {
				return errors.New(fileName + "value must >=" + strconv.Itoa(rule))
			}
		}

		if strings.Index(valid, INT_MAX) != -1 {
			v := value.Int()
			split := strings.Split(valid, "=")
			rule, err := strconv.Atoi(split[1])
			if err != nil {
				return errors.New(fileName + "验证规则有误")
			}

			if int(v) >= rule {
				return errors.New(fileName + "value must <=" + strconv.Itoa(rule))
			}
		}

		if value.Type().Name() == "string" {
			if strings.Index(valid, STR_LENTH) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)

				if err != nil {
					return errors.New(fileName + " " + STR_LENTH + " rule is error" )
				}

				if len(v) != length {
					return errors.New(fileName + "str length must be " + lenStr)
				}
			}

			if strings.Index(valid, STR_MAX_LENGTH) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)

				if err != nil {
					return errors.New(fileName + " " + STR_LENTH + " rule is error" )
				}

				if len(v) > length {
					return errors.New(fileName + "str length <= " + lenStr)
				}
			}

			if strings.Index(valid, STR_MIN_LENGTH) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)

				if err != nil {
					return errors.New(fileName + " " + STR_LENTH + " rule is error" )
				}

				if len(v) < length {
					return errors.New(fileName + "str length >= " + lenStr)
				}
			}
		}
	}
	return nil
}

package validatar

import (
	"github.com/asaskevich/govalidator"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)
// 判断字符是否为数字
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// 判断字符是否为英文字符
func isAlpha(r rune) bool {

	if r >= 'A' && r <= 'Z' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func ispriv(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s{
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r == '_' {
			return false
		}
	}
	return true
}

//验证字符数字范围
func CheckIntScope(s string, min int64, max int64) bool {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}

	if val < min || val > max {
		return false
	}
	return true
}

//验证字符是否数字
func CheckStringDigit(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	return true
}

//验证字符是否数字+逗号,
func CheckStringCommonDigit(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s{
		if (r < '0' || r > '9') && r != ',' {
			return false
		}
	}
	return true
}

//验证字符长度
func CheckStringLenth(val string, min, max int) bool {
	if min == 0 || min ==0 {
		return false
	}
	count := utf8.RuneCountInString(val)
	if count < min || count > max {
		return  false
	}
	return true
}

//验证是否英文字母+,
func CheckStringAlpha(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r == ',' {
			return false
		}
	}
	return true
}

//验证是否URL
func CheckUrl(val string) bool {
	return govalidator.IsURL(val)
}

//验证是否英文字母数字组合
func CheckStringAlnum(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s{
		if !isDigit(r) && !isAlpha(r) &&
			r != ' ' && r != '-' && r != '!' && r != '_' &&
			r != '@' && r != '?' && r != '+' && r != ':' &&
			r != '.' && r != '/' && r != '(' && r != '\'' &&
			r != ')' && r != '·' && r != '&' {
			return false
		}
	}
	return true
}

//检查时间格式
func CheckTime(s string) bool {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	_, err := time.ParseInLocation("15:04:05", s, loc)

	if err != nil {
		return false
	}
	return true
}

//检查日期格式
func CheckDate(s string) bool {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	_, err := time.ParseInLocation("2006-01-02", s, loc)

	if err != nil {
		return false
	}
	return true
}

//检查日期时间格式 "YYYY-MM-DD HH:ii:ss"
func CheckDateTime(s string) bool {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	_, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc)

	if err != nil {
		return false
	}
	return true
}

//验证是否中文
func CheckStringChn(s string) bool {
	for _, r := range s{
		if !unicode.Is(unicode.Han, r) &&
			!isAlpha(r) && (r < '0' || r > '9') && r != '_' &&
			r != ' ' && r != '-' && r != '!' && r != '@' && r != ':' &&
			r != '?' && r != '+' && r != '.' && r != '/' && r != '\'' &&
			r != '(' && r != ')' && r != '·' && r != '&' {
			return false
		}
	}
	return true
}

//验证是否英文+汉字
func CheckStringNormal(s string) bool {
	if !CheckStringChn(s) && !CheckStringDigit(s) {
		return false
	}
	return true
}

//验证是否module格式
func CheckStringModule(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s{
		if !isAlpha(r) && r != '/' {
			return false
		}
	}
	return true
}
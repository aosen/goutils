/*
Author: Aosen
Date: 2016-01-20
QQ: 316052486
Desc:
一系列Go开发工具箱
*/

package goutils

import (
	"encoding/xml"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html/charset"
)

//将Jsonp转化为Json
// JsonpToJson modify jsonp string to json string
// Example: forbar({a:"1",b:2}) to {"a":"1","b":2}
func JsonpToJson(json string) string {
	start := strings.Index(json, "{")
	end := strings.LastIndex(json, "}")
	start1 := strings.Index(json, "[")
	if start1 > 0 && start > start1 {
		start = start1
		end = strings.LastIndex(json, "]")
	}
	if end > start && end != -1 && start != -1 {
		json = json[start : end+1]
	}
	json = strings.Replace(json, "\\'", "", -1)
	regDetail, _ := regexp.Compile("([^\\s\\:\\{\\,\\d\"]+|[a-z][a-z\\d]*)\\s*\\:")
	return regDetail.ReplaceAllString(json, "\"$1\":")
}

// The GetWDPath gets the work directory path.
func GetWDPath() string {
	wd := os.Getenv("GOPATH")
	if wd == "" {
		panic("GOPATH is not setted in env.")
	}
	return wd
}

//判断目录是否存在
// The IsDirExists judges path is directory or not.
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("util isDirExists not reached")
}

//判断文件是否存在
// The IsFileExists judges path is file or not.
func IsFileExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return !fi.IsDir()
	}

	panic("util isFileExists not reached")
}

//判断字符串是否为数字字符串
// The IsNum judges string is number or not.
func IsStringNum(a string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(a)
}

//将xml转化为map[string]string
// simple xml to string  support utf8
func XML2mapstr(xmldoc string) map[string]string {
	var t xml.Token
	var err error
	inputReader := strings.NewReader(xmldoc)
	decoder := xml.NewDecoder(inputReader)
	decoder.CharsetReader = func(s string, r io.Reader) (io.Reader, error) {
		return charset.NewReader(r, s)
	}
	m := make(map[string]string, 32)
	key := ""
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			key = token.Name.Local
		case xml.CharData:
			content := string([]byte(token))
			m[key] = content
		default:
			// ...
		}
	}

	return m
}

//将字符串转化成hash
//string to hash
func MakeHash(s string) string {
	const IEEE = 0xedb88320
	var IEEETable = crc32.MakeTable(IEEE)
	hash := fmt.Sprintf("%x", crc32.Checksum([]byte(s), IEEETable))
	return hash
}

//获取相对路径中的参数，返回参数字典
///Book/ShowBookList.aspx?tclassid=3&page=1
func GetKVInRelaPath(path string) (kv map[string]string) {
	//获取参数字符串
	l := strings.Split(path, "?")
	if len(l) == 2 {
		kvslist := strings.Split(l[1], "&")
		for _, el := range kvslist {
			kvs := strings.Split(el, "=")
			if len(kvs) == 2 {
				kv[kvs[0]] = kvs[1]
			}
		}
	}
	return
}

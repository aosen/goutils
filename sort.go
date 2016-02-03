/*
Author: Aosen
Date: 2016-02-02
Desc: 所有与排序相关的函数
*/

package goutils

import "sort"

//将map[string]string字典按照字母先后顺序排序:
//key + value .... key + value
//例如：将foo=1,bar=2,baz=3 排序为bar=2,baz=3,
//foo=1，参数名和参数值链接后，得到拼装字符串bar2baz3foo1
type kl []string

func (self kl) Len() int {
	return len(self)
}

func (self kl) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self kl) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func MapDictSortToStr(dict map[string]string) (ret string) {
	//以key为元素形成列表
	keylist := kl{}
	for key, _ := range dict {
		keylist = append(keylist, key)
	}
	//将keylist排序
	sort.Sort(keylist)
	length := keylist.Len()
	for i := 0; i < length; i++ {
		ret = ret + keylist[i] + dict[keylist[i]]
	}
	return
}

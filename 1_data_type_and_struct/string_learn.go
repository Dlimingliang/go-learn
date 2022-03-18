package main

import (
	"fmt"
	"strings"
)

func main() {

	var defaultS string
	fmt.Println(defaultS == "") // true

	s := "aaabbbccc"
	//判断字符串是否以**开头
	fmt.Println(strings.HasPrefix(s, "a"))
	//判断字符串是否以**结尾
	fmt.Println(strings.HasSuffix(s, "a"))
	//判断字符串是否包含子字符串
	fmt.Println(strings.Contains(s, "b"))
	//第一个出现位置，没有就-1
	fmt.Println(strings.Index(s, "a"))
	//字符串出现最后的位置
	fmt.Println(strings.LastIndex(s, "a"))
	//替换字符串 将字符串中前n个log字符串替换为新字符串 n为-1 表示替换所有
	fmt.Println(strings.Replace(s, "aaa", "ddd", 1))
	//计算字符串str在字符串s中出现的非重叠次数
	fmt.Println(strings.Count(s, "bb"))
	//重复字符串
	fmt.Println(strings.Repeat(s, 2))
	//转大小写
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	//修剪字符串
	fmt.Println(strings.Trim(s, "bbb"))
	fmt.Println(strings.Trim(s, "a"))
	//分隔字符串，拼接slice
	str2 := "hello world"
	sl := strings.Fields(str2)
	fmt.Println(strings.Join(sl, "-"))

}

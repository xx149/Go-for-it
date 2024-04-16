package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转换为整数
	intStr := "123"
	intValue, err := strconv.Atoi(intStr)
	if err != nil {
		fmt.Println("字符串转整数出错:", err)
	} else {
		fmt.Println("整数值:", intValue)
	}
	// 字符串转换为整数（指定进制）
	hexStr := "1a"
	hexValue, err := strconv.ParseInt(hexStr, 16, 0)
	if err != nil {
		fmt.Println("16进制字符串转整数出错:", err)
	} else {
		fmt.Println("16进制整数值:", hexValue)
	}
}

/*
 * @Author: your name
 * @Date: 2022-04-11 21:33:22
 * @LastEditTime: 2022-04-11 21:39:32
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /struct_json/Users/bytedance/go/src/github.com/singgel/golang-base/strconv_Parse/main.go
 */
package main

import (
	"fmt"
	"strconv"
)

// Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
func main() {
	b, _ := strconv.ParseBool("true")
	fmt.Printf("boolean: %v\n", b)
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("float: %v\n", f)
	i, _ := strconv.ParseInt("-2", 10, 64)
	fmt.Printf("int: %v\n", i)
	u, _ := strconv.ParseUint("2", 10, 64)
	fmt.Printf("uint: %v\n", u)
}

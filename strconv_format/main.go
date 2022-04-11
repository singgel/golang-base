/*
 * @Author: your name
 * @Date: 2022-04-11 21:39:52
 * @LastEditTime: 2022-04-11 21:42:51
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /struct_json/Users/bytedance/go/src/github.com/singgel/golang-base/strconv_format/main.go
 */
package main

import (
	"fmt"
	"strconv"
)

// Format系列函数实现了将给定类型数据格式化为string类型数据的功能。
// isPrint() 返回一个字符是否是可打印的
// CanBackquote() 返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。
func main() {
	s1 := strconv.FormatBool(true)
	fmt.Println(s1)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s2)
	s3 := strconv.FormatInt(-2, 16)
	fmt.Println(s3)
	s4 := strconv.FormatUint(2, 16)
	fmt.Println(s4)
}

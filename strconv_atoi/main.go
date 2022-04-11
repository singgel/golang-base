/*
 * @Author: your name
 * @Date: 2022-04-11 21:31:08
 * @LastEditTime: 2022-04-11 21:33:16
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /struct_json/Users/bytedance/go/src/github.com/singgel/golang-base/strconv_atoi/main.go
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
}

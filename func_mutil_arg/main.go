/*
 * @Author: your name
 * @Date: 2022-04-11 15:16:29
 * @LastEditTime: 2022-04-11 15:20:30
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /struct_json/Users/bytedance/go/src/github.com/singgel/golang-base/func_mutil_arg/main.go
 */
package main

import "fmt"

func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	sum := intSum3(1, 2, 3)
	fmt.Println(sum)
	sum1, sum2 := calc(1, 2)
	fmt.Printf("sum1: %v, sum2: %v\n", sum1, sum2)
}

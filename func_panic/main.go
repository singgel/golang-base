/*
 * @Author: your name
 * @Date: 2022-04-28 14:24:11
 * @LastEditTime: 2022-04-28 14:45:58
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/func_panic/main.go
 */
package main

import "fmt"

func printA() {
	fmt.Println("A")
}

func printB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("B recover from: %v\n", err)
		}
	}()
	panic("not implemented")
}

func printC() {
	fmt.Println("C")
}

func main() {
	printA()
	printB()
	printC()
}

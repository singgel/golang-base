/*
 * @Author: your name
 * @Date: 2022-04-11 21:48:01
 * @LastEditTime: 2022-04-11 21:55:18
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /struct_json/Users/bytedance/go/src/github.com/singgel/golang-base/interface_define/main.go
 */
package main

import "fmt"

type Sayer interface {
	Say()
}

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}

func main() {
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()
}

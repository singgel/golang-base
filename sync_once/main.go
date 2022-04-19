/*
 * @Author: your name
 * @Date: 2022-04-19 18:28:07
 * @LastEditTime: 2022-04-19 18:31:49
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_once/main.go
 */
package main

import (
	"fmt"
	"sync"
)

func main() {

	doOnce()

}

func doOnce() {

	var once sync.Once

	onceBody := func() {

		fmt.Println("Only once")

	}

	//用于等待协程执行完毕

	done := make(chan bool)

	//启动10个协程执行once.Do(onceBody)

	for i := 0; i < 10; i++ {

		go func() {

			//把要执行的函数(方法)作为参数传给once.Do方法即可

			once.Do(onceBody)

			done <- true

		}()

	}

	for i := 0; i < 10; i++ {

		<-done

	}

}

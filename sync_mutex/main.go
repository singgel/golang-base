/*
 * @Author: your name
 * @Date: 2022-04-19 18:28:21
 * @LastEditTime: 2022-04-19 18:30:00
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_mutex/main.go
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum int

	mutex sync.Mutex
)

func main() {

	for i := 0; i < 100; i++ {

		go add(10)

	}

	for i := 0; i < 10; i++ {

		go fmt.Println("和为:", readSum())

	}

	time.Sleep(2 * time.Second)

}

//增加了一个读取sum的函数，便于演示并发

func readSum() int {

	mutex.Lock()
 
	defer mutex.Unlock()
 
	b:=sum
 
	return b
 
 }

func add(i int) {

	mutex.Lock()

	defer mutex.Unlock()

	sum += i

}

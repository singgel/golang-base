/*
 * @Author: your name
 * @Date: 2022-04-19 16:27:52
 * @LastEditTime: 2022-04-19 16:48:23
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_wg/main.go
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	stopCh := make(chan bool)
	wg.Add(1)

	go func() {

		defer wg.Done()

		watchDog("【监控狗1】", stopCh)

	}()

	time.Sleep(5 * time.Second)

	stopCh <- true //发停止指令

	wg.Wait()

}

func watchDog(name string, stopCh chan bool) {

	//开启for select循环，一直后台监控

	for {

		select {
		case <-stopCh:

			fmt.Println(name, "停止指令已收到，马上停止")

			return

		default:

			fmt.Println(name, "正在监控……")

		}

		time.Sleep(1 * time.Second)

	}

}

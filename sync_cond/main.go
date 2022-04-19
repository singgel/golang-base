/*
 * @Author: your name
 * @Date: 2022-04-19 17:02:36
 * @LastEditTime: 2022-04-19 17:03:13
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_cond/main.go
 */
//10个人赛跑，1个裁判发号施令
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	race()
}

func race() {

	cond := sync.NewCond(&sync.Mutex{})

	var wg sync.WaitGroup

	wg.Add(11)

	for i := 0; i < 10; i++ {

		go func(num int) {

			defer wg.Done()

			fmt.Println(num, "号已经就位")

			cond.L.Lock()

			cond.Wait() //等待发令枪响

			fmt.Println(num, "号开始跑……")

			cond.L.Unlock()

		}(i)

	}

	//等待所有goroutine都进入wait状态

	time.Sleep(2 * time.Second)

	go func() {

		defer wg.Done()

		fmt.Println("裁判已经就位，准备发令枪")

		fmt.Println("比赛开始，大家准备跑")

		cond.Broadcast() //发令枪响

	}()

	//防止函数提前返回退出

	wg.Wait()

}

/*
 * @Author: your name
 * @Date: 2022-04-20 20:01:47
 * @LastEditTime: 2022-04-20 20:15:18
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_race/main.go
https://juejin.cn/post/6844903918233714695

go run -race sync_race/main.go
 首先，通过 time.AfterFunc 创建 timer，定时的间隔从 randomDuration 函数获得，定时函数打印消息，然后通过 timer 的 Reset 方法重置定时器，重复利用。
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 程序中存在 2 个 goroutine 非同步读写变量 t。
// 如果初始定时时间非常短，就可能出现在主函数还未对 t 赋值，定时函数已经执行，而此时 t 仍然是 nil，无法调用 Reset 方法。
// func main() {
// 	start := time.Now()
// 	var t *time.Timer
// 	t = time.AfterFunc(randomDuration(), func() {
// 		fmt.Println(time.Now().Sub(start))
// 		t.Reset(randomDuration())
// 	})

// 	time.Sleep(5 * time.Second)
// }

func main() {
	start := time.Now()
	reset := make(chan bool)
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		reset <- true
	})
	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

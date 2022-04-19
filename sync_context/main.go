/*
 * @Author: your name
 * @Date: 2022-04-19 16:50:22
 * @LastEditTime: 2022-04-19 16:56:34
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_context/main.go
 */
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, stop := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		watchDog(ctx, "watch dog")
	}()
	time.Sleep(time.Second * 5)
	stop()

	wg.Wait()

}

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "copy that, stop now")
			return
		default:
			fmt.Println(name, "watching")
		}
		time.Sleep(1 * time.Second)
	}
}

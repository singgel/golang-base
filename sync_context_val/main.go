/*
 * @Author: your name
 * @Date: 2022-04-19 17:19:09
 * @LastEditTime: 2022-04-19 17:35:08
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_context_val/main.go
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
	wg.Add(4) //记得这里要改为4，原来是3，因为要多启动一个协程

	//省略其他无关代码
	ctx := context.Background()
	valCtx := context.WithValue(ctx, "userId", 2)
	cancelCtx, cancel := context.WithCancel(valCtx)

	go func() {

		defer wg.Done()

		getUser(cancelCtx)

	}()

	//省略其他无关代码
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}

func getUser(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():

			fmt.Println("【获取用户】", "协程退出")

			return

		default:

			userId := ctx.Value("userId")

			fmt.Println("【获取用户】", "用户ID为：", userId)

			time.Sleep(1 * time.Second)

		}

	}

}

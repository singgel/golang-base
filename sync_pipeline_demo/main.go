/*
 * @Author: your name
 * @Date: 2022-04-19 17:48:25
 * @LastEditTime: 2022-04-19 18:00:50
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_pipeline_demo/main.go
 */
package main

import "fmt"

func main() {
	buys := buy(10)
	builds := build(buys)
	packs := pack(builds)

	for p := range packs {
		fmt.Println(p)
	}
}

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		for i := 0; i < n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range in {
			out <- "组装" + v
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range in {
			out <- "打包" + v
		}
	}()
	return out
}

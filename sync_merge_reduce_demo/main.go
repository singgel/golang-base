/*
 * @Author: your name
 * @Date: 2022-04-19 18:10:23
 * @LastEditTime: 2022-04-19 18:14:50
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_merge_reduce_demo/main.go
 */
package main

import (
	"fmt"
	"sync"
)

func main() {
	coms := buy(100)    //采购100套配件

	//三班人同时组装100部手机
 
	phones1 := build(coms)
 
	phones2 := build(coms)
 
	phones3 := build(coms)
 
	//汇聚三个channel成一个
 
	phones := merge(phones1,phones2,phones3)
 
	packs := pack(phones) //打包它们以便售卖
 
	//输出测试，看看效果
 
	for p := range packs {
 
	   fmt.Println(p)
 
	}
}

//扇入函数（组件），把多个chanel中的数据发送到一个channel中

func merge(ins ...<-chan string) <-chan string {

	var wg sync.WaitGroup

	out := make(chan string)

	//把一个channel中的数据发送到out中

	p := func(in <-chan string) {

		defer wg.Done()

		for c := range in {

			out <- c

		}

	}

	wg.Add(len(ins))

	//扇入，需要启动多个goroutine用于处于多个channel中的数据

	for _, cs := range ins {

		go p(cs)

	}

	//等待所有输入的数据ins处理完，再关闭输出out

	go func() {

		wg.Wait()

		close(out)

	}()

	return out

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

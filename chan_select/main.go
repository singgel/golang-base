/*
 * @Author: your name
 * @Date: 2022-04-19 16:31:03
 * @LastEditTime: 2022-04-19 16:33:47
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/chan_select/main.go
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	//声明三个存放结果的channel

	firstCh := make(chan string)

	secondCh := make(chan string)

	threeCh := make(chan string)

	//同时开启3个goroutine下载

	go func() {

		firstCh <- downloadFile("firstCh")

	}()

	go func() {

		secondCh <- downloadFile("secondCh")

	}()

	go func() {

		threeCh <- downloadFile("threeCh")

	}()

	//开始select多路复用，哪个channel能获取到值，

	//就说明哪个最先下载好，就用哪个。

	select {

	case filePath := <-firstCh:

		fmt.Println(filePath)

	case filePath := <-secondCh:

		fmt.Println(filePath)

	case filePath := <-threeCh:

		fmt.Println(filePath)

		// default:
		// 	fmt.Println("玩完了!")

	}
	fmt.Println("玩完了!")

}

func downloadFile(chanName string) string {

	//模拟下载文件,可以自己随机time.Sleep点时间试试

	time.Sleep(time.Second)

	return chanName + ":filePath"

}

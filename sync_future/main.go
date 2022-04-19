/*
 * @Author: your name
 * @Date: 2022-04-19 18:19:36
 * @LastEditTime: 2022-04-19 18:22:47
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /golang-base/sync_future/main.go
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	vagetableCh := washVegetables()
	waterCh := boilWater()

	fmt.Println("烧水洗菜，稍等会儿")

	time.Sleep(5 * time.Second)
	fmt.Println("要做火锅了，看看菜水好了吗")

	vegetables := <-vagetableCh
	water := <-waterCh

	fmt.Println("准备好了", vegetables, water)
}

//洗菜

func washVegetables() <-chan string {

	vegetables := make(chan string)

	go func() {

		time.Sleep(5 * time.Second)

		vegetables <- "洗好的菜"

	}()

	return vegetables

}

//烧水

func boilWater() <-chan string {

	water := make(chan string)

	go func() {

		time.Sleep(5 * time.Second)

		water <- "烧开的水"

	}()

	return water

}

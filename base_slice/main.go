/*
 * @Author: hekuangsheng hekuangsheng@bytedance.com
 * @Date: 2022-05-07 17:09:12
 * @LastEditors: hekuangsheng hekuangsheng@bytedance.com
 * @LastEditTime: 2022-05-07 17:09:18
 * @FilePath: /golang-base/base_slice/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	for idx, val := range slice {
		fmt.Printf("数组具体元素的地址 %p\n", &slice[idx])
		fmt.Printf("val的地址 %p\n", &val)
	}
}

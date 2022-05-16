/*
 * @Author: hekuangsheng hekuangsheng@bytedance.com
 * @Date: 2022-05-07 15:15:51
 * @LastEditors: hekuangsheng hekuangsheng@bytedance.com
 * @LastEditTime: 2022-05-07 16:29:20
 * @FilePath: /golang-base/file_write/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("hello 小王子") //直接写入字符串数据

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件

	error := ioutil.WriteFile("./xx.txt", []byte("str hello沙河"), 0666)
	if error != nil {
		fmt.Println("write file failed, error:", error)
		return
	}
}

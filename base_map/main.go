/*
 * @Author: your name
 * @Date: 2022-04-14 10:57:17
 * @LastEditTime: 2022-04-15 12:09:24
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /struct_json/Users/bytedance/go/src/github.com/singgel/golang-base/base_map/main.go
 */
package main

import "fmt"

type EmployeeIdInfo struct {
	employeeId  int32  `thrift:"employeeId,1,required" json:"employeeId"`
	emailPrefix string `thrift:"emailPrefix,2,required" json:"emailPrefix"`
}

func main() {
	m := make(map[*EmployeeIdInfo]string, 3)
	empList := []EmployeeIdInfo{
		{emailPrefix: "小王子", employeeId: 18},
		{emailPrefix: "娜扎", employeeId: 23},
		{emailPrefix: "大王八", employeeId: 9000},
	}

	for _, emp := range empList {
		m[&emp] = emp.emailPrefix
	}
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}

/*
 * @Author: your name
 * @Date: 2022-04-14 10:57:17
 * @LastEditTime: 2022-04-19 16:40:08
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /struct_json/Users/bytedance/go/src/github.com/singgel/golang-base/base_obj_map/main.go
 */
package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() { // Issue 46399
	var m sync.Map
	for i, v := range [3]string{"hello", "world", "Go"} {
		m.Store(i, v)
	}
	m.Range(func(key, value any) bool {
		m.Range(func(key, value any) bool {
			// We should be able to load the key offered in the Range callback,
			// because there are no concurrent Delete involved in this tested map.
			if v, ok := m.Load(key); !ok || !reflect.DeepEqual(v, value) {
				fmt.Printf("Nested Range loads unexpected value, got %+v want %+v", v, value)
			}

			// We didn't keep 42 and a value into the map before, if somehow we loaded
			// a value from such a key, meaning there must be an internal bug regarding
			// nested range in the Map.
			if _, loaded := m.LoadOrStore(42, "dummy"); loaded {
				fmt.Println("Nested Range loads unexpected value, want store a new value")
			}

			// Try to Store then LoadAndDelete the corresponding value with the key
			// 42 to the Map. In this case, the key 42 and associated value should be
			// removed from the Map. Therefore any future range won't observe key 42
			// as we checked in above.
			val := "sync.Map"
			m.Store(42, val)
			if v, loaded := m.LoadAndDelete(42); !loaded || !reflect.DeepEqual(v, val) {
				fmt.Printf("Nested Range loads unexpected value, got %v, want %v", v, val)
			}
			return true
		})

		// Remove key from Map on-the-fly.
		m.Delete(key)
		return true
	})

	// After a Range of Delete, all keys should be removed and any
	// further Range won't invoke the callback. Hence length remains 0.
	length := 0
	m.Range(func(key, value any) bool {
		length++
		return true
	})

	if length != 0 {
		fmt.Printf("Unexpected sync.Map size, got %v want %v", length, 0)
	}
}

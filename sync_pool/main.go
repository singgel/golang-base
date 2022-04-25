package main

import (
 "fmt"
 "sync"
 "sync/atomic"
)

// 用来统计实例真正创建的次数
var numCalcsCreated int32

// 创建实例的函数
func createBuffer() interface{} {
 // 这里要注意下，非常重要的一点。这里必须使用原子加，不然有并发问题；
 atomic.AddInt32(&numCalcsCreated, 1)
 buffer := make([]byte, 1024)
 return &buffer
}

func main() {
 // 创建实例
 bufferPool := &sync.Pool{
  New: createBuffer,
 }

 // 多 goroutine 并发测试
 numWorkers := 1024 * 1024
 var wg sync.WaitGroup
 wg.Add(numWorkers)

 for i := 0; i < numWorkers; i++ {
  go func() {
   defer wg.Done()
   // 申请一个 buffer 实例
   buffer := bufferPool.Get()
   _ = buffer.(*[]byte)
   // 释放一个 buffer 实例
   defer bufferPool.Put(buffer)
  }()
 }
 wg.Wait()
 fmt.Printf("%d buffer objects were created.\n", numCalcsCreated)
}
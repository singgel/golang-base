<!--
 * @Author: your name
 * @Date: 2022-04-19 10:27:41
 * @LastEditTime: 2022-04-19 20:58:29
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%A
 * @FilePath: /golang-base/golang_CSP.md
-->
## CSP form wiki
[Communicating sequential processes](https://en.wikipedia.org/wiki/Communicating_sequential_processes)  
翻译：在计算机科学中，交谈循序程序（英语：Communicating sequential processes，缩写为CSP），又译为通信顺序进程、交换消息的循序程序，是一种形式语言，用来描述并发性系统间进行交互的模式。它是叫做进程代数或进程演算的关于并发的数学理论家族的一员，基于了通过通道的消息传递。CSP高度影响了Occam的设计，也影响了编程语言如Limbo、RaftLib、Go、 Crystal和Clojure的core.async等。  

[about CSP history](https://clojure.org/news/2013/06/28/clojure-clore-async-channels#_history)  
翻译：这种风格的根源至少可以追溯到Hoare的通信顺序过程（CSP），其次是在例如ocam、JavaCSP和Go编程语言中的实现和扩展。  
在现代化身中，通道的概念成为一流的，这样做为我们提供了我们所寻求的间接性和独立性。  
通道的一个关键特征是阻塞。在最原始的形式中，无缓冲通道充当会合点，任何读取器都将等待写入器，反之亦然。可以引入缓冲，但不鼓励无界缓冲，因为带阻塞的有界缓冲可能是协调起搏和背压的重要工具，确保系统不会承担超出其能力的工作。

## Actor CSP
Actor模型非常适用于多个组件独立工作，相互之间仅仅依靠消息传递的情况。如果想在多个组件之间维持一致的状态  
1.线程池方案  
Java1.5后，Doug Lea的Executor系列被包含在默认的JDK内，是典型的线程池方案。  

2.异步回调方案、GreenThread/Coroutine/Fiber方案(也就是大家常说的协程)  
为了解决回调方法带来的难题，这种方案的思路是写代码的时候还是按顺序写，但遇到IO等阻塞调用时，将当前的代码片段暂停，保存上下文，让出当前线程。等IO事件回来，然后再找个线程让当前代码片段恢复上下文继续执行，写代码的时候感觉好像是同步的，仿佛在同一个线程完成的，但实际上系统可能切换了线程，但对程序无感。（全都在用户态）  

3.Goroutine  
内置了一个调度器，实现了Coroutine的多线程并行调度，同时通过对网络等库的封装，对用户屏蔽了调度细节。提供了Channel机制，用于Goroutine之间通信，实现CSP并发模型（Communicating Sequential Processes）。因为Go的Channel是通过语法关键词提供的，对用户屏蔽了许多细节。其实Go的Channel和Java中的SynchronousQueue是一样的机制，如果有buffer其实就是ArrayBlockQueue。  

## Java Akka 
* Akka（Scala,Java）基于线程和异步回调模式实现  
* 要么是Akka提供的异步框架，要么通过Future-callback机制，转换成回调模式  
* Quasar (Java) 为了解决Akka的阻塞回调问题，Quasar通过字节码增强的方式，在Java中实现了Coroutine/Fiber  

## Golang Goroutine
* 内置了Coroutine机制。因为要用户态的调度，必须有可以让代码片段可以暂停/继续的机制。内置了一个调度器，实现了Coroutine的多线程并行调度，同时通过对网络等库的封装，对用户屏蔽了调度细节  
* Go实现了M:N的调度，Goroutine任务执行，相当于一种rebalance机制  
* 系统启动时，会启动一个独立的后台线程空闲挂起[runtime.gopark]，忙碌调出[将event中的pollDesc取出来，找到关联的阻塞Goroutine]  

1.Go通过Goroutine的调度解决了CPU利用率的问题，goroutine 被 Go runtime 所调度，这一点和线程不一样。也就是说，Go 语言的并发是由 Go 自己所调度的，自己决定同时执行多少个 goroutine，什么时候执行哪几个。这些对于我们开发者来说完全透明，只需要在编码的时候告诉 Go 语言要启动几个 goroutine，至于如何调度执行，我们不用关心  

2.操作系统的线程一般都有固定的栈内存（通常为2MB）,而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB），所以在 Go 语言中一次创建数万个 goroutine 也是可能的。并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小， Go 的 runtime 会自动为 goroutine 分配合适的栈空间  

3.互联网在线应用场景下，如果每个请求都扔到一个Goroutine里，当资源出现瓶颈的时候，会导致大量的Goroutine阻塞，最后用户请求超时。(比如带锁的共享资源，比如数据库连接等。这时候就需要用Goroutine池来进行控流)  

## Golang CSP VS Actor
* CSP模型里消息和Channel是主体，处理器是匿名的（channel与数据类型绑定）  

类比Java中的Future机制：  
Java的Future机制是异步通信的机制，在主线程中可以启动一个FutureTask，启动之后要是不想阻塞的等待FutureTask的执行结果，可以在FutureTask执行的同时，非阻塞的干其他的事情，当我们要获得FutureTask结果的时候，调用Task的get方法获取结果，在get的时候，要是FutureTask已经执行完毕，就可以立即拿到结果，但要是FutureTask尚未执行完毕，就会阻塞的等待，直到FutureTask执行完毕，才能够继续执行下面的代码 

go里面有阻塞式和非阻塞式两种： 
没有声明容量的为阻塞式： retCh := make(chan string)  
声明了容量的的Channel为非阻塞式： retCh := make(chan string, 1)  

* Actor模型里Actor是主体，Mailbox（类似于CSP的Channel）是透明的（队列与类型不强相关）  

---
# 补充
## Golang sync
* [sync.Cond：Wait、Signal、Broadcast](https://github.com/singgel/golang-base/blob/main/sync_cond/main.go)  
注意：在调用 Signal 或者 Broadcast 之前，要确保目标协程处于 Wait 阻塞状态，不然会出现死锁问题。  
* [sync.WaitGroup](https://github.com/singgel/golang-base/blob/main/sync_wg/main.go)  
用来控制等待所有协程结束  


## Context
* [Context](https://github.com/singgel/golang-base/blob/main/sync_context/main.go)  
一个任务会有很多个协程协作完成，一次 HTTP 请求也会触发很多个协程的启动，而这些协程有可能会启动更多的子协程，并且无法预知有多少层协程、每一层有多少个协程。  
Context 就是用来简化解决这些问题的，并且是并发安全的。Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作。  
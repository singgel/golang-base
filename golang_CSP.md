<!--
 * @Author: your name
 * @Date: 2022-04-19 10:27:41
 * @LastEditTime: 2022-04-20 20:43:21
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%A
 * @FilePath: /golang-base/golang_CSP.md
-->
[toc]

# CSP简述
## CSP form wiki
[Communicating sequential processes](https://en.wikipedia.org/wiki/Communicating_sequential_processes)  
翻译：在计算机科学中，交谈循序程序（英语：Communicating sequential processes，缩写为CSP），又译为通信顺序进程、交换消息的循序程序，是一种形式语言，用来描述并发性系统间进行交互的模式。它是叫做进程代数或进程演算的关于并发的数学理论家族的一员，基于了通过通道的消息传递。CSP高度影响了Occam的设计，也影响了编程语言如Limbo、RaftLib、Go、 Crystal和Clojure的core.async等  

[about CSP history](https://clojure.org/news/2013/06/28/clojure-clore-async-channels#_history)  
翻译：这种风格的根源至少可以追溯到Hoare的通信顺序过程（CSP），其次是在例如ocam、JavaCSP和Go编程语言中的实现和扩展  
在现代化身中，通道的概念成为一流的，这样做为我们提供了我们所寻求的间接性和独立性  
通道的一个关键特征是阻塞。在最原始的形式中，无缓冲通道充当会合点，任何读取器都将等待写入器，反之亦然。可以引入缓冲，但不鼓励无界缓冲，因为带阻塞的有界缓冲可能是协调起搏和背压的重要工具，确保系统不会承担超出其能力的工作。

# 并发演进
[并发之痛 Thread，Goroutine，Actor](http://jolestar.com/parallel-programming-model-thread-goroutine-actor/)  
1. 竞态条件（race conditions） 如果每个任务都是独立的，不需要共享任何资源，那线程也就非常简单。但世界往往是复杂的  
2. 依赖关系以及执行顺序 如果线程之间的任务有依赖关系，需要等待以及通知机制来进行协调  

基于Java扩展：
> 内存（线程的栈空间）:每个线程都需要一个栈（Stack）空间来保存挂起（suspending）时的状态。Java的栈空间（64位VM）默认是1024k，不算别的内存，只是栈空间，启动1024个线程就要1G内存。
> 调度成本（context-switch）:国外一篇论文专门分析线程切换的成本，基本上得出的结论是切换成本和栈空间使用大小直接相关。
> CPU使用率:想提高CPU利用率，最大限度的压榨硬件资源，从这个角度考虑，我们应该用多少线程呢？没有固定答案，因为网络的时间不是固定的，另外比如锁，比如数据库连接池，就会更复杂。
因此我们从以上的讨论可以得出一个结论：  
1. 线程的成本较高（内存，调度）不可能大规模创建  
2. 应该由语言或者框架动态解决这个问题  

## Actor CSP
Actor模型非常适用于多个组件独立工作，相互之间仅仅依靠消息传递的情况。如果想在多个组件之间维持一致的状态  
1. 线程池方案  
Java1.5后，Doug Lea的Executor系列被包含在默认的JDK内，是典型的线程池方案  
* 缺点： *
线程池一定程度上控制了线程的数量，实现了线程复用，降低了线程的使用成本。但还是没有解决数量的问题  

2. 异步回调方案、GreenThread/Coroutine/Fiber方案(也就是大家常说的协程)  
为了解决回调方法带来的难题，这种方案的思路是写代码的时候还是按顺序写，但遇到IO等阻塞调用时，将当前的代码片段暂停，保存上下文，让出当前线程。等IO事件回来，然后再找个线程让当前代码片段恢复上下文继续执行，写代码的时候感觉好像是同步的，仿佛在同一个线程完成的，但实际上系统可能切换了线程，但对程序无感。（全都在用户态）  
* 缺点： *
从一个线程切换到另一个线程需要完整的上下文切换。因为可能需要多次内存访问，索引这个切换上下文的操作开销较大，会增加运行的cpu周期。

3. Goroutine  
内置了一个调度器，实现了Coroutine的多线程并行调度，同时通过对网络等库的封装，对用户屏蔽了调度细节。提供了Channel机制，用于Goroutine之间通信，实现CSP并发模型（Communicating Sequential Processes）。因为Go的Channel是通过语法关键词提供的，对用户屏蔽了许多细节。其实Go的Channel和Java中的SynchronousQueue是一样的机制，如果有buffer其实就是ArrayBlockQueue  
* 好处： *
区别于操作系统内核调度操作系统线程，goroutine 的调度是Go语言运行时（runtime）层面的实现，是完全由 Go 语言本身实现的一套调度系统——go scheduler。它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行。下面[Golang Goroutine](#golang-goroutine)会有介绍  
单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的， goroutine 则是由Go运行时（runtime）自己的调度器调度的，完全是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多  

## Java Akka 
* Akka（Scala,Java）基于线程和异步回调模式实现  
* 要么是Akka提供的异步框架，要么通过Future-callback机制，转换成回调模式  
* Quasar (Java) 为了解决Akka的阻塞回调问题，Quasar通过字节码增强的方式，在Java中实现了Coroutine/Fiber  

## Golang Goroutine
* 内置了Coroutine机制。因为要用户态的调度，必须有可以让代码片段可以暂停/继续的机制。内置了一个调度器，实现了Coroutine的多线程并行调度，同时通过对网络等库的封装，对用户屏蔽了调度细节  
* Go实现了M:N的调度，Goroutine任务执行，相当于一种rebalance机制  
* 系统启动时，会启动一个独立的后台线程空闲挂起[runtime.gopark]，忙碌调出[将event中的pollDesc取出来，找到关联的阻塞Goroutine]  

### Goroutine优缺点
1. Go通过Goroutine的调度解决了CPU利用率的问题，goroutine 被 Go runtime 所调度，这一点和线程不一样。也就是说，Go 语言的并发是由 Go 自己所调度的，自己决定同时执行多少个 goroutine，什么时候执行哪几个。这些对于我们开发者来说完全透明，只需要在编码的时候告诉 Go 语言要启动几个 goroutine，至于如何调度执行，我们不用关心  

2. 操作系统的线程一般都有固定的栈内存（通常为2MB）,而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB），所以在 Go 语言中一次创建数万个 goroutine 也是可能的。并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小， Go 的 runtime 会自动为 goroutine 分配合适的栈空间  

3. 互联网在线应用场景下，如果每个请求都扔到一个Goroutine里，当资源出现瓶颈的时候，会导致大量的Goroutine阻塞，最后用户请求超时。(比如带锁的共享资源，比如数据库连接等。这时候就需要用Goroutine池来进行控流)  

### Goroutine调度机制
![gpm](pic/gpm.png)
- G：表示 goroutine，每执行一次go f()就创建一个 G，包含要执行的函数和上下文信息  
- 全局队列（Global Queue）：存放等待运行的 G  
- P：表示 goroutine 执行所需的资源，最多有 GOMAXPROCS 个，GOMAXPROCS 默认值是机器上的 CPU 核心数  
- P 的本地队列：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建 G 时，G 优先加入到 P 的本地队列，如果本地队列满了会批量移动部分 G 到全局队列  
- M：线程想运行任务就得获取 P，从 P 的本地队列获取 G，当 P 的本地队列为空时，M 也会尝试从全局队列或其他 P 的本地队列获取 G。M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去  
- Goroutine 调度器和操作系统调度器是通过 M 结合起来的，每个 M 都代表了1个内核线程，操作系统调度器负责把内核线程分配到 CPU 的核上执行  

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
## Channel
* [channel](https://github.com/singgel/golang-base/blob/main/chan_select/main.go)  
：一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的  
使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。  
Go语言中提供了单向通道来处理这种需要限制通道只能进行某种操作的情况。
``` go
<- chan int // 只接收通道，只能接收不能发送
chan <- int // 只发送通道，只能发送不能接收
```

## Golang sync
* [sync.Mutex、sync.RWMutex](https://github.com/singgel/golang-base/blob/main/sync_mutex/main.go)  
注意：使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的  
* [sync.WaitGroup](https://github.com/singgel/golang-base/blob/main/sync_wg/main.go)  
用来控制等待所有协程结束  
* [sync.Cond：Wait、Signal、Broadcast](https://github.com/singgel/golang-base/blob/main/sync_cond/main.go)  
注意：在调用 Signal 或者 Broadcast 之前，要确保目标协程处于 Wait 阻塞状态，不然会出现死锁问题  

* [sync/atomic](https://github.com/singgel/golang-base/blob/main/sync_atomic/main.go)  
atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者 sync 包的函数/类型实现同步更好  

## Context
* [Context](https://github.com/singgel/golang-base/blob/main/sync_context/main.go)  
一个任务会有很多个协程协作完成，一次 HTTP 请求也会触发很多个协程的启动，而这些协程有可能会启动更多的子协程，并且无法预知有多少层协程、每一层有多少个协程  
Context 就是用来简化解决这些问题的，并且是并发安全的。Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作  

## Race
* [race](https://github.com/singgel/golang-base/blob/main/sync_race/main.go)  
Go语言中单元测试的时候加上-race参数，可以实现并发测试，但 race-enabled 程序耗费的 CPU 和内存通常是正常程序的十倍，在真实环境下一直启用竞态检测是非常不切合实际的  
[issues/3970](https://github.com/golang/go/issues/3970)这个官方源码案例告诉我们：它不会发出假的提示，认真严肃地对待它的每条警示非常必要。但它并非万能，还是需要以你对并发特性的正确理解为前提，才能真正地发挥出它的价值  
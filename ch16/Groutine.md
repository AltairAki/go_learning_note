# Go 协程机制
Thread vs. Groutine

1. 创建时默认的 stack 的大小
 - JDK5 以后 Java Thread stack 默认为1M
 - Groutine 的 Stack 初始化大小为2K （创建起来更快）
2. 和 KSE (Kernel Space Entity) （直接由 CPU 调度）
 - Java Thread 是 1:1  
 - Groutine 是 M:N

 ![image](http://altairaki.cn/wp-content/uploads/2020/08/%E5%8D%8F%E7%A8%8B%E6%9C%BA%E5%88%B6.jpg)

M - System Thread M代表内核级线程，一个M就是一个线程，goroutine就是跑在M之上的；M是一个很大的结构，里面维护小对象内存cache（mcache）、当前执行的goroutine、随机数发生器等等非常多的信息
P - Processor Golang实现的协程处理器，主要用途就是用来执行goroutine的，所以它也维护了一个goroutine队列，里面存储了所有需要它来执行的goroutine
G - Groutine 代表一个协程，它有自己的栈，instruction pointer和其他信息（正在等待的channel等等），用于调度。

如果一个协程运行的时间特别长，把整个协程队列都阻塞吗？
 
Go 起来的时候，会有一个守护线程计数，记每个Processor完成协程的数量，当一段时间发现某个Processor完成的协程数量，没有变化。就会往这个协程任务栈内置一个标记。 当这个协程运行到非内联函数的时候就会读到这个标记，就会把自己中断下来，插到等候协程队列的队尾。Processor切换成别的协程继续运行。


当某一个协程被系统中断了，如：IO 需要等待的时候，为了提高整体的并发。Processor 会把自己移到另一个可使用的系统线程当中，继续执行这个队列里挂的其他协程。当中断的协程被唤醒，完成IO之后。会把自己加入到其中某一个 processor 的协程等待队列里，或者是全局等待队列里。 
 
当一个协程被中断时，它在寄存器里的运行状态，也会保存在这个协程对象里。当协程获得再一次运行的机会时，运行状态又会重新写入寄存器，继续运行。


参考：[golang核武器goroutine调度原理、channel详解](https://www.cnblogs.com/wdliu/p/9272220.html)
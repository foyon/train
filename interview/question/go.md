### php数组为什么是有序的

​	php数组底层是hashtable，每个元素都有下一层的指针，内存是连续的，所以foreach比for、while效率高



yaf框架底层原理

#php垃圾回收机制

#go垃圾回收机制，结构体包含什么

#go slice与map区别与底层实现

#go内存及 堆 栈

#go性能优化思路

#go tool -S trace.out 

#go pprof 

#attach 火焰图等

#go channel 关闭等问题，channel 底层实现

#go interface底层实现

#go 类型断言

#golang context包ctx

#golang实现线程池

### golang 主goroutine 广播模式

利用close chan特性， _, ok := range chanT ; chan close情况会直接返回false
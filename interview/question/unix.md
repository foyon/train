系统相关面试题

### 解决hash冲突方法

* 链地址法：hash(k)=i,hashmap，相同的i，hash单链追加

* 再哈希法：这种方法是同时构造多个不同的哈希函数：

  ​                   Hi=RH1（key） i=1，2，…，k 

​                          当哈希地址Hi=RH1（key）发生冲突时，再计算Hi=RH2（key）……，直到冲突不再产生。这种方法不易产生聚集，但增加了计算时间。

* 开放定址法
* 建立公共溢出区： 分为基本表和溢出表，冲突的一律放入溢出表。

Kill 与kill -s 9 区别

select与epoll区别，epoll结构体包含什么

死锁

进程与线程，通信机制

Tcp、ip协议相关（20B～） udp（8B）

客户度和服务端各自生产自己的序列号 +1 syn 确认机制，

滑动窗口控制流量，最终校验机制





linux性能优化思路，火焰图

Byte、bit

consul

ngx_lua有哪些阶段

分布式锁

LNMP fast-cgi源码相关



#### GPU 与CPU区别

cpu 很强的通用性来处理不同业务场景，GPU处理单一高吞吐，性能超高。

Gpu thread > cpu

ALU GPU> cpu, cache Gpu << cpu,  ALU计算单元GPU超多，每个cache小，CPU反之。

就算密集型，易于并行的适用于GPU，成百上千个核



###  SSH相关 客户端、服务端证书生成



敏感词匹配算法？



安全的登陆状况？



### int查询快还是字符串查询快？

int转二进制

char 通过ascii 进行逆转，多了一次逆解操作，不过现在很多在编译阶段已经处理，效率差不多。



进程、线程和协程











### DDD & OOP 

####  DDD

​		 领域驱动设计模式，面向对象分析，对事物抽象。首先考虑的是业务，提供某一类model，不依赖后端架构实现。约定的是交互标准。Clean架构、六边形架构、CQRS、Event Source等架构。

#### OOP

​		面向project，面向接口的封装开发，把组件的实现和接口分开。



protobuf & json



k8s包含什么？各负责什么

api

ectd

continer

布隆过滤？bitmap原理？

raft选举算法原理？

熔断限流？
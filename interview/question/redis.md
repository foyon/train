#### redis面试题系列

### redis 实例扩容机制？

1、按业务分实例。主从模式

2、垂直类扩容，直接改参数。

3、集群类：

​			新建一个目标实例（集群），新老之间同步数据redisport、redishake，然后程序该地址切换。有损，有覆盖丢失风险，时序问题。 

​			或者就是改lvs代理地址  NAT、DR、TRUN，vip直接面向用户请求，作为用户请求的目的目标地址。NAT改变用户header头里真实CIP地址。

### mysql 与redis 如何保证数据一致性？

锁，try catch

mysql事物

先delete redis，再delete mysql， 再delete redis

利用mysql binlog，记录offset 导入redis（阿里提供开源软件）



### redis分布式集群

https://www.cnblogs.com/kaleidoscope/p/9630316.html

### 基于代理的分片

- 客户端发送请求到一个代理，代理解析客户端的数据，将请求转发至正确的节点，然后将结果回复给客户端。

- 开源方案

  - Twemproxy
  - codis
  - ![img](https://upload-images.jianshu.io/upload_images/1521743-5b8e35bb04543259.png)

  #### Redis Cluster

  - Redis官网推出，可线性扩展到1000个节点

  - 无中心架构

  - 一致性哈希思想

  - 客户端直连redis服务，免去了proxy代理的损耗

  - #### Redis-cluster原理

    - Hash slot。集群内的每个redis实例监听两个tcp端口，6379（默认）用于服务客户端查询，16379（默认服务端口 + 10000）用于集群内部通信。
      - key的空间被分到16384个hash slot里；
      - 计算key属于哪个slot，CRC16(key) & 16384。
    - 节点间状态同步：gossip协议，最终一致性。节点间通信使用轻量的二进制协议，减少带宽占用。

  

  ![img](https://upload-images.jianshu.io/upload_images/1521743-158efb87b969a0a8.png)

  

  - 

### redis分布式锁

set key val expire NX （仅当key不存在的时候才能进行操作，redis对于command处理是单进程）

### redis rdb和aof持久化

都是fork子进程进行处理

rdb： bgsave 异步

​          save 同步，周期从内存 copy 写入磁盘

aof： 每次写命令计入日志，通过fork进行写入磁盘






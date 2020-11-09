/**
* @File:          distribute_lock.go
* @Aim:           分布式锁实现
* @Author:        foyon
* @Created Time:  2020-11-09
 */

/*
1、基于go 内存加锁
2、基于go-channel特性
3、redis 原子性 setnx
	业务发展到一定量级的话，就需要从多方面来考虑了。
	首先是你的锁是否在任何恶劣的条件下都不允许数据丢失，如果不允许，
	那么就不要使用Redis的setnx的简单锁
    设置过期时间因为网络问题会存在数据不一致性问题
4、通过mysql自增primary key
5、通过zookeeper
	基于ZooKeeper的锁与基于Redis的锁的不同之处在于Lock成功之前会一直阻塞
	需要考虑吞吐量
6、通过etcd raft一致性算法

=========
不同的应用场景取锁
按业务进行分片设置规则
*/

package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var counter int
var lock sync.Mutex
var wg sync.WaitGroup

func main() {
	//getLockMem()
	//getLockChannel()
	getRedisLock()

}

// 通过内存临界加锁
func getLockMem() { // {{{
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//临界加锁操作
			lock.Lock()
			counter++
			//lock.Lock()
			lock.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
} // }}}

//通过channel 阻塞特性
func getLockChannel() { // {{{

	chl := newChannelLock()
	for i := 0; i < 1000; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			//临界加锁操作
			if !chl.getChannelLock() {
				fmt.Printf("can not get lock [%v]\n", i)
				return
			}

			counter++
			chl.UnLockChannel()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
} // }}}

type ChLock struct {
	C chan struct{}
}

func newChannelLock() ChLock { // {{{
	var chl ChLock
	//如果无缓冲，会造成dead lock
	chl.C = make(chan struct{}, 1)
	chl.C <- struct{}{}
	return chl
} // }}}

func (chl ChLock) getChannelLock() bool { // {{{
	getLock := false

	select {
	case <-chl.C:
		getLock = true
	default:
	}

	return getLock
} // }}}

func (chl ChLock) UnLockChannel() { // {{{
	chl.C <- struct{}{}
} // }}}

//通过redis setNx实现分布式锁
func getRedisLock() { // {{{
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			redisLock()
		}()
	}
	wg.Wait()
} // }}}

func redisLock() { // {{{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var lockKey = "counter_lock"
	var counterKey = "counter"

	//redis cmd: set foyon "test" EX 5 NX
	//setnx仅当redis不存在的时候操作key
	//set 命令可以支持，需要设置过期时间防止死锁
	//[NX|XX] NX仅当不存在的时候可以取锁
	//获取到锁后，执行一致性步骤完毕 进行Del 即可
	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
} // }}}

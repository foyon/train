package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
1.一批任务(合并包)
设定超时
各job执行状态
报错信息捕捉
主动关闭各job

*/

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(val int, wg *sync.WaitGroup) {
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			fmt.Println("finished job id:", val)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}

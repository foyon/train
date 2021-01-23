package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//以合并包任务为例,当然不一定适用这种场景哈
func main() {
	outChan := make(chan int)
	errChan := make(chan error)
	finishChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(outChan chan<- int, errChan chan<- error, val int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			fmt.Println("finished job id:", val)
			outChan <- val
			if val == 60 {
				errChan <- errors.New("error in 60")
			}

		}(outChan, errChan, i, &wg)
	}

	go func() {
		wg.Wait()
		fmt.Println("finish all job")
		close(finishChan)
	}()

Loop:
	for {
		//阻塞
		select {
		case val := <-outChan:
			fmt.Println("finished:", val)
		case err := <-errChan:
			fmt.Println("error:", err)
			break Loop
		case <-finishChan:
			break Loop
		//超时限制
		case <-time.After(3 * time.Millisecond):
			fmt.Println("echo timeout")
			//break Loop
			//return
			break
		}
	}
}

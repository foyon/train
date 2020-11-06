/**
* @File:          goroutine_pool.go
* @Aim:           go程池实现,比较复杂 不推荐
* @Author:        foyon
* @Created Time:  2020-11-06
 */
package main

import (
	"fmt"
	"time"
)

type Gopool struct {
	//job传递
	Queue chan func() error
	//开启go程数
	Worker int
	//任务总数
	Total int

	//执行结果
	Result chan error

	FinishCallBack func()
}

func (gp *Gopool) Init(worker int, total int) { // {{{
	gp.Queue = make(chan func() error, total)
	gp.Worker = worker
	gp.Total = total
	gp.Result = make(chan error, total)
} // }}}

func (gp *Gopool) Start() { // {{{
	for i := 0; i < gp.Worker; i++ {
		go func() {
			for {
				task, ok := <-gp.Queue
				if !ok {
					break
				}

				err := task()
				gp.Result <- err
			}

		}()
	}

	for j := 0; j < gp.Total; j++ {
		if res, ok := <-gp.Result; !ok {
			fmt.Printf(" error: %v", res)
		}
	}

	if gp.FinishCallBack != nil {
		gp.FinishCallBack()
	}
} // }}}

func (gp *Gopool) Stop() { // {{{
	close(gp.Queue)
	close(gp.Result)
} // }}}

func (gp *Gopool) AddTask(task func() error) { // {{{
	gp.Queue <- task
} // }}}

func (gp *Gopool) SetFinshCallBack(callback func()) { // {{{
	gp.FinishCallBack = callback
} // }}}

/*##############################################################*/

func main() {
	testPool()
}

func testPool() { // {{{

	jobs := []string{
		"a",
		"b",
		"c",
		"d",
	}

	pool := new(Gopool)
	pool.Init(3, len(jobs))

	for _, v := range jobs {
		//拷贝
		v := v
		pool.AddTask(func() error {
			return Echo(v)
		})
	}

	isFinish := false
	pool.SetFinshCallBack(func() {
		func(isFinish *bool) {
			*isFinish = true
		}(&isFinish)
	})

	pool.Start()

	for !isFinish {
		time.Sleep(time.Millisecond * 100)
	}

	pool.Stop()
	fmt.Println("exec over")
} // }}}

func Echo(v string) error { // {{{
	fmt.Printf("this goroutine echo %v \n", v)
	return nil
} // }}}

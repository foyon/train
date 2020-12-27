package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	smap := sync.Map{}
	wg := sync.WaitGroup{}

	smap.Store("age", 10)

	wg.Add(5)

	go func() {
		defer wg.Done()
		i := 0
		for i < 1000 {
			//smap.Store("one", 10)
			i++
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		for i < 1000 {
			//smap.Store("one", 11)
			//tf, _ := smap.Load("one")
			//fmt.Printf("%s", tf)
			i++
		}
	}()

	go func() {
		defer wg.Done()
		//tick := time.NewTicker(2 * time.Second)
		for i := 0; ; i++ {
			set(smap, i)
			time.Sleep(5 * time.Second)
		}

	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		for true {
			get(smap)
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		time.Sleep(time.Second * 2)
		//fmt.Println(smap.Load("one"))
		fmt.Printf("hhh")
		wg.Done()

	}()
	wg.Wait()
}

func set(s sync.Map, i int) {
	s.Store("one", i)
}

func get(s sync.Map) {
	v, _ := s.Load("one")
	fmt.Printf("%d", v)
}

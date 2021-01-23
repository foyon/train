package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func main3() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for b := 0; b < 10; b++ {
		go func() {
			fmt.Println("go-1i: ", b)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("go-2i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	var m []map[string]interface{}
	t := map[string]interface{}{
		"test": 1,
	}
	m = append(m, t)
	m = append(m, t)
	fmt.Printf("%+v", count(m))

	return
	mark := 151333
	ms := strconv.Itoa(mark)
	mt := []rune(ms)
	fmt.Printf("%s", string(mt[0]))

	mark2 := string(9743)
	fmt.Printf("%s", string(mark2[0]))
}
func main4() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	//for {
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
	//}
}

func main2() {
	strTest := "+"
	fmt.Println(interface{}(strTest).(string))

	var s2 = [...]map[string]interface{}{{`你好`: "foyon"}, {`1`: "1"}}

	for k, v := range s2 {
		fmt.Printf("%v => %s\n", k, v)
	}

	//sliceTest := []map[string]interface{}{{},{}}

	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

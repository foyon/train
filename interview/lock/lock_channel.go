package main

import (
	"net/http"
	"sync"
)

var mutex = sync.Mutex{}
var ch = make(chan bool, 1)

func UseMutex() {
	mutex.Lock()

	_, _ = http.Get("https://cn.bing.com/search?q=s&qs=n&form=QBLH&sp=-1&pq=s&sc=9-1&sk=&cvid=1E9A04C55A8F4807B6C96432425F5F86")
	mutex.Unlock()
}

func UseChan() {
	ch <- true

	_, _ = http.Get("https://cn.bing.com/search?q=s&qs=n&form=QBLH&sp=-1&pq=s&sc=9-1&sk=&cvid=1E9A04C55A8F4807B6C96432425F5F86")
	<-ch
}

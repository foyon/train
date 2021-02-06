package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)
type WriteFile struct {
	Fd *bufio.Writer
	FileObj *os.File
}

//不要并发写入，open同一个fd，会导致切片出错
//同时打开fd，不同线程写是一个串行阻塞写入情况

func main(){
	//Case02()
	 Case01()
}
/*
1. 单线程 同一个文件，open fd 可以被重复打开，在未close情况两个fd可写，是串行的，buffer不够的情况未知
2.多go func 并发执行会导致数据不准确，对一个被open的再次open 进行写入 可以写入
结论：在往文件进行写入尽量避免



*/

func Case01() {
	fmt.Println("vim-go")


	go func(){
		wf2 := NewWriteFile()
		defer wf2.Fd.Flush()
		for i:=0;i<100;i++{
			go wf2.testWrite1(fmt.Sprintf("wf02-1 step:%d", i*2))
			go wf2.testWrite2(fmt.Sprintf("wf02-1 step:%d", i*20))
		}
	}()

	go func(){
		wf := NewWriteFile()
		defer wf.Fd.Flush()

		for i:=0;i<100;i++{
			if i == 5 {
				time.Sleep(5 * time.Second)
			}
			go wf.testWrite1(fmt.Sprintf("wf01-1 step:%d", i))
			go wf.testWrite2(fmt.Sprintf("wf01-2 step:%d", i*10))
		}
	}()
  time.Sleep(15 * time.Second)

	return
}

//同一个Fd 不同go func 写入

func Case02() {
	wf := NewWriteFile()
	defer wf.Fd.Flush()
	go func(){
		for i:=1;i<10000;i++{
			wf.testWrite1(fmt.Sprintf("wf01-1 step:%d", i))
		}
	}()

	go func(){
		for i:=0;i<10000;i++{
			//go wf.testWrite2(fmt.Sprintf("wf01-2 step:%d", i*10))
		}
	}()

	time.Sleep(10 * time.Second)
	return
}

func NewWriteFile()  *WriteFile{
	return &WriteFile{
		Fd: func() *bufio.Writer{
			fileObj, err := os.OpenFile("./test_file.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
				//defer fileObj.Close()
				if err != nil{
					panic(err)
				}
				writeObj := bufio.NewWriter(fileObj)
				return writeObj
		}(),
	}
}

func (w *WriteFile)testWrite1(i string){
	//fmt.Printf("test write 1 begin")

	rs, err := w.Fd.WriteString(fmt.Sprintf("%s \n", i))
	if err != nil {
		fmt.Printf("rs: %d err:%s \n", rs, err)
	}

}
func (w *WriteFile)testWrite2(i string){
	_, _ = w.Fd.WriteString(fmt.Sprintf("%s \n", i))

}
/**
* @File:          dxme.go
* @Aim:
* @Author:        foyon
* @Created Time:  2020-10-19
 */
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	stu := student{
		Name: "fan",
	}
	fmt.Println(stu)
	//f(&stu)
	jsonTostruct()

}

type student struct {
	Name string
	Cal  func()
}

func f(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		//msg.Name
		//type interface {} is interface with no methods)
		//msg.Cal
		fmt.Printf("%+v", msg)
		fmt.Printf("exec")
	}
}

type People struct {
	//name string 无法渲染,巨坑
	Name string `json:"name"`
	//id =0 无法渲染
	Id int
}

func jsonTostruct() {
	js := `{
        "name":"s11",
		"id":2
    }`
	var p People
	//p := new(People)
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("%+v", p)
}

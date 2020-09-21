package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/cleanimg", cleanimg)
	e := http.ListenAndServe(":9001", nil)

	if nil != e {
		fmt.Println(e.Error())
	}
}

/*
*通过http请求清除指定图片
*key 清楚图片的名字
 */
func cleanimg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//key := valuesGetDefault(r.Form, "key", false)
	//fmt.Fprintf(w, key)
	key := r.Form.Get("key")
	imgPath := "/data/wwwroot/user_feedback/rdb/img/" + key
	os.Remove(imgPath)
	fmt.Fprintf(w, "ok")
}

package main

import (
	"fmt"
	"net/http"
	"os"
	//"os/signal"
	"externalorder"
	"time"
)

func main() {
	go func() {
		http.HandleFunc("/cleanimg", cleanimg)
		e := http.ListenAndServe(":9001", nil)

		if nil != e {
			fmt.Println(e.Error())
		}
	}()

	/**
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutdown Server ...")
	*/

	mainTicker := time.NewTicker(time.Second * 5)
	for k := range mainTicker.C {
		// 定时对各种服务进行状态检查
		//player.CheckIdle()
		//tkproxy.RunATick()
		//cagproxy.RunATick()
		//caggr.RunATick()
		//eca.RunATick()
		//sms.RunATick()

		// 外部指令检查
		externalorder.CheckOrder()
		if externalorder.IsNeedExit() {
			break
		}

		/**
		fmt.Println(k)
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		fmt.Println("Shutdown Server ...")
		*/

	}

	/**
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutdown Server ...")
	*/

	mainTicker.Stop()

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

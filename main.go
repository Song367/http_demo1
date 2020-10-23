package main

import (
	"fmt"
	httppoint "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"golang.org/x/time/rate"
	"http_demo1/Servers"
	"http_demo1/uti"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)


func main3(){
	user:=Servers.UserServer{}

	// 限流
	limit:=rate.NewLimiter(1,5)
	// 使用rate.allow的限流函数
	endpoint:=Servers.Ratelimit(limit)(Servers.GetUserInfo(user))
	server:=httppoint.NewServer(endpoint,Servers.DecodeRequest,Servers.EncodeResponse)

	route:=mux.NewRouter()
	{
		route.Methods("GET","DELETE").Path("/user/{uid:\\d+}").Handler(server)
		route.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type","application/json")
			writer.Write([]byte(`{"status":"200"}`))
		})
	}

	/* 获取本地ip */
	//conn, err := net.Dial("udp", "www.google.com.hk:80")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//defer conn.Close()
	//fmt.Println(conn.LocalAddr().String())
	//fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
	errChan:=make(chan error)
	go (func(){
		uti.ConsulReg()

		err:=http.ListenAndServe(":"+strconv.Itoa(uti.ServerPort),route)
		//err:=http.ListenAndServe(":8080",route)
		if err!=nil{
			log.Error(err)
			errChan<-err
		}
	})()
	go (func(){
		sig_c:=make(chan os.Signal)
		signal.Notify(sig_c,syscall.SIGINT,syscall.SIGTERM)      // 监听信号
		errChan<-fmt.Errorf("%s",<-sig_c)
	})()
	getChan:=<-errChan
	uti.UnRegister()
	fmt.Println("getchan",getChan)
}



package sse

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func Test_cheng(t *testing.T) {
	srv := New()

	// Create a new Mux and set the handler
	mux := http.NewServeMux()
	mux.HandleFunc("/events", srv.ServeHTTP)
	srv.CreateStream("test")

	go func() {
		for a := 0; a < 1000; a++ {
			srv.Publish("test", &Event{Data: []byte("haha\n")})
			time.Sleep(time.Second * 2)
		}
	}()

	// 示例端点，用于发布事件
	/*mux.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		// 发布一个事件
		srv.Publish("messages", &Event{
			Data: []byte("你好，SSE！"),
		})
		fmt.Fprintln(w, "事件已发布")
	})*/

	go func() {
		c := NewClient("http://127.0.0.1:8080/events?stream=test")
		c.Subscribe("test", "test", func(msg *Event) {
			if msg.Data != nil {
				fmt.Println("--->", string(msg.Data))
			}
		})
	}()
	http.ListenAndServe(":8080", mux)
}

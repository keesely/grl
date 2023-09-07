package curl

import (
	"fmt"
	"math/rand"
	"time"
)

// type CURL_OPTIONS struct {
// 	Method string
// 	Url    string
// 	Header map[string]string
// 	Body   string
// }

type Response struct {
	Body string
}

func request(url string) string {
	// mock request time delay rand 1-3s
	rand := time.Duration(1 + rand.Intn(3))
	time.Sleep(time.Second * rand)
	// TODO
	return "response by " + url
}

func nowForamt() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

func Go(url string, parallel int) {
	// 实现简单的并发请求 返回结果输出到通道
	// 非阻塞
	channels := make(chan Response, parallel)
	for i := 0; i < parallel; i++ {
		nurl := fmt.Sprintf("%s/%d", url, i+1)
		go func() {
			res := request(nurl)
			fmt.Println("request", nurl, "response", res, " at ", nowForamt())
			channels <- Response{Body: res}
		}()
	}
	// 接收通道数据并输出
	for i := 0; i < parallel; i++ {
		res := <-channels
		fmt.Println(res.Body, " at ", nowForamt())
	}
	close(channels)
}

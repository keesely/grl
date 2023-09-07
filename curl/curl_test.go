package curl_test

// 测试Go()并发 url: https://pkg.go.dev/net/http#hdr-Clients_and_Transports

import (
	"fmt"
	"grl/curl"
	"testing"
)

func TestGo(t *testing.T) {
	url := "https://pkg.go.dev/net/http"
	parallel := 10
	curl.Go(url, parallel)
	fmt.Println("done")
}

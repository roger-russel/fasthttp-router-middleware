package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func errorCode(expected int, received int, route string) string {
	return fmt.Sprintf("Expecting Http Status Code: %v, but received: %v at %s", expected, received, route)
}

func TestMain(t *testing.T) {

	var url = "http://127.0.0.1:8101"

	doneChan := make(chan struct{})
	go func() {
		main()
	}()

	time.Sleep(1 * time.Second) // Wait for fasthttp start

	go func() {

		var resp []byte

		code, _, _ := fasthttp.Get(resp, url)
		if code != 200 {
			t.Errorf(errorCode(200, code, url))
		}

		code, _, _ = fasthttp.Get(resp, url+"/unauthorized")
		if code != 401 {
			t.Errorf(errorCode(401, code, url+"/unauthorized"))
		}

		code, _, _ = fasthttp.Get(resp, url+"/forbidden")
		if code != 403 {
			t.Errorf(errorCode(403, code, url+"/forbidden"))
		}

		code, _, _ = fasthttp.Get(resp, url+"/authorized")

		if code != 200 {
			t.Errorf(errorCode(200, code, url+"/authorized"))
		}

		doneChan <- struct{}{}
	}()

	<-doneChan
}

func ExampleMain() {

	time.Sleep(1 * time.Second) // Wait for fasthttp start

	main()
	// Output:
	// Error in ListenAndServe:  listen tcp4 :8101: bind: address already in use
}

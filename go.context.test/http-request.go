// http-request.go
//

package main

import (
	"fmt"
	"context"
	"time"
)

// 创建一个int的context
type reqCtx int

func (*reqCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*reqCtx) Done() <-chan struct{} {
	return nil
}

func (*reqCtx) Err() error {
	return nil
}

func (*reqCtx) Value(key interface{}) interface{} {
	return nil
}

type goReq struct {
	ctx *reqCtx
	url string
	Timeout int
}

func (gr goReq) Get() string {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	gr.ctx = ctx
	return ""
}

func main() {
	var ctx reqCtx = 22
	var gr goReq = goReq{&ctx, "https://baidu.com", 10}
	gr.Get()

	fmt.Println("use fmt....")
}


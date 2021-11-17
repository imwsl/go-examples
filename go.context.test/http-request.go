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
	url
}


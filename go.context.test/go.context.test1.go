//
// context.Context 
// 
// go.context.test.go
//
//
//

/*
	context.Context
	
	type Context interface {
		Deadline() (deadline time.Time, ok bool)
		Done() <-chan struct {}
		Err() err
		Value(key interface{}) interface{}
	}
	
	1、Deadline: 返回context.Context被取消的时间，也就是完成工作的截止日期。
	2、Done: 返回一个Channel，这个Channel会在当前工作完成或上下文被取消后关闭，多次调用Done方法返回同一个Channel。
	3、Err: 返回context结束的原因:
		a、context.Context被取消，返回Canceled错误。
		b、context.Context超时，返回DeadlineExceeded错误。
		
	context.Context在不同的goroutine之间同步特定数据，取消信号及处理请求的截止日期。
*/

package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	
	go handle(ctx, 500 * time.Millisecond)
	select {
		case <-ctx.Done():
			fmt.Println("Done() main ", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
		case <-ctx.Done():
			fmt.Println("Done() handle ", ctx.Err())
		case <-time.After(duration):
			fmt.Println("process request with ", duration)
	}
}

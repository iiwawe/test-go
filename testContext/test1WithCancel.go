package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger


//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

/**
 * 获取context
 * context.WithCancel
 * 执行cancel()函数发送取消信号
 */
func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()

}

/**
 * 获取context
 * context.WithCancel
 * 执行cancel()函数发送取消信号
 */
func timeoutHandler() {
	 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// go doTimeOutStuff(ctx)
	go doStuff(ctx)
	time.Sleep(10 * time.Second)

	cancel()

}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	//someHandler()
	timeoutHandler()
	logg.Printf("down all")
}

package main

import (
	"time"
	"context"
	"log"
)

func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok { //设置了deadl
			log.Printf("deadline set")
			if time.Now().After(deadline) {
				log.Printf(ctx.Err().Error())
				return
			}

		}

		select {
		case <-ctx.Done():
			log.Printf("done")
			return
		default:
			log.Printf("work")
		}
	}
}

func deadlineHandler() {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// go doTimeOutStuff(ctx)
	go doTimeOutStuff(ctx)
	time.Sleep(10 * time.Second)

	cancel()

}

func main() {

	//someHandler()
	//timeoutHandler()
	deadlineHandler()
	log.Printf("down all")
}
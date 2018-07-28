package main

import (
	"os"
	"context"
	"time"
	"log"
	"testing"
	"net/http"
)

var logg *log.Logger



func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok { //设置了deadl
			logg.Printf("deadline set")
			if time.Now().After(deadline) {
				logg.Printf(ctx.Err().Error())
				return
			}

		}
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func TestContext1(t *testing.T) {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}

/*
 *
 */
func timeoutHandler() {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	 go doTimeOutStuff(ctx)
	//go doStuff(ctx)

	time.Sleep(10 * time.Second)

	cancel()

}

func TestContext2(t *testing.T) {
	logg = log.New(os.Stdout, "", log.Ltime)
	timeoutHandler()
	logg.Printf("end")
}

/*
 *
 */
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
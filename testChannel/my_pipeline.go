package testChannel

import "sync"

func gen(num ...int) <- chan int {
	out := make(chan int)
	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 为每一个输入channel cs 创建一个 goroutine output
	// output 将数据从 c 拷贝到 out，直到 c 关闭，然后 调用 wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 启动一个 goroutine，用于所有 output goroutine结束时，关闭 out
	// 该goroutine 必须在 wg.Add 之后启动
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

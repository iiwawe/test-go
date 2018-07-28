package testChannel

import (
	"fmt"
	"testing"
	"time"
)

func TestMyChannel1(t *testing.T) {
	go fmt.Println("1")
	fmt.Println("2")
}

func TestMyChannel2(t *testing.T) {
	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	fmt.Println("2")
}

//无缓存的Channel放进去的值必须取出来
//否则会造成阻塞
func TestMyChannel3(t *testing.T) {
	ch := make(chan int)

	ch <- 1    //线程阻塞，之后代码无法执行

	go func() {
		i := <-ch
		fmt.Println(i)
	}()
	fmt.Println("2")
}

func TestMyChannel33333(t *testing.T) {
	ch := make(chan int)

	go func() {                //先开启以恶搞线程读取chan, 不要使主线程堵塞
		i := <-ch
		fmt.Println(i)
	}()

	ch <- 1    //线程阻塞

	fmt.Println("2")
}

/*
 *
 * 生产和消费的例子，使用不带缓存的channel
 *
 */
func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send: ", i)
	}
}

func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive: ", v)
	}
}

func TestMyChannel4(t *testing.T) {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}

/*
 *
 * 生产和消费的例子，使用带缓存的channel
 *
 */
func produce2(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer2(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func TestMyChannel5(t *testing.T) {
	ch := make(chan int, 10)
	go produce2(ch)
	go consumer2(ch)
	time.Sleep(1 * time.Second)
}

/*
 *
 * 测试非缓存chan的阻塞情况
 * ？？？？存在疑问？？？？
 *
 */
func TestMyChannel6(t *testing.T) {
	nums := make([]int, 5, 8)
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
			fmt.Println(len(out))
		}
		close(out)
	}()
	fmt.Println("---end---")
}

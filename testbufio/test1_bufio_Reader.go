package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//test1()
	//test2()
	test3()
}

// 示例：Peek、Read、Discard、Buffered
func test1() {

	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	buf := bufio.NewReaderSize(sr, 0)
	b := make([]byte, 10)

	fmt.Println(buf.Buffered()) // 0
	s, _ := buf.Peek(5)
	s[0], s[1], s[2] = 'a', 'b', 'c'
	fmt.Printf("%d   %q\n", buf.Buffered(), s) // 16   "abcDE"

	buf.Discard(1)

	for n, err := 0, error(nil); err == nil; {
		n, err = buf.Read(b)
		fmt.Printf("%d   %q   %v\n", buf.Buffered(), b[:n], err)
	}
	// 5   "bcDEFGHIJK"   <nil>
	// 0   "LMNOP"   <nil>
	// 6   "QRSTUVWXYZ"   <nil>
	// 0   "123456"   <nil>
	// 0   "7890"   <nil>
	// 0   ""   EOF
}

//示例：ReadLine
func test2() {
	sr, _ := os.Open("F:\\gopath\\src\\bonc.com\\testGo\\testbufio\\file")
	//sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n1234567890")
	buf := bufio.NewReaderSize(sr, 1024)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFGHIJKLMNOP"   true   <nil>
	// "QRSTUVWXYZ"   false   <nil>
	// "1234567890"   false   <nil>
	// ""   false   EOF

	fmt.Println("----------")

	// 尾部有一个换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG\n"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFG"   false   <nil>
	// ""   false   EOF

	fmt.Println("----------")

	// 尾部没有换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFG"   false   <nil>
	// ""   false   EOF
}

//示例 test readString
func test3() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n1234567890")
	buf := bufio.NewReaderSize(sr, 1024)
	delim := byte('Y')
	line, err := buf.ReadString(delim)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(line)
}

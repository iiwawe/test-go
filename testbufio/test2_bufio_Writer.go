package main

import (
	"bufio"
	"fmt"
	"os"
)

// 示例：Available、Buffered、WriteString、Flush
func main() {
	name := "F:\\gopath\\src\\bonc.com\\testGo\\testbufio\\flll"
	//file, err := os.Create(name)
	//if err != nil {
	//	fmt.Errorf("create error")
	//}

	//注意：添加权限
	//O_RDONLY：只读模式(read-only)
	//O_WRONLY：只写模式(write-only)
	//O_RDWR：读写模式(read-write)
	//O_APPEND：追加模式(append)
	//O_CREATE：文件不存在就创建(create a new file if none exists.)
	//O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
	//O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
	//O_TRUNC：打开并清空文件
	file, err := os.OpenFile(name, os.O_RDWR, 0)
	defer file.Close()
	if err != nil {
		fmt.Errorf("read error")
	}

	buf := bufio.NewWriterSize(file, 0)
	fmt.Println(buf.Available(), buf.Buffered()) // 4096 0

	buf.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	fmt.Println(buf.Available(), buf.Buffered()) // 4070 26

	// 缓存后统一输出，避免终端频繁刷新，影响速度
	buf.Flush()  //ABCDEFGHIJKLMNOPQRSTUVWXYZ

	//ioutil.WriteFile(name, []byte("www"), 0655)
}
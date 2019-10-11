package main

import (
	"fmt"
	"github.com/CodisLabs/codis/pkg/models/zk"
	"time"
)

func main() {
	//s := float32(5.4)
	//ss := 300 * s
	//d := float32(160.00)
	//dd := d / 1240
	//fmt.Println(dd)
	testZk()
}

func testZk() {
	client, err := zkclient.New("172.16.11.133", "", time.Second)
	if err != nil {
		fmt.Print("err")
	}
	data, err := client.Read("/redis/raas/cluster/redis:test-w456", true)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(string(data))
}

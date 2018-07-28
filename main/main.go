package main

import (

	"github.com/CodisLabs/codis/pkg/models/zk"
	"time"
	"fmt"
)

func main() {
	client, err := zkclient.New("172.16.11.133", "", time.Second)
	if err != nil{
             fmt.Print("err")
	}
	data, err := client.Read("/redis/raas/cluster/redis:test-w456", true)
	if err != nil{
		fmt.Print(err.Error())
	}
	fmt.Println(string(data))
}

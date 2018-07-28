package testChannel

import (
	"testing"
	"fmt"
)

/*
 *
 * 流水线入门：求平方数
 *
 */
func TestMyPipline(t *testing.T) {
	//设置流水线
	c := gen(2, 3, 8)
	out := sq(c)
	//消费输出结果
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

/*
 *
 * 流水线进阶：扇入和扇出
 *
 */
func TestMerge(t *testing.T)  {
	in := gen(2, 3)

	// 启动两个 sq 实例，即两个goroutines处理 channel "in" 的数据
	c1 := sq(in)
	c2 := sq(in)

	// merge 函数将 channel c1 和 c2 合并到一起，这段代码会消费 merge 的结果
	for n := range merge(c1, c2) {
		fmt.Println(n) // 打印 4 9, 或 9 4
	}
}

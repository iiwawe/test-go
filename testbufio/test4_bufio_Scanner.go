package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//test4()
	test5()
}

// 示例：扫描
func test4() {
	// 逗号分隔的字符串，最后一项为空
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 定义匹配函数（查找逗号分隔的字符串）
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if atEOF {
			// 告诉 Scanner 扫描结束。
			return 0, data, bufio.ErrFinalToken
		} else {
			// 告诉 Scanner 没找到匹配项，让 Scan 填充缓存后再次扫描。
			return 0, nil, nil
		}
	}
	// 指定匹配函数
	scanner.Split(onComma)
	// 开始扫描
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	// 检查是否因为遇到错误而结束
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

// 示例：带检查扫描
func test5() {
	const input = "1234 5678 1234567901234567890 90"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 自定义匹配函数
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 获取一个单词
		advance, token, err = bufio.ScanWords(data, atEOF)
		// 判断其能否转换为整数，如果不能则返回错误
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		// 这里包含了 return 0, nil, nil 的情况
		return
	}
	// 设置匹配函数
	scanner.Split(split)
	// 开始扫描
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}

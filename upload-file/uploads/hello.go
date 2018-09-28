package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	// 以\n 为分隔符读取一段内容
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Found an error", err)

	} else {
		input = input[:len(input)-1]
		fmt.Println("Hello,", input)
	}
}

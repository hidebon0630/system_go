package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1 行め
2 行め
3 行め `

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}

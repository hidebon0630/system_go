package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1 行め
2 行め
3 行め `

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}

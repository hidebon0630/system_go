package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC822))
	// 15 Aug 21 19:37 JST

	fmt.Println(now.Format("2006/01/02 03:04:05 MST"))
	// 2021/08/15 07:37:51 JST
}

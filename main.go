package main

import "fmt"

// インターフェースを定義
type Talker interface {
	Talk()
}

// 構造体を宣言
type Greeter struct {
	name string
}

func (g *Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	// インターフェースの型を持つ変数を宣言
	var talker Talker
	// インターフェースを満たす構造体のポインタは代入できる
	talker = &Greeter{"ninja"}
	talker.Talk()
}

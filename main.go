package main

import (
	"fmt"
)

func main() {
	a := byte(12) // 12 は 8ビットのバイナリで 00001100
	b := byte(10) // 10 は 8ビットのバイナリで 00001010

	c := a & b // ビットANDを実行すると 00001000 になる（10進数で8）

	fmt.Printf("%d & %d = %d\n", a, b, c)
}

package main

import (
	"bytes"
	"fmt"
	huf "huffman/t2"
)

func main() {
	bs := []byte("aabbbccccddddde")
	h := huf.NewHuffman()
	bs1, _ := h.Zip(bs)
	bs2, _ := h.Unzip(bs1)
	if bytes.Equal(bs, bs2) {
		fmt.Printf("%s\n", "OK")
	} else {
		fmt.Printf("%s\n", "FAIL")
	}
}

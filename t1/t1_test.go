package t1

import (
	"fmt"
	"testing"
)

func TestGetHuffmanMap2(t *testing.T) {
	s := "aabbbccccddddd"
	hm := GetHuffmanMap([]byte(s))
	for k, v := range hm {
		fmt.Printf("k : %c,v: %v\n", k, v)
	}
}

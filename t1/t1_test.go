package t1

import (
	"fmt"
	"testing"
)

func TestZip1(t *testing.T) {
	bs := []byte("aabbbccccddddd")
	h := NewHuffman()
	bs1, _ := h.Zip(bs)
	fmt.Printf("rate=%.2f\n", (float64(len(bs1)) / float64(len(bs))))
	for k, v := range h.bm {
		fmt.Printf("%c ; %s \n", k, v)
	}
	Print01Bytes(bs1)
	//解压
	bs2, _ := h.Unzip(bs1)
	fmt.Printf("%s\n", string(bs2))

}
func TestSort(t *testing.T) {
	hs := make([]*HTNode, 0)
	hs = append(hs, &HTNode{data: 'a', weight: 2})
	hs = append(hs, &HTNode{data: 'b', weight: 3})
	hs = append(hs, &HTNode{data: 'c', weight: 4})
	hs = append(hs, &HTNode{data: 'd', weight: 5})
	hs = append(hs, &HTNode{data: 'e', weight: 1})
	hs2 := Sort(hs)
	for i, v := range hs2 {
		fmt.Printf("i=%d,d=%c,w=%d\n", i, v.data, v.weight)
	}
	fmt.Printf("\n")
}
func TestGetHuffmanMap2(t *testing.T) {
	s := "aabbbccccddddd"
	h := NewHuffman()
	h.Zip([]byte(s))
}
func TestPrint10(t *testing.T) {
	s := "abc"
	Print01Bytes([]byte(s))
}

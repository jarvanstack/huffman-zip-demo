package t1

import "fmt"

type HuffmanMap map[byte]int
type Node struct {
}
type HuffmanTree struct {
}

//获得 ASC 码的频率然后用
//构建成为 huffman 树,
//通过 huffman 压缩
//通过同样的 huffman 解压
func main() {
	fmt.Printf("%s\n", "h")
}

func GetHuffmanTree(hm HuffmanMap) {

}

//获得 ASC 码的频率然后用
func GetHuffmanMap(data []byte) HuffmanMap {
	m := make(map[byte]int)
	for _, v := range data {
		m[v] += 1
	}
	return m
}

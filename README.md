## huffman 

golang赫夫曼实现文件的压缩和解压

## quick start

```go
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

```

output 


```bash
[root@jarvan huffman]# go run main.go 
OK
[root@jarvan huffman]# 
```

## 思路

```bash
//构建Huffman树和压缩表
//压缩使用压缩表压缩
//解压的是否使用同样的 huffman 树解压
```
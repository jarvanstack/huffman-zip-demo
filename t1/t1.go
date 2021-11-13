package t1

import (
	"bytes"
	"fmt"
)

//获得 ASC 码的频率然后用
//构建成为 huffman 树,
//通过 huffman 压缩
//通过同样的 huffman 解压

type Huffman struct {
	root *HTNode         //huffman 树根节点
	bm   map[byte]string //每个字节对应的bits 比如 a:10
}

type HTNode struct {
	data   byte
	weight int
	code   string
	left   *HTNode
	right  *HTNode
}

//赋值 code 给 ht
func AddCode(h *HTNode, parentCode, addCode string, m map[byte]string) {
	h.code = parentCode + addCode
	if h.left != nil {
		AddCode(h.left, h.code, "0", m)
	}
	if h.right != nil {
		AddCode(h.right, h.code, "1", m)
	}
	if h.IsLeaf() {
		if h.data != 255 {
			//赋值map
			m[h.data] = h.code
		}
	}
}

//是否是叶子节点
func (h *HTNode) IsLeaf() bool {
	if h.left == nil && h.right == nil {
		return true
	}
	return false
}

//打印 中序遍历打印
func (h *HTNode) PrintT() {
	fmt.Printf("data:%c,w:%d,code:%s\n", h.data, h.weight, h.code)
	if h.left != nil {
		h.left.PrintT()
	}
	if h.right != nil {
		h.right.PrintT()
	}
}

func NewHuffman() *Huffman {
	return &Huffman{}
}

//解码一个 byte 如果结束了返回 -1
//01100110 01110111 01110100 10010010 00000000 00000000
func (h *Huffman) Unzip(src []byte) ([]byte, error) {
	//将 src 变为 str 就是字符串 01 那种
	bs := ByteToStr01(src)
	// fmt.Printf("bs: %v\n", string(bs))
	//然后一个读取返回
	resp := bytes.Buffer{}
	si := 0
	for {
		b, i := Get1Byte(bs[si:], h.root, 0)
		si += i
		if i > 0 {
			resp.WriteByte(b)
		}
		if bs[si] == '0' {
			//结束压缩
			break
		}
	}
	return resp.Bytes(), nil
}

//将编码后的二进制文件变成字符串
func ByteToStr01(bs []byte) []byte {
	bss := bytes.Buffer{}
	for _, n := range bs {
		bss.WriteString(fmt.Sprintf("%08b", n)) // prints 00000000 11111101
	}
	return bss.Bytes()
}

//解码一个 byte 如果结束了返回 -1
//遍历这棵树
//返回条件是叶子节点
func Get1Byte(bs []byte, root *HTNode, level int) (byte, int) {
	level++
	if root.IsLeaf() {
		return root.data, level
	}
	if bs[level] == '0' {
		//左子树
		return Get1Byte(bs, root.left, level)
	} else {
		//右子树
		return Get1Byte(bs, root.right, level)
	}
}

//map [byte] string a 001  用于压缩
//huffman 树用于解压
func (h *Huffman) Zip(src []byte) ([]byte, error) {
	//获得 ASC 码的频率然后用
	tm := h.getTimesMap(src)
	//构建成为 huffman 树
	//赋值
	hs := make([]*HTNode, 0)
	for b, i := range tm {
		hs = append(hs, &HTNode{data: b, weight: i, code: ""})
	}
	//构建 huffman 树
	root := h.NewHT(hs)
	//huffman 树的节点用于解码
	h.root = root
	//赋值 code 和 map 表
	m := make(map[byte]string)
	AddCode(root, "", "1", m)
	h.bm = m
	//压缩
	bs := bytes.Buffer{}
	for _, b := range src {
		bs.WriteString(m[b])
	}
	bss := Str01ToBytes(bs.String())
	//按照 0101 打印
	return bss, nil
}
func Print01Bytes(bs []byte) {
	for _, n := range bs {
		fmt.Printf("%08b ", n) // prints 00000000 11111101
	}
}
func Str01ToBytes(str string) []byte {
	//字节数量
	strSize := len(str)
	byteSize := strSize / 8
	if strSize%8 != 0 {
		byteSize++
	}
	bs := make([]byte, byteSize)
	counter := 0
	for i := 0; i < byteSize; i++ {
		for j := 7; j >= 0; j-- {
			if str[counter] == '1' {
				//设置为 1
				bs[i] |= (1 << j)
			}
			counter++
			if counter == strSize {
				break
			}
		}
	}
	return bs
}
func (h *Huffman) NewHT(hs []*HTNode) *HTNode {
	//退出条件
	if len(hs) <= 1 {
		if len(hs) == 0 {
			return nil
		}
		return hs[0]
	}
	hs = Sort(hs)
	//最小的2个成为新的
	n1 := hs[0]
	n2 := hs[1]
	n3 := &HTNode{
		data:   255,
		weight: n1.weight + n2.weight,
		left:   n1,
		right:  n2,
	}
	hs[1] = n3
	return h.NewHT(hs[1:])
}

//a b c d e
//2 3 4 5 1
//0 1 2 3 4
func Sort(hs []*HTNode) []*HTNode {
	//排序
	for i := 0; i < len(hs); i++ {
		maxIndex := 0
		endIndex := len(hs) - i
		for j := 0; j < endIndex; j++ {
			if hs[maxIndex].weight < hs[j].weight {
				maxIndex = j
			}
		}
		//交换
		endIndex--
		if hs[endIndex].weight != hs[maxIndex].weight {
			temp := hs[endIndex]
			hs[endIndex] = hs[maxIndex]
			hs[maxIndex] = temp
		}
	}
	return hs
}

//获得 ASC 码的频率然后用
func (h *Huffman) getTimesMap(data []byte) map[byte]int {
	m := make(map[byte]int)
	for _, v := range data {
		m[v] += 1
	}
	return m
}

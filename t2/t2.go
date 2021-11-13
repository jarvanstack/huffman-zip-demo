package t2

import (
	"bytes"
	"fmt"
)

//构建Huffman树和压缩表
//压缩使用压缩表压缩
//解压的是否使用同样的 huffman 树解压

type Huffman struct {
	root *hTNode         //huffman 树根节点
	bm   map[byte]string //每个字节对应的bits 比如 a:10
}

func NewHuffman() *Huffman {
	return &Huffman{}
}

//压缩
func (h *Huffman) Zip(src []byte) ([]byte, error) {
	//1.获得 ASC 码的频率然后用
	tm := h.getTimesMap(src)
	//2.构建成为 huffman 树
	hs := make([]*hTNode, 0)
	for b, i := range tm {
		hs = append(hs, &hTNode{data: b, weight: i, bits: ""})
	}
	//保存 huffman 根节点
	h.root = h.newHT(hs)
	//赋值 code 和 map 表
	m := make(map[byte]string)
	//获得每一个节点对应的 bit
	valueBits(h.root, "", "1", m)
	//3. 构建 huffman 对应的压缩表 map[byte]string 比如 a:1010
	h.bm = m
	//3. 压缩
	//3.1 获得压缩后的字符串
	bs := bytes.Buffer{}
	for _, b := range src {
		bs.WriteString(m[b])
	}
	//3.2将 "0101" 字符串压缩为 byte
	bss := strToEncodeBytes(bs.Bytes())
	//按照 0101 打印
	return bss, nil
}

//解压
func (h *Huffman) Unzip(src []byte) ([]byte, error) {
	//将编码后的二进制文件变成二进制字符串bytes
	bs := bytesToBS(src)
	resp := bytes.Buffer{}
	si := 0
	for {
		b, i := bytesDecodeToByte(bs[si:], h.root, 0)
		si += i
		if i > 0 {
			//解压一个字节
			resp.WriteByte(b)
		}
		//如果下一个bit从0开始就是结束(我们压缩的是否都是从首位都是1)
		if si == len(bs) || bs[si] == '0' {
			//结束压缩
			break
		}
	}
	return resp.Bytes(), nil
}

//将节点 huffman 树化
func (h *Huffman) newHT(hs []*hTNode) *hTNode {
	//退出条件
	if len(hs) <= 1 {
		if len(hs) == 0 {
			return nil
		}
		return hs[0]
	}
	hs = sort(hs)
	//最小的2个成为新的
	n1 := hs[0]
	n2 := hs[1]
	n3 := &hTNode{
		data:   255,
		weight: n1.weight + n2.weight,
		left:   n1,
		right:  n2,
	}
	hs[1] = n3
	return h.newHT(hs[1:])
}

//获得 ASC 码的频率
func (h *Huffman) getTimesMap(data []byte) map[byte]int {
	m := make(map[byte]int)
	for _, v := range data {
		m[v] += 1
	}
	return m
}

//huffman 树的节点
type hTNode struct {
	data   byte
	weight int
	bits   string
	left   *hTNode
	right  *hTNode
}

//是否是叶子节点
func (h *hTNode) isLeaf() bool {
	if h.left == nil && h.right == nil {
		return true
	}
	return false
}

//测试
//前序遍历打印哈夫曼树
func (h *hTNode) printHT() {
	fmt.Printf("data:%c,w:%d,code:%s\n", h.data, h.weight, h.bits)
	if h.left != nil {
		h.left.printHT()
	}
	if h.right != nil {
		h.right.printHT()
	}
}

//Huffman 树赋值 bits 每个节点的 bits
func valueBits(h *hTNode, parentCode, addCode string, hm map[byte]string) {
	h.bits = parentCode + addCode
	if h.left != nil {
		valueBits(h.left, h.bits, "0", hm)
	}
	if h.right != nil {
		valueBits(h.right, h.bits, "1", hm)
	}
	if h.isLeaf() {
		if h.data != 255 {
			//赋值map
			hm[h.data] = h.bits
		}
	}
}

//将编码后的二进制文件变成二进制字符串
func bytesToBS(bs []byte) []byte {
	bss := bytes.Buffer{}
	for _, n := range bs {
		bss.WriteString(fmt.Sprintf("%08b", n)) // prints 00000000 11111101
	}
	return bss.Bytes()
}

//将二进制字符串解码一次,然后返回解码的字节和解码小号的字符串的长度
func bytesDecodeToByte(bs []byte, root *hTNode, level int) (byte, int) {
	level++
	if root.isLeaf() {
		return root.data, level
	}
	if bs[level] == '0' {
		//左子树
		return bytesDecodeToByte(bs, root.left, level)
	} else {
		//右子树
		return bytesDecodeToByte(bs, root.right, level)
	}
}

// 将 "0101" 字符串压缩为 byte
func strToEncodeBytes(str []byte) []byte {
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

//a b c d e
//2 3 4 5 1
//0 1 2 3 4
func sort(hs []*hTNode) []*hTNode {
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

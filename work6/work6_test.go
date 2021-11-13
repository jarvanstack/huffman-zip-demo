package work6

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	//一个只能装2条数据的 LRU 数据库
	s := NewServer(8080, 2)
	//LRU 数据库运行
	s.Run()
}
func TestClient(t *testing.T) {
	c := NewClient("127.0.0.1:8080")
	c.Set("key1", "value1")
	c.Set("key2", "value2")
	c.Set("key3", "value3")
	//成功 value1
	key2, _ := c.Get("key2")
	fmt.Printf("key2: %v\n", key2)
	//过期并返回错误
	key1, _ := c.Get("key1")
	fmt.Printf("key1: %v\n", key1)
}

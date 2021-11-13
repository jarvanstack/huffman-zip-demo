package work6

//NewClient()创建客户端
//  addr: 服务端的暴露的接口的地址
//  使用 work3 的 rpc 通信
func NewClient(addr string) Client {
	return nil
}

type Client interface {
	//Set() 如果没有就新建,有就更新, 失败返回错误
	Set(key string, value string) error
	//Get() 获取缓存的内容, 失败返回错误
	Get(key string) (string, error)
	//Delete() 删除缓存的内容, 失败返回错误
	Delete(key string) error
}

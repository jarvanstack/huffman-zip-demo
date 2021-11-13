package work6

//NewServer()
//  port: 创建服务开放端口
//  maxSize: 最大条数(这里因为做最大内存量比如 1GB 不方便做,所以简化为数据条数)
func NewServer(port int, maxSize int64) Server {
	return nil
}

type Server interface {
	//开启数据库
	Run()
}

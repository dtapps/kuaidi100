package kuaidi100

import (
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Customer string // 授权码
	Key      string // 密钥
}

// Client 实例
type Client struct {
	config struct {
		customer string // 授权码
		key      string // 密钥
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.customer = config.Customer
	c.config.key = config.Key

	return c, nil
}

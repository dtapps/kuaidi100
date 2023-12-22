package kuaidi100

import (
	"go.dtapp.net/golog"
)

func (c *Client) Config(customer string) *Client {
	c.config.customer = customer
	return c
}

// ConfigApiGormFun 接口日志配置
func (c *Client) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		c.gormLog.client = client
		c.gormLog.status = true
	}
}

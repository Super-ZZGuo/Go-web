package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// 初始对象
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求信息
	Path   string
	Method string
	// 解析路由参数
	Params map[string]string
	// 响应信息
	StatusCode int
	// 中间件
	// HandlerFunc -> 自定义中间件
	// index -> 记录进行到第几个中间件
	handlers []HandlerFunc
	index    int
}

// Context 构造函数
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// Next 将控制权交给下一个中间件
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// Fail 错误处理
func (c *Context) Fail(code int, err string) {
	// inedx = 0 表示第一个中间件
	// 故 c.index = len(c.handlers) 表示跳过中间索引对应的中间件
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) Param(key string) string {
	//value, _ := c.Params[key]
	return c.Params[key]
}

// PostForm 提供访问 PostForm 方法
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 提供访问 Query 方法
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 添加状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置请求头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 快速构建 String 响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 快速构建 JSON 响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 快速构建 Data 响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 快速构建 HTML 响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
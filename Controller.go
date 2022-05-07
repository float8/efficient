package efficient

import (
	"net/http"
	"reflect"
	"strings"
)

// Controller 定义了一些基本的http请求处理程序操作，例如
// http上下文，模板和视图，会话和xsrf。
type Controller struct {
	//控制器名称
	controllerName string
	//组名
	groupName string
}

//ControllerInterface 是一个统一所有控制器处理程序的接口。
type ControllerInterface interface {
	Prepare(ctx Context)
	Get(ctx Context)
	Post(ctx Context)
	Delete(ctx Context)
	Put(ctx Context)
	Head(ctx Context)
	Patch(ctx Context)
	Options(ctx Context)
	Execute(ctx Context, ctr ControllerInterface)
	Finish(ctx Context)
}

// Prepare 在请求函数执行之前，在Init之后运行。
func (c *Controller) Prepare(cxt Context) {}

// Finish 在请求函数执行后运行。
func (c *Controller) Finish(cxt Context) {}

// Get 添加一个请求函数来处理GET请求。
func (c *Controller) Get(cxt Context) {
	http.Error(cxt.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Post 添加一个请求函数来处理POST请求。
func (c *Controller) Post(cxt Context) {
	http.Error(cxt.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Delete 添加一个请求函数来处理DELETE请求
func (c *Controller) Delete(cxt Context) {
	http.Error(cxt.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Put  添加一个请求函数来处理PUT请求
func (c *Controller) Put(cxt Context) {
	http.Error(cxt.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Put  添加一个请求函数来处理 HEAD 请求
func (c *Controller) Head(cxt Context) {
	http.Error(cxt.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Patch  添加一个请求函数来处理 PATCH 请求
func (c *Controller) Patch(cxt Context) {
	http.Error(cxt.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Options  添加一个请求函数来处理 OPTIONS 请求
func (c *Controller) Options(cxt Context) {
	http.Error(cxt.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}

//Execute 执行请求方法
func (c *Controller) Execute(cxt Context, ctr ControllerInterface) {
	switch cxt.Request.Method {
	case http.MethodGet:
		ctr.Get(cxt)
	case http.MethodPost:
		ctr.Post(cxt)
	case http.MethodDelete:
		ctr.Delete(cxt)
	case http.MethodPut:
		ctr.Put(cxt)
	case http.MethodHead:
		ctr.Head(cxt)
	case http.MethodPatch:
		ctr.Patch(cxt)
	case http.MethodOptions:
		ctr.Options(cxt)
	default:
		http.Error(cxt.Writer, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// Group 添加一个请求函数来处理 Group 请求。
func (c *Controller) Group() {}

// Use 添加一个中间件。
func (c *Controller) Use() {}

//GetGroup 获取执行的组名称。
func (c *Controller) GetGroup() string {
	return c.groupName
}

//GetController 获取执行的控制器名称。
func (c *Controller) GetController(ctr ControllerInterface) string {
	if c.controllerName == "" {
		reflectVal := reflect.ValueOf(ctr)
		ct := reflect.Indirect(reflectVal).Type()
		c.controllerName = strings.TrimSuffix(ct.Name(), "Controller")
	}
	return c.controllerName
}

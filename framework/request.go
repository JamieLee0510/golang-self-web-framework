package framework

import (
	"mime/multipart"

	"github.com/spf13/cast"
)

//代表請求包含的方法
type IRequest interface{
	//請求地址中帶的參數
	//ex： foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}

	// 路由匹配中帶的參數
	// ex： /book/:id
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	Param(key string) interface{}

	//form表單中帶的參數
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	FormFile(key string) (*multipart.FileHeader, error)
	Form(key string) interface{}


	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)
	
	// request 的基礎資訊
	Uri() string
	Method() string
	Host() string
	ClientIp() string
	
	// header
	Headers() map[string][]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)

}


//獲取請求地址的所有參數
func (ctx *Context) QueryAll() map[string][]string{
	if ctx.request != nil{
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
}


// 請求地址uri中帶的參數
// ex: foo.com?a=1&b=bar&c[]=bar
// #region

// 獲取int類型的參數
func (ctx *Context) QueryInt(key string, def int) (int, bool){
	params := ctx.QueryAll()
	if vals, ok := params["key"]; ok{
		if len(vals) >0{
			//透過cast庫將string 轉換為 int
			return cast.ToInt(vals[0]), true
		}
	}
	return def,false
}

func (ctx *Context) QueryInt64(key string, def int64) (int64, bool){
	params := ctx.QueryAll()
	if vals, ok := params["key"]; ok{
		if len(vals) >0{
			//透過cast庫將string 轉換為 int
			return cast.ToInt64(vals[0]), true
		}
	}
	return def,false
}


func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) Query(key string) interface{} {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals[0]
	}
	return nil
}

// #endregion

// 路由匹配中帶的參數
// ex: /book/:id



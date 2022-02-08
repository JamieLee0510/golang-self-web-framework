/// self packaging request

package gin

import (
	"mime/multipart"

	"github.com/spf13/cast"
)

//代表請求包含的方法
type IRequest interface{
	//請求地址中帶的參數
	//ex： foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中帶的參數
	// ex： /book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	//form表單中帶的參數
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}


	// // json body
	// BindJson(obj interface{}) error

	// // xml body
	// BindXml(obj interface{}) error

	// // 其他格式
	// GetRawData() ([]byte, error)
	
	// // request 的基礎資訊
	// Uri() string
	// Method() string
	// Host() string
	// ClientIp() string
	
	// // header
	// Headers() map[string][]string
	// Header(key string) (string, bool)

	// // cookie
	// Cookies() map[string]string
	// Cookie(key string) (string, bool)

}


//獲取請求地址的所有參數
func (ctx *Context) QueryAll() map[string][]string{
	if ctx.Request != nil{
		return map[string][]string(ctx.Request.URL.Query())
	}
	return map[string][]string{}
}


// 請求地址uri中帶的參數
// ex: foo.com?a=1&b=bar&c[]=bar
// #region 請求的相關接口實現 

// 獲取int類型的參數
func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool){
	params := ctx.QueryAll()
	if vals, ok := params["key"]; ok{
		if len(vals) >0{
			//透過cast庫將string 轉換為 int
			return cast.ToInt(vals[0]), true
		}
	}
	return def,false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool){
	params := ctx.QueryAll()
	if vals, ok := params["key"]; ok{
		if len(vals) >0{
			//透過cast庫將string 轉換為 int
			return cast.ToInt64(vals[0]), true
		}
	}
	return def,false
}


func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}


// #endregion


// #region params related

//獲取路由參數
func (ctx *Context) DingParam(key string) interface{}{
	params := ctx.QueryAll()
	if ctx.params != nil{
		if val, ok := params[key]; ok{
			return val
		}
	}

	return nil
}

// 路由匹配中帶的參數
// ex: /book/:id
func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := ctx.DingParam(key); val != nil {
		// 通過cast進行類型轉換
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.DingParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.DingParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.DingParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) ParamBool(key string, def bool) (bool, bool) {
	if val := ctx.DingParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	if val := ctx.DingParam(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

// #endregion

// #region bind body related

package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
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
// #region 請求的相關接口實現 

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


// #region params related

//獲取路由參數
func (ctx *Context) Param(key string) interface{}{
	if ctx.params != nil{
		if val, ok := ctx.params[key]; ok{
			return val
		}
	}

	return nil
}

// 路由匹配中帶的參數
// ex: /book/:id
func (ctx *Context) ParamInt(key string, def int) (int, bool) {
	if val := ctx.Param(key); val != nil {
		// 通過cast進行類型轉換
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) ParamBool(key string, def bool) (bool, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

// #endregion

// #region bind body related

//將 body 解析到 obj 結構體中
func (ctx *Context) BindJson(obj interface{}) error{
	if ctx.request != nil{
		//使用 ioutil 讀取 body 中的文本
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil{
			return err
		}

		//重新複製一份body到ctx.request裡面，為後續的邏輯處理做準備
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		//解析body到obj結構體中
		err = json.Unmarshal(body,obj)
		if err != nil { 
			return err 
		}
	} else {
		return errors.New("ctx.request is empty")
	}
	return nil
}

// xml body
func (ctx *Context) BindXml(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = xml.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request is empty")
	}
	return nil
}

// #endregion



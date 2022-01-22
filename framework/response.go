package framework

type IResponse interface{
	///很多返回值使用IResponse接口本身，這樣可以使用鏈式調用
	///ex: c.SetOkStatus().Json("ok, UserLoginController: " + foo)

	//Json 輸出
	Json(obj interface{}) IResponse

	// Jsonp 輸出
	Jsonp(obj interface{}) IResponse
	
	// xml 輸出
	xml(obj interface{}) IResponse

	// html 輸出
	Html(template string,obj interface{}) IResponse

	// string
	Text(format string, values ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 設置 status
	SetStatus(code int) IResponse

	// 設置200 status
	SetOkStatus() IResponse


}
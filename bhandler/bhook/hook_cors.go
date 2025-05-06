package bhook

import "github.com/gogf/gf/v2/net/ghttp"

// BambooHookDefaultCors 为请求启用默认的跨域配置并继续执行中间件链。
func BambooHookDefaultCors(r *ghttp.Request) {
	r.Response.CORSDefault()
}

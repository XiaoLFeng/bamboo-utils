package bmiddle

import "github.com/gogf/gf/v2/net/ghttp"

// BambooMiddleDefaultCors
//
// # 默认跨域处理
//
// 默认跨域处理，用于处理跨域请求的处理，设置默认的跨域请求头信息；
// 默认的跨域处理，将会放行所有的跨域请求，不做任何的限制处理；
//
// # 参数
//   - r		http 信息(*ghttp.Request)
func BambooMiddleDefaultCors(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

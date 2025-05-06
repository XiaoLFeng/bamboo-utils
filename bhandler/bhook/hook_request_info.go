/*
 * ***********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ***********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ***********************************************************
 */

package bhook

import (
	"github.com/bamboo-services/bamboo-utils/blog"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// BambooHookRequestInfo 记录并解析 HTTP 请求的详细信息，包括请求方法、路径、参数、头部等内容，并统计请求耗时。
//
// 该钩子函数在请求处理之前被调用，用于记录并解析 HTTP 请求的详细信息，包括请求方法、路径、参数、头部等内容，并统计请求耗时。
//
// 参数：
//   - r：ghttp.Request 对象，表示当前 HTTP 请求。
//
// 返回值：无
func BambooHookRequestInfo(r *ghttp.Request) {
	// 获取请求的参数
	blog.BambooInfo(r.Context(), "REQU", "请求信息：[%s] %s", r.Method, r.URL.Path)
	if len(r.GetBody()) > 0 {
		g.Log().Infof(r.Context(), "\t[BODY]请求体参数:")
		// 解析请求体参数
		decode, _ := gjson.Decode(r.GetBody())
		for key, value := range decode.(map[string]interface{}) {
			g.Log().Infof(r.Context(), "\t\t[%v] \t%v", key, value)
		}
	}
	if len(r.GetQueryMap()) > 0 {
		g.Log().Infof(r.Context(), "\t[PARA]请求参数:")
		for key, value := range r.GetQueryMap() {
			g.Log().Infof(r.Context(), "\t\t[%v] \t%v", key, value)
		}
	}
	if r.Request.Header != nil {
		g.Log().Debugf(r.Context(), "\t[HEAD]请求头部:")
		for key, value := range r.Request.Header {
			g.Log().Debugf(r.Context(), "\t\t[%v] \t%v", key, value)
		}
	}

	r.Middleware.Next()
}

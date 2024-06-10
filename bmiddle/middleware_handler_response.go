package bmiddle

import (
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/btype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"net/http"
)

// BambooMiddleHandler
//
// # 基本数据返回
//
// 基本数据信息返回，用于最终的报错或者数据进行编译打包，打包完成后输出 json 参数信息.
//
// # 参数
//   - r		http 信息(*ghttp.Request)
func BambooMiddleHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 自定义缓冲内容，退出默认信息返回
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = bcode.BCode(bcode.Success)
	)
	if r.GetError() != nil {
		// 对 GoFrame 默认的报错作处理
		if gerror.Code(r.GetError()) != gcode.CodeNil {
			switch gerror.Code(r.GetError()).Code() {
			case 51:
				code = bcode.RequestParameterIncorrect
			}
			msg = r.GetError().Error()
		} else {
			// 自定义处理
			if berror.GetECode(err) == nil {
				code = bcode.ServerInternalError
			}
			code = *berror.GetECode(err)
			msg = err.Error()
		}
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = bcode.PageNotFound
			case http.StatusForbidden:
				code = bcode.StatusForbidden
			default:
				code = bcode.UnknownError
			}
			err = berror.NewError(code, msg)
			r.SetError(err)
		} else {
			code = bcode.Success
		}
	}

	if g.IsNil(res) {
		r.Response.WriteJson(btype.BaseResponse{
			Output:       code.Output(),
			Code:         code.Code(),
			Message:      code.Message(),
			ErrorMessage: msg,
		})
	} else {
		r.Response.WriteJson(btype.BaseResponse{
			Output:       code.Output(),
			Code:         code.Code(),
			Message:      code.Message(),
			ErrorMessage: msg,
			Data:         res,
		})
	}
	if msg != "" {
		glog.Noticef(r.Context(), "[RESPONSE] <%d>%s %s - %s", code.Code(), code.Output(), code.Message(), msg)
	} else {
		glog.Noticef(r.Context(), "[RESPONSE] <%d>%s %s", code.Code(), code.Output(), code.Message())
	}
}

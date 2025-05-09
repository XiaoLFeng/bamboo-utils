package bmiddle

import (
	"errors"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/XiaoLFeng/bamboo-utils/butil/private"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"go/types"
	"net/http"
)

// BambooHandlerResponse 处理 HTTP 请求的统一响应逻辑。
//
// 根据请求的中间件处理结果返回统一的 JSON 格式响应。当存在错误或未定义状态时，会构造适配的错误码和消息返回。
//
// 参数:
//   - r: 当前 HTTP 请求对象。
//
// 功能描述:
//   - 检查缓存区内容，若存在则直接返回内容为 JSON 格式。
//   - 处理自定义或未定义的错误并返回错误状态码和信息。
//   - 在无错误且正常状态下返回标准成功响应。
func BambooHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// 检查如果缓存区存在内容直接返回结果
	if r.Response.BufferLength() > 0 {
		r.Response.WriteJson(r.Response.BufferString())
		return
	}

	var errorCode berror.ErrorCode

	// 否则获取是否有错误
	if r.GetError() != nil {
		// 检查报错是否可以转换为自定义错误码
		var customErr berror.Error
		if errors.As(r.GetError(), &customErr) {
			// 自定义错误码
			errorCode = customErr.GetErrorCode()
		} else {
			// 检查是否是原生 GoFrame 错误
			var normalError *gerror.Error
			if errors.As(r.GetError(), &normalError) {
				errorCode = private.GErrorTransform(normalError)
			} else {
				errorCode = berror.ErrDeveloperError
			}
		}
		if errorCode.GetErrorCode().Code > 9999 {
			r.Response.Status = int(errorCode.GetErrorCode().Code / 100)
		} else if errorCode.GetErrorCode().Code > 999 {
			r.Response.Status = int(errorCode.GetErrorCode().Code / 10)
		} else {
			r.Response.Status = int(errorCode.GetErrorCode().Code)
		}
		returnResult := bmodels.ResponseDTO[types.Nil]{
			Context: gctx.CtxId(r.GetCtx()),
			Time:    gtime.TimestampMilli(),
			Code:    errorCode.Code,
			Message: errorCode.Message,
		}
		if g.Log().GetLevel() == glog.LEVEL_DEV {
			returnResult.Overhead = butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds())
		}
		r.Response.WriteJson(returnResult)
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			errorCode = berror.ErrUndefinitionError
			errorCode.Data = r.Response.Buffer()
			returnResult := bmodels.ResponseDTO[types.Nil]{
				Context: gctx.CtxId(r.GetCtx()),
				Time:    gtime.TimestampMilli(),
				Code:    errorCode.Code,
				Message: errorCode.Message,
			}
			if g.Log().GetLevel() == glog.LEVEL_DEV {
				returnResult.Overhead = butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds())
			}
			r.Response.Status = int(berror.ErrUndefinitionError.Code / 100)
			r.Response.WriteJson(returnResult)
			return
		}
	}

	// 返回统一格式的响应
	if !g.IsNil(r.GetHandlerResponse()) {
		r.Response.WriteJson(r.GetHandlerResponse())
	} else {
		returnResult := bmodels.ResponseDTO[types.Nil]{
			Context: gctx.CtxId(r.GetCtx()),
			Time:    gtime.TimestampMilli(),
			Code:    200,
			Message: "操作成功",
		}
		if g.Log().GetLevel() == glog.LEVEL_DEV {
			returnResult.Overhead = butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds())
		}
		r.Response.WriteJson(returnResult)
	}
}

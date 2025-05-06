package bresult

import (
	"context"
	"github.com/bamboo-services/bamboo-utils/bmodels"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"go/types"
)

// SuccessHasData 创建一个包含成功状态码、消息和数据的响应结构体。
//
// 参数:
//   - ctx: 请求的上下文消息对象。
//   - message: 响应的描述信息。
//   - data: 要返回的具体数据内容。
//
// 返回:
//   - 包含状态码、描述信息及数据的通用响应结构体。
func SuccessHasData[T interface{}](ctx context.Context, message string, data T) *bmodels.ResponseDTO[T] {
	returnResult := bmodels.ResponseDTO[T]{
		Context: gctx.CtxId(ctx),
		Code:    200,
		Data:    data,
		Time:    gtime.TimestampMilli(),
		Message: message,
	}
	if g.Log().GetLevel() == glog.LEVEL_DEV {
		request := ghttp.RequestFromCtx(ctx)
		returnResult.Overhead = butil.Ptr(gtime.Now().Sub(request.EnterTime).Milliseconds())
	}
	return &returnResult
}

// Success 创建一个成功的响应对象。
//
// 参数:
//   - ctx: 请求的上下文对象。
//   - message: 响应的消息内容。
//
// 返回:
//   - dto.ResponseDTO[types.Nil]: 包含状态码 200 和消息的响应对象。
func Success(ctx context.Context, message string) *bmodels.ResponseDTO[types.Nil] {
	returnResult := bmodels.ResponseDTO[types.Nil]{
		Context: gctx.CtxId(ctx),
		Code:    200,
		Time:    gtime.TimestampMilli(),
		Message: message,
	}
	if g.Log().GetLevel() == glog.LEVEL_DEV {
		request := ghttp.RequestFromCtx(ctx)
		returnResult.Overhead = butil.Ptr(gtime.Now().Sub(request.EnterTime).Milliseconds())
	}
	return &returnResult
}

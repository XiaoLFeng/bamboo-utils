package bresult

import (
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/btype"
	"github.com/gogf/gf/v2/net/ghttp"
)

// WriteResponse
//
// # 写入响应
//
// 用于写入响应信息, 用于返回错误码信息.
//
// # 参数
//   - ctx			请求的上下文(*context.Context)
//   - getCode		错误码(BCode)
//   - msg			错误信息(string)
//   - res			数据(interface)
func WriteResponse(ctx context.Context, getCode bcode.BCode, msg string, res interface{}) {
	r := ghttp.RequestFromCtx(ctx)
	r.Response.WriteJson(btype.BaseResponse{
		Output:       getCode.Output(),
		Code:         getCode.Code(),
		Message:      getCode.Message(),
		ErrorMessage: msg,
		Data:         res,
	})
}

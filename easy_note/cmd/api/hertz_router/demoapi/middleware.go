// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by hertz generator.

package demoapi

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/easy_note/cmd/api/mw"
	"github.com/cloudwego/biz-demo/easy_note/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/requestid"
	"go.opentelemetry.io/otel/trace"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use recovery mw
		recovery.Recovery(recovery.WithRecoveryHandler(
			func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
				hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
				c.JSON(consts.StatusInternalServerError, utils.H{
					"code":    errno.ServiceErr.ErrCode,
					"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
				})
			},
		)),
		// use requestid mw
		requestid.New(
			requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
				traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
				return traceID
			}),
		),
		// use gzip mw
		gzip.Gzip(gzip.DefaultCompression),
	}
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _noteMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _updatenoteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletenoteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _querynoteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createnote_nduserMw() []app.HandlerFunc {
	// your code...
	//return []app.HandlerFunc{
	//	mw.SeataTransactionMiddleware(),
	//}
	return nil
}

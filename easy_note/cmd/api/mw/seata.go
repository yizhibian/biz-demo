package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/seata/seata-go/pkg/client"
	"github.com/seata/seata-go/pkg/constant"
	"github.com/seata/seata-go/pkg/tm"
)

func InitSeataClient() {

	client.InitPath("pkg/configs/seata/seatago.yml")

}

func SeataTransactionMiddleware() app.HandlerFunc {

	return func(c context.Context, ctx *app.RequestContext) {
		xid := string(ctx.GetHeader(constant.XidKey))
		if xid == "" {
			xid = string(ctx.GetHeader(constant.XidKeyLowercase))
		}

		if len(xid) == 0 {
			hlog.Errorf("Hertz: header not contain header: %s, global transaction xid", constant.XidKey)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c = tm.InitSeataContext(c)
		tm.SetXID(c, xid)
		ctx.Next(c)

		hlog.Infof("global transaction xid is :%s", xid)
	}
}

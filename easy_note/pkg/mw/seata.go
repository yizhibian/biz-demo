package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/seata/seata-go/pkg/client"
	"github.com/seata/seata-go/pkg/constant"
	"github.com/seata/seata-go/pkg/tm"
)

func RMInitSeataClient() {

	//cfg := client.LoadPath("pkg/configs/seata/seatago.yml")
	client.InitPath("pkg/configs/seata/seatago.yml")

}

func SeataTransactionMiddleware() app.HandlerFunc {

	return func(c context.Context, ctx *app.RequestContext) {
		xid := string(ctx.GetHeader(constant.XidKey))
		if xid == "" {
			xid = string(ctx.GetHeader(constant.XidKeyLowercase))
		}

		if len(xid) == 0 {
			hlog.Errorf("Gin: header not contain header: %s, global transaction xid", constant.XidKey)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c = tm.InitSeataContext(c)
		tm.SetXID(c, xid)
		ctx.Next(c)

		hlog.Infof("global transaction xid is :%s", xid)
	}
}

func RMSeataTransactionMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(c context.Context, req, resp interface{}) (err error) {
		//xid := string(ctx.GetHeader(constant.XidKey))
		//if xid == "" {
		//	xid = string(ctx.GetHeader(constant.XidKeyLowercase))
		//}
		//
		//if len(xid) == 0 {
		//	hlog.Errorf("Gin: header not contain header: %s, global transaction xid", constant.XidKey)
		//	ctx.AbortWithStatus(http.StatusBadRequest)
		//	return
		//}

		//c.Value(tm.seataContextVariable)
		//
		//c = tm.InitSeataContext(c)
		//tm.SetXID(c, xid)
		//ctx.Next(c)
		c = tm.InitSeataContext(c)
		xid := tm.GetXID(c)

		klog.Info("the global xid is ", xid)

		if err = next(c, req, resp); err != nil {
			return err
		}

		return nil
		//hlog.Infof("global transaction xid is :%s", xid)
	}
}

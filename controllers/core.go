package controllers

import (
	"github.com/andela/micro-api-gateway-qroute/pb/core"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Core struct {
	conn   *grpc.ClientConn
	client core.CoreServiceClient
}

func (ct *Core) CreateServiceVendor() func(echo.Context) error {
	fn := func(c echo.Context) error {
		var payload core.CreateVendorRequest
		c.Bind(&payload)
		ctx := c.Get("context").(context.Context)

		return ct.client.CreateServiceVendor(ctx, &payload)
	}
	return fn
}

func CoreResource() *Core {
	conn := connectTo("core")
	client := pulse.NewCoreServiceClient(conn)
	return &Core{conn, client}
}

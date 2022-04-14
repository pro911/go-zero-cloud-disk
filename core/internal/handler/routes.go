// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/pro911/go-zero-cloud-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: CoreHandler(serverCtx),
			},
		},
	)
}

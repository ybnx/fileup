// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"fileup/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/emaillogin",
				Handler: EmailLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/passwdlogin",
				Handler: PasswdLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sendcode",
				Handler: SendCodeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/changePasswd",
				Handler: ChangePasswdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/signout",
				Handler: SignoutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/userinfo",
				Handler: UserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}

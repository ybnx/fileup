package handler

import (
	"net/http"

	"fileup/api/user/internal/logic"
	"fileup/api/user/internal/svc"
	"fileup/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChangePasswdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangePasswdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChangePasswdLogic(r.Context(), svcCtx)
		resp, err := l.ChangePasswd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

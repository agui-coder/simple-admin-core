package api

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/api"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /api/create_or_update api CreateOrUpdateApi
//
// Create or update API information | 创建或更新API
//
// Create or update API information | 创建或更新API
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateApiReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := api.NewCreateOrUpdateApiLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateApi(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

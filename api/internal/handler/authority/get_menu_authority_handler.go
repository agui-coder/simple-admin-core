package authority

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/authority"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /authority/menu/role authority GetMenuAuthority
//
// Get role's menu authorization list | 获取角色菜单权限列表
//
// Get role's menu authorization list | 获取角色菜单权限列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: MenuAuthorityInfoResp

func GetMenuAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := authority.NewGetMenuAuthorityLogic(r, svcCtx)
		resp, err := l.GetMenuAuthority(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

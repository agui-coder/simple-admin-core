package oauthprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
)

type DeleteOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOauthProviderLogic {
	return &DeleteOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOauthProviderLogic) DeleteOauthProvider(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.OauthProvider.Delete().Where(oauthprovider.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
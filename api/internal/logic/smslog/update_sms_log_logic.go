package smslog

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSmsLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSmsLogLogic {
	return &UpdateSmsLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSmsLogLogic) UpdateSmsLog(req *types.SmsLogInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.UpdateSmsLog(l.ctx,
		&mcms.SmsLogInfo{
			Id:          req.Id,
			PhoneNumber: req.PhoneNumber,
			Content:     req.Content,
			SendStatus:  req.SendStatus,
			Provider:    req.Provider,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}

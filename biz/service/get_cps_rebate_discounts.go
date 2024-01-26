package service

import (
	"context"
	"howell/howell_api/biz/model/coder/hao/howell_api"
	"howell/howell_api/biz/model/models"
	"howell/howell_api/biz/rpc"
	"howell/howell_api/kitex_gen/coder/hao/howell_rpc"
)

func GetCpsRebateDiscounts(ctx context.Context, req howell_api.GetCpsRebateDiscountsRequest) (data *howell_api.GetCpsRebateDiscountsData, err error) {
	rpcReq := &howell_rpc.MGetCpsRebateDiscountsRequest{
		EntityIdList: []string{req.GetEntityID()},
	}

	entityMap, err := rpc.MGetCpsRebateDiscounts(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	data = &howell_api.GetCpsRebateDiscountsData{}
	entity := entityMap[req.GetEntityID()]
	if entity != nil {
		apiEntity := &models.CpsRebateDiscounts{
			ID:    entity.Id,
			AppID: entity.AppId,
			Name:  entity.Name,
			CpsType: models.CpsTypePtr(
				models.CpsType(entity.GetCpsType())),
			JumpLink: entity.JumpLink,
			Extra:    entity.Extra,
			Status:   entity.Status,
		}
		data.Info = apiEntity
	}

	return data, nil
}

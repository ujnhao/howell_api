package service

import (
	"context"
	"howell/howell_api/biz/model/coder/hao/howell_api"
	api_models "howell/howell_api/biz/model/models"
	"howell/howell_api/biz/rpc"
	"howell/howell_api/kitex_gen/coder/hao/howell_rpc"
	rpc_models "howell/howell_api/kitex_gen/models"
)

func CreateCpsRebateDiscounts(ctx context.Context, req *api_models.CpsRebateDiscounts) (data *howell_api.CreateCpsRebateDiscountsData, err error) {
	rpcCRD := &rpc_models.CpsRebateDiscounts{
		Id:       req.ID,
		AppId:    req.AppID,
		Name:     req.Name,
		CpsType:  rpc_models.CpsTypePtr(rpc_models.CpsType(req.GetCpsType())),
		JumpLink: req.JumpLink,
		Extra:    req.Extra,
		Status:   req.Status,
	}
	rpcReq := &howell_rpc.CreateCpsRebateDiscountsRequest{
		CRDEntity: rpcCRD,
	}
	resp, err := rpc.CreateCpsRebateDiscounts(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &howell_api.CreateCpsRebateDiscountsData{
		EntityID: resp.EntityId,
	}, nil
}

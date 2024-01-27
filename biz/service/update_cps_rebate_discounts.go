package service

import (
	"context"
	"howell/howell_api/biz/model/coder/hao/howell_api"
	"howell/howell_api/biz/model/converters"
	api_models "howell/howell_api/biz/model/models"
	"howell/howell_api/biz/rpc"
	"howell/howell_api/kitex_gen/coder/hao/howell_rpc"
)

func CreateCpsRebateDiscounts(ctx context.Context, req *api_models.CpsRebateDiscounts) (data *howell_api.CreateCpsRebateDiscountsData, err error) {
	rpcCRD := converters.ConvCRDAPIToRPC(req)
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

func UpdateCpsRebateDiscounts(ctx context.Context, req howell_api.UpdateCpsRebateDiscountsRequest) error {
	rpcCRD := converters.ConvCRDAPIToRPC(req.CrdEntity)
	rpcReq := &howell_rpc.UpdateCpsRebateDiscountsRequest{
		EntityId:  req.GetEntityID(),
		CRDEntity: rpcCRD,
	}
	err := rpc.UpdateCpsRebateDiscounts(ctx, rpcReq)
	if err != nil {
		return err
	}

	return nil
}

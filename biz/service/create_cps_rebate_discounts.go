package service

import (
	"context"
	"howell/howell_api/biz/model/coder/hao/howell_api"
	"howell/howell_api/biz/rpc"
	"howell/howell_api/kitex_gen/coder/hao/howell_rpc"
)

func CreateCpsRebateDiscounts(ctx context.Context, req *howell_api.CpsRebateDiscounts) (data *howell_api.CreateCpsRebateDiscountsData, err error) {
	rpcCRD := &howell_rpc.CpsRebateDiscounts{}
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

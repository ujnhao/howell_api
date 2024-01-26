package service

import (
	"context"
	"howell/howell_api/biz/model/coder/hao/howell_api"
	"howell/howell_api/biz/model/converters"
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
	data.Info = converters.ConvCRDRPCToAPI(entity)
	return data, nil
}

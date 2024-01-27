package service

import (
	"context"
	"howell/howell_api/biz/model/coder/hao/howell_api"
	api_common "howell/howell_api/biz/model/common"
	"howell/howell_api/biz/model/converters"
	api_models "howell/howell_api/biz/model/models"
	"howell/howell_api/biz/rpc"
	"howell/howell_api/kitex_gen/coder/hao/howell_rpc"
)

func QueryCpsRebateDiscounts(ctx context.Context, req howell_api.QueryCpsRebateDiscountsRequest) (data *howell_api.QueryCpsRebateDiscountsData, err error) {
	rpcReq := &howell_rpc.QueryCpsRebateDiscountsRequest{
		EntityIdList: req.EntityIDList,
		PageIndex:    req.PageIndex,
		PageSize:     req.PageSize,
	}
	entityList, pagination, err := rpc.QueryCpsRebateDiscounts(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	data = &howell_api.QueryCpsRebateDiscountsData{}
	apiItemList := make([]*api_models.CpsRebateDiscounts, 0)
	for _, entity := range entityList {
		apiItem := converters.ConvCRDRPCToAPI(entity)
		apiItemList = append(apiItemList, apiItem)
	}
	data.ItemList = apiItemList
	data.Pagination = (*api_common.Pagination)(pagination)
	return data, nil
}

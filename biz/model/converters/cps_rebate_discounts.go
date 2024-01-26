package converters

import (
	api_models "howell/howell_api/biz/model/models"
	rpc_models "howell/howell_api/kitex_gen/models"
)

func ConvCRDRPCToAPI(rpcInfo *rpc_models.CpsRebateDiscounts) *api_models.CpsRebateDiscounts {
	if rpcInfo == nil {
		return nil
	}
	apiEntity := &api_models.CpsRebateDiscounts{
		ID:    rpcInfo.Id,
		AppID: rpcInfo.AppId,
		Name:  rpcInfo.Name,
		CpsType: api_models.CpsTypePtr(
			api_models.CpsType(rpcInfo.GetCpsType())),
		JumpLink: rpcInfo.JumpLink,
		Extra:    rpcInfo.Extra,
		Status:   rpcInfo.Status,
	}
	return apiEntity
}

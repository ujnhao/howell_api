package converters

import (
	api_common "howell/howell_api/biz/model/common"
	api_models "howell/howell_api/biz/model/models"
	rpc_models "howell/howell_api/kitex_gen/models"
)

func ConvCRDRPCToAPI(rpcInfo *rpc_models.CpsRebateDiscounts) *api_models.CpsRebateDiscounts {
	if rpcInfo == nil {
		return nil
	}
	apiEntity := &api_models.CpsRebateDiscounts{
		ID:       rpcInfo.Id,
		AppID:    rpcInfo.AppId,
		Name:     rpcInfo.Name,
		CpsType:  (*api_common.CpsType)(rpcInfo.CpsType),
		ActTpye:  (*api_common.ActType)(rpcInfo.ActTpye),
		ActURL:   rpcInfo.ActUrl,
		Images:   rpcInfo.Images,
		Resource: rpcInfo.Resource,
		Extra:    rpcInfo.Extra,
	}
	return apiEntity
}

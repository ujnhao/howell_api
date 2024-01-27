package converters

import (
	api_common "howell/howell_api/biz/model/common"
	api_models "howell/howell_api/biz/model/models"
	rpc_common "howell/howell_api/kitex_gen/common"
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
		ActType:  (*api_common.ActType)(rpcInfo.ActType),
		ActURL:   rpcInfo.ActUrl,
		Images:   rpcInfo.Images,
		Resource: rpcInfo.Resource,
		Extra:    rpcInfo.Extra,
	}
	return apiEntity
}

func ConvCRDAPIToRPC(apiInfo *api_models.CpsRebateDiscounts) *rpc_models.CpsRebateDiscounts {
	if apiInfo == nil {
		return nil
	}
	apiEntity := &rpc_models.CpsRebateDiscounts{
		Id:       apiInfo.ID,
		AppId:    apiInfo.AppID,
		Name:     apiInfo.Name,
		CpsType:  (*rpc_common.CpsType)(apiInfo.CpsType),
		ActType:  (*rpc_common.ActType)(apiInfo.ActType),
		ActUrl:   apiInfo.ActURL,
		Images:   apiInfo.Images,
		Resource: apiInfo.Resource,
		Extra:    apiInfo.Extra,
	}
	return apiEntity
}

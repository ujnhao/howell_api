package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"howell/howell_api/kitex_gen/coder/hao/howell_rpc"
	"howell/howell_api/kitex_gen/coder/hao/howell_rpc/howellrpcservice"
	"howell/howell_api/kitex_gen/common"
	"howell/howell_api/kitex_gen/models"
	"time"
)

func CreateCpsRebateDiscounts(ctx context.Context, req *howell_rpc.CreateCpsRebateDiscountsRequest) (resp *howell_rpc.CreateCpsRebateDiscountsResponse, err error) {
	c, err := howellrpcservice.NewClient("coder.hao.howell_rpc", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		return nil, err
	}

	resp, err = c.CreateCpsRebateDiscounts(ctx, req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func MGetCpsRebateDiscounts(ctx context.Context, req *howell_rpc.MGetCpsRebateDiscountsRequest) (res map[string]*models.CpsRebateDiscounts, err error) {
	c, err := howellrpcservice.NewClient("coder.hao.howell_rpc", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		return nil, err
	}

	resp, err := c.MGetCpsRebateDiscounts(ctx, req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		return nil, err
	}

	return resp.EntityMap, nil
}

func QueryCpsRebateDiscounts(ctx context.Context, req *howell_rpc.QueryCpsRebateDiscountsRequest) ([]*models.CpsRebateDiscounts, *common.Pagination, error) {
	c, err := howellrpcservice.NewClient("coder.hao.howell_rpc", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.QueryCpsRebateDiscounts(ctx, req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		return nil, nil, err
	}

	return resp.ItemList, resp.Pagination, nil
}

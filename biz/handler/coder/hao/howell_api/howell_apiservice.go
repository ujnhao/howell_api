// Code generated by hertz generator.

package howell_api

import (
	"context"
	howell_api "howell/howell_api/biz/model/coder/hao/howell_api"
	"howell/howell_api/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateCpsRebateDiscounts .
// @router /api/cps_rebate_discounts/create [POST]
func CreateCpsRebateDiscounts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req howell_api.CreateCpsRebateDiscountsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	data, err := service.CreateCpsRebateDiscounts(ctx, req.CrdEntity)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(howell_api.CreateCpsRebateDiscountsResponse)
	resp.Data = data

	c.JSON(consts.StatusOK, resp)
}

// GetCpsRebateDiscounts .
// @router /api/cps_rebate_discounts/get [GET]
func GetCpsRebateDiscounts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req howell_api.GetCpsRebateDiscountsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	data, err := service.GetCpsRebateDiscounts(ctx, req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(howell_api.GetCpsRebateDiscountsResponse)
	resp.Data = data

	c.JSON(consts.StatusOK, resp)
}

// QueryCpsRebateDiscounts .
// @router /api/cps_rebate_discounts/query [POST]
func QueryCpsRebateDiscounts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req howell_api.QueryCpsRebateDiscountsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	data, err := service.QueryCpsRebateDiscounts(ctx, req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(howell_api.QueryCpsRebateDiscountsResponse)
	resp.Data = data

	c.JSON(consts.StatusOK, resp)
}

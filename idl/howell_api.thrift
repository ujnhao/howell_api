namespace go coder.hao.howell_api 

enum CpsType {
    Unknow = 0
}

struct CpsRebateDiscounts {
    1:optional string id 
    2:optional string app_id
    3:optional string name
    4:optional CpsType cps_type   
    5:optional string jump_link 
    6:optional string extra  
    7:optional i32 status
}

struct CreateCpsRebateDiscountsRequest {
    1: required CpsRebateDiscounts crd_entity
}

struct CreateCpsRebateDiscountsData {
    1: optional string entity_id 
}

struct CreateCpsRebateDiscountsResponse {
    1: optional CreateCpsRebateDiscountsData data 

    100: i32 status
    101: string message
}

service HowellAPIService {
    CreateCpsRebateDiscountsResponse CreateCpsRebateDiscounts(1: CreateCpsRebateDiscountsRequest req) (api.post="/api/cps_rebate_discounts/create");
}

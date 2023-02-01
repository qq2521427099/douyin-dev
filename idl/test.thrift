namespace go test

struct BaseResp {
    1:i64 status_code
    2:string status_message
}

struct EchoReq{
    1:string message
}
struct EchoResp{
    1:BaseResp base_resp
    2:string message
}

service Echo {
    EchoResp Test(1:EchoReq echoReq)
}
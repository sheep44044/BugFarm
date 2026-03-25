namespace go hello_example

struct HelloReq {
    1: string Name (api.query="name")
}

struct HelloResp {
    1: string RespBody
}

service HelloService {
    HelloResp HelloMethod (1: HelloReq req)(api.get="/hello")
    HelloResp ByeMethod(1: HelloReq request) (api.get="/bye");
}
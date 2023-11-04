namespace go device

// -- basic --

struct BaseResp {
    1: i64 code,
    2: string msg,
}

// -- model --

// TODO: split address info in 3,4;
// TODO: better value,
struct Device {
    1: i64 id,
    2: string name,
    3: double longitude,
    4: double latitude,
    5: string desc,
    6: string value,
    7: string config,
    8: string version,
}

// -- Req & Rsp --

struct PingRequst {

}

struct PingResponse {
    1: BaseResp base,
}

struct CreateDeviceRequst {
    1: string device_name,
    2: optional double longitude,
    3: optional double latitude,
    4: optional string desc,
    5: optional string value,
    6: string config,
}

struct CreateDeviceResponse {
    1: BaseResp base,
}

struct RemoveDeviceRequst {
    1: i64 device_id,
}

struct RemoveDeviceResponse {
    1: BaseResp base,
}

struct UpdateDeviceRequst {
     1: required i64 device_id,
     2: optional string name,
     3: optional double longitude,
     4: optional double latitude,
     5: optional string desc,
     6: optional string value,
     7: optional string config,
     8: optional string version,
}

struct UpdateDeviceResponse {
    1: BaseResp base,
}

struct GetDeviceRequst {
    1: i64 device_id,
}

struct GetDeviceResponse {
    1: BaseResp base,
    2: Device device,
}


struct ListDeviceRequst {

}

struct ListDeviceResponse {
    1: BaseResp base,
    2: optional list<Device> device_list
}

// -- Service --

service DeviceService {
    PingResponse Ping(1: PingRequst req),
    CreateDeviceResponse Create(1: CreateDeviceRequst req),
    RemoveDeviceResponse Remove(1: RemoveDeviceRequst req),
    UpdateDeviceResponse Update(1: UpdateDeviceRequst req),
    GetDeviceResponse GetDeviceByID(1: GetDeviceRequst req),
    ListDeviceResponse ListDevice(1: ListDeviceRequst req),
}
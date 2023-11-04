namespace go api

// Model

struct User {
    1: i64 id,
    2: string name,
    3: string avatar,
    4: string email,
}

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

struct UserRegisterRequest {
    1: required string username,
    2: required string password,
}

struct UserRegisterResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required i64 user_id,
    4: required string token,
}

struct UserLoginRequest {
    1: required string username,
    2: required string password,
}

struct UserLoginResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required i64 user_id,
    4: required string token,
}

struct UserInfoRequest {
    1: required i64 user_id,
    2: required string token,
}

struct UserInfoResponse {
    1: required i64 status_code = 0,
    2: optional i64 status_msg,
    3: required User user,
}

struct DevicePingRequst {

}

struct DevicePingResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct CreateDeviceRequst {
    1: string name,
    2: optional double longitude,
    3: optional double latitude,
    4: optional string desc,
    5: optional string value,
    6: string config,
}

struct CreateDeviceResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct RemoveDeviceRequst {
    1: i64 device_id,
}

struct RemoveDeviceResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
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
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct GetDeviceRequst {
    1: i64 device_id,
}

struct GetDeviceResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: Device device,
}


struct ListDeviceRequst {

}

struct ListDeviceResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: optional list<Device> device_list
}


service BasicService {
    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post="/user/register/")
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/user/login/")
    UserInfoResponse UserInfo(1: UserInfoRequest req) (api.get="/user/info")

    DevicePingResponse Ping(1: DevicePingRequst req) (api.get="/device/ping")
    CreateDeviceResponse DeviceCreate(1: CreateDeviceRequst req) (api.post="/device/")
    RemoveDeviceResponse DeviceRemove(1: RemoveDeviceRequst req) (api.delete="/device/info")
    UpdateDeviceResponse DeviceUpdate(1: UpdateDeviceRequst req) (api.post="/device/info")
    GetDeviceResponse GetDeviceByID(1: GetDeviceRequst req) (api.get="/device/info")
    ListDeviceResponse ListDevice(1: ListDeviceRequst req) (api.get="/device/list")
}
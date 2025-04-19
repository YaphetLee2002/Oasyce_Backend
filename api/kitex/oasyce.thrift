include "user.thrift"
include "common.thrift"

service UserService {
    /** 获取当前用户信息*/
    user.GetUserResponse GetUser(1: user.GetUserRequest req) throws (1: common.Error err) (
        api.post = '/GetUser'),
}

local key = KEYS[1]
--使用次数，也就是验证次数
local cntKey = key .. ":cnt"
--预期中的验证码
local expectedCode = ARGV[1]

local cnt = tonumber(redis.call("get", cntKey))
local code = redis.call("get", key)

if cnt <= 0 then
    --说明一直输入错误,有人搞你
    return -1
end

if code == expectedCode then
    --输入正确
    --用完不能再用了
    redis.call("set", cntKey, -1)
    --redis.call("del", key)
    return 0
else
    -- 手抖输入错误
    -- 可验证次数减1
    redis.call("decr", cntKey)
    return -2
end
--你的验证码在redis上的key
local key = KEYS[1]

--使用次数
local cntKey = key .. ":cnt"
local val = ARGV[1]

--验证码有效时间10分钟，600s
local ttl = tonumber(redis.call("ttl", key))

if ttl == -1 then
    --key 存在，但是过期时间没有
    return -2
elseif ttl == -2 or ttl < 540 then
    redis.call("set", key, val)
    redis.call("expire", key, 600)
    redis.call("set", cntKey, 3)
    redis.call("expire", key, 600)
    return 0
else
    --发送太频繁，还没有过1分钟
    return -1
end

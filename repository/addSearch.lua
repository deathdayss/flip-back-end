local len = redis.call('zcard', KEYS[1])
if len == 10 then
    redis.call('zremrangebyrank', KEYS[1], 0, 0)
end
redis.call('zadd', KEYS[1], ARGV[1], ARGV[2])
if len == 0 then
    redis.call('expire', KEYS[1], 3600)
end
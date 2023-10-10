package stp

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type DistributedLocker interface {
	Lock() int
	Unlock() int
}

// distributed locker by redis
const (
	lockScript = `
		redis.replicate_commands()
		local lockerKey = KEYS[1]
		local value = ARGV[1]
		local expireSeconds = ARGV[2]

		local locked = redis.call('EXISTS', lockerKey)
		if locked == 1 then
			return 1
		end

		local setResultTable = redis.call('SET', lockerKey, value)
		if setResultTable['ok'] ~= 'OK' then
			return 2
		end

		local expireResult = redis.call('EXPIRE', lockerKey, expireSeconds)
		if expireResult ~= 1 then
			return 3
		end

		return 0
	`
	unlockScript = `
		redis.replicate_commands()
		local lockerKey = KEYS[1]

		local locked = redis.call('EXISTS', lockerKey)
		if locked == 0 then
			return 1
		end

		local delResult = redis.call('DEL', lockerKey)
		if delResult ~= 1 then
			return 2
		end

		return 0
	`
)

type DistributedRedisLocker struct{}

func NewDistributedRedisLocker() *DistributedRedisLocker { return &DistributedRedisLocker{} }

func (drl *DistributedRedisLocker) Lock(context context.Context, scripter redis.Scripter, key string, overtimeSeconds int, value interface{}) int {
	script := redis.NewScript(lockScript)
	result, err := script.Run(context, scripter, []string{key}, value, overtimeSeconds).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		return -1
	}
	return result
}

func (drl *DistributedRedisLocker) Unlock(context context.Context, scripter redis.Scripter, key string) int {
	script := redis.NewScript(unlockScript)
	result, err := script.Run(context, scripter, []string{key}).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		return -1
	}
	return result
}

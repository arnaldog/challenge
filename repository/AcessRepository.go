package repository

import (
	"awesomeProject/model"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func init() {
	fmt.Println("Access package initialized")
}

type AccessRepository interface {
	Register(access model.Access)
	Count(model.Access) int
	CountByIp(ip string) int
	SetCountByIp(ip string, count int) int
}

type AccessRepositoryImpl struct {
	Context     context.Context
	RedisClient *redis.Client
}

func (a AccessRepositoryImpl) Count(access model.Access) int {
	return a.CountByIp(access.Ip)
}

func (a AccessRepositoryImpl) Register(access model.Access) {
	var count = a.CountByIp(access.Ip)
	a.SetCountByIp(access.Ip, count+1)
}

func (a AccessRepositoryImpl) CountByIp(ip string) int {
	var count = 0
	val, err := a.RedisClient.Get(a.Context, ip).Result()
	if err == redis.Nil {
		val = "0"
	} else if err != nil {
		panic(err)
	}
	count, err = strconv.Atoi(val)

	if err != nil {
		panic(err)
	}

	return count
}

func (a AccessRepositoryImpl) SetCountByIp(ip string, count int) int {
	var err = a.RedisClient.Set(a.Context, ip, count, 0).Err()
	if err != nil {
		panic(err)
	}

	return count
}

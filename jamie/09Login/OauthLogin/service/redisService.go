package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"jamie/domain"
	"time"
)

var ctx = context.Background()

func RedisSessionCreate(cli *redis.Client, user domain.User) (string, error) {
	newUUID, err := uuid.NewUUID()
	marshal, _ := json.Marshal(&user) //newEmp json 형태로 변환

	_, err = cli.Set(ctx, newUUID.String(), marshal, time.Hour*1).Result()
	return newUUID.String(), err
}
func RedisSessionRead(cli *redis.Client, sessionID string) (domain.User, error) {
	var findUser domain.User
	val, err := cli.Get(ctx, sessionID).Result()
	if err != nil {
		return findUser, err
	}
	bytes := []byte(val)
	err = json.Unmarshal(bytes, &findUser)
	return findUser, err
}

func RedisSessionDelete(cli *redis.Client, sessionID string) error {
	_, err := cli.Del(ctx, sessionID).Result()
	return err
}

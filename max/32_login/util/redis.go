package util

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"loginMod"
	"time"
)

func RedisSave(red *redis.Client, user loginMod.User) (string, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	marshal, _ := json.Marshal(&user)

	_, err = red.Set(context.TODO(), newUUID.String(), marshal, time.Hour*1).Result()

	return newUUID.String(), err
}

func RedisGet(red *redis.Client, sessionId string) (loginMod.User, error) {

	var findUser loginMod.User
	value, err := red.Get(context.TODO(), sessionId).Result()

	if err != nil {
		return findUser, err
	}

	bytesData := []byte(value)
	err = json.Unmarshal(bytesData, &findUser)
	return findUser, err
}

func RedisDelete(red *redis.Client, sessionId string) error {

	_, err := red.Del(context.TODO(), sessionId).Result()
	return err

}

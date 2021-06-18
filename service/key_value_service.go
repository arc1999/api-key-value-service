package service

import (
	"api-key-value-service/model"
	"api-key-value-service/redis"
	"context"
)

type KeyValueService interface {
	Set(input model.KeyValue)  (*model.KeyValue,error)
	Get(key string)  (*model.KeyValue,error)
	GetAll()  ([]model.KeyValue,error)
}
type KeyValueServiceImpl struct {
}

var redisClient = redis.GetRedisClient()

func (i KeyValueServiceImpl) Set(input model.KeyValue)  (*model.KeyValue,error) {
	err := redisClient.Set(context.TODO(), input.Key, input.Value, 0).Err()
	if err != nil {
		return nil, err
	}
	return i.Get(input.Key)
}
func (i KeyValueServiceImpl) Get(key string)  (*model.KeyValue,error) {
	val,err := redisClient.Get(context.TODO(), key).Result()
	if err != nil {
		return nil, err
	}
	return &model.KeyValue{
		Key:   key,
		Value: val,
	},nil
}
func (i KeyValueServiceImpl) GetAll()  ([]model.KeyValue,error) {

	var cursor uint64
	var result []model.KeyValue
	var keys []string
	for {
		var err error
		keys, cursor, err = redisClient.Scan(context.TODO(), cursor, "*", 10).Result()
		if err != nil {
			return nil, err
		}
		if cursor == 0 {
			break
		}
	}
	for _,k:= range keys{
		d, err :=i.Get(k)
		if err != nil {
			return nil, err
		}
		result=append(result,*d)
	}
	return result,nil
}
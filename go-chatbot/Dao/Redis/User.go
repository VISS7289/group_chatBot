package Redis

import (
	"time"
)

func RedisSet(key string,value string) error {
	err := rdb.Set(key, value, 300*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisFind(key string) (string,error) {
	val, err := rdb.Get(key).Result()
	if err != nil {
		return "",err
	}
	return val,nil
}
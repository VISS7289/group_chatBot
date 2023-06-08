package Redis

import (
	"strconv"
	"time"
)

// 将指定键的值加1并存储到Redis数据库中
func RedisAdd(key string) (err error) {
	res:=rdb.Get(key).Val()
	val :=0
	if res==""{
		val =1
	}else{
		val ,err = strconv.Atoi(res)
		if err != nil {
			return err
		}
		val = val + 1
	}
	err = rdb.Set(key, val, time.Hour * 24).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisDelete(key string) error {
	return rdb.Del(key).Err()
}

func RedisVal(key string) string {
	return rdb.Get(key).Val()
}


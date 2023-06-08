package Redis

import (
	"time"
)

// 向Redis数据库中设置一个键值对
func RedisSet(key string,value string) error {
	err := rdb.Set(key, value, 300*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从Redis数据库中查找指定键的值
func RedisFind(key string) (string,error) {
	val, err := rdb.Get(key).Result()
	if err != nil {
		return "",err
	}
	return val,nil
}
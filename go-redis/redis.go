package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-iot-platform/lib/marshaller"
	"github.com/go-iot-platform/lib/structs"
	"github.com/vmihailenco/msgpack"
)

// NewRedis instance
func NewRedis(options interface{}) *Redis {
	opts := new(redis.Options)
	structs.Merge(opts, options)
	client := redis.NewClient(opts)
	return &Redis{
		db: client,
	}
}

// Redis struct
type Redis struct {
	db       *redis.Client
	pipeline *redis.Pipeline
}

// DB constructor
func (r *Redis) DB() *redis.Client {
	return r.db
}

// Close close redis
func (r *Redis) Close() error {
	return r.db.Close()
}

// Ping ping
func (r *Redis) Ping() *redis.StatusCmd {
	return r.db.Ping()
}

// Del del object by keys
func (r *Redis) Del(keys ...string) error {
	_, err := r.db.Del(keys...).Result()
	return err
}

// HDel del object by key and field
func (r *Redis) HDel(key string, fields ...string) error {
	_, err := r.db.HDel(key, fields...).Result()
	return err
}

// Get get object by key
func (r *Redis) Get(key string) (string, error) {
	return r.db.Get(key).Result()
}

// Set set object by key
func (r *Redis) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.db.Set(key, value, expiration).Result()
}

// SetObject set object by key and field
func (r *Redis) SetObject(objectKey string, field string, value interface{}) (bool, error) {
	bytes, err := msgpack.Marshal(value)
	if err != nil {
		return false, err
	}
	return r.db.HSet(objectKey, field, bytes).Result()
}

// GetObject get object by key and field
func (r *Redis) GetObject(objectKey string, field string, result interface{}) error {
	temp, err := r.db.HGet(objectKey, field).Bytes()
	if err != nil {
		return err
	}
	err = msgpack.Unmarshal(temp, result)
	return err
}

// CheckExistedObject will return true if the object is existed.
func (r *Redis) CheckExistedObject(objectKey string, field string) (bool, error) {
	existed, err := r.db.HExists(objectKey, field).Result()
	if err != nil {
		return false, err
	}
	if !existed {
		return false, nil
	}
	return true, nil
}

/*
MergeCache merge cache
- if exists , merge cache
- else save cache
*/
func (r *Redis) MergeCache(key string, field string, value interface{}) error {
	var cacheItem interface{}
	var mapValue map[string]interface{}
	conv := marshaller.ConventionalMarshaller{Value: value}

	b, err := conv.MarshalJSON()
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &mapValue)
	if err != nil {
		return err
	}
	r.GetObject(key, field, &cacheItem)
	if cacheItem == nil {
		_, err = r.SetObject(key, field, mapValue)
		return err
	}
	//merge cache
	var temp interface{}
	err = structs.MergeOverwriteCamel(cacheItem, mapValue, &temp)
	if err != nil {
		return err
	}
	_, err = r.SetObject(key, field, temp)
	return err
}

// GetData get data
func (r *Redis) GetData(key string, field string, result interface{}) error {
	var cacheItem interface{}
	err := r.GetObject(key, field, &cacheItem)
	if err != nil {
		return err
	}
	b, err := json.Marshal(cacheItem)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, result)
	return err
}

// SaveData save data
func (r *Redis) SaveData(key string, field string, value interface{}) error {
	var mapValue map[string]interface{}
	conv := marshaller.ConventionalMarshaller{Value: value}

	b, err := conv.MarshalJSON()
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &mapValue)
	if err != nil {
		return err
	}
	_, err = r.SetObject(key, field, mapValue)
	return err
}

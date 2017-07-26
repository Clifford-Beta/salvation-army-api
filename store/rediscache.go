package store

import (
	"github.com/go-redis/redis"
	"time"
	"salv_prj/model"
	l4g "github.com/alecthomas/log4go"
)

type Redis struct {
	Address string
	Password string
	DB int
	master *redis.Client
}
//"localhost:6379"
func InitCache(address, password string, db int) *Redis{
	rediscache := &Redis{
		DB:db,
		Password:password,
		Address:address,
	}
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       db,  // use default DB
	})
	rediscache.master = client
	_,err := client.Ping().Result()
	if err != nil {
		l4g.Critical("store.cache.ping","The cache is down, all queries will be directed to the db", err)
		//time.Sleep(time.Second)
	}else{
		l4g.Info("store.cache.ping","The cache is set and ready to accept connections")

	}

	return rediscache
}

func (r *Redis)GetMaster()*redis.Client{
	return r.master
}

func (r *Redis)Close()error{
	return r.master.Close()
}


func (r *Redis)Set( data map[string]interface{}, expire time.Duration)StoreChannel{
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		for k,v := range data {
			err := r.GetMaster().Set(k, v, expire).Err()
			if err != nil {
				result.Err = model.NewLocAppError("RedisCache.Get", "An error occured while trying to save object to redis", data, err.Error())
				break
			}
		}

		result.Data = "Done"
		storeChannel <- result
		close(storeChannel)
	}()

	return storeChannel
}

func (r *Redis)Get(keys []string) StoreChannel {
	storeChannel := make(StoreChannel, len(keys))
	go func() {
		result := StoreResult{}
		for _,v := range keys {
			val,err := r.GetMaster().Get(v).Result()
			if err != nil {
				result.Err = model.NewLocAppError("RedisCache.Get", "An error occured while trying to fetch object from redis",
					map[string]interface{}{"keys":keys}, err.Error())
			}
			result.Data = val
			storeChannel <- result
			}
		close(storeChannel)
	}()

	return storeChannel
}
package store

import (
	"testing"
	"os"
)

var(
	key="mine"
	val="This is it"
)
func TestInitCache(t *testing.T) {
	var Cache = InitCache("localhost:6379","",0)
	//if Cache.GetMaster().Ping() == "PONG"
	_, err := Cache.GetMaster().Ping().Result()
	if err != nil {
		t.Error("Redis Ping test failed with %v",err)
	}
}

func TestRedis_Set(t *testing.T) {
	returned := <- Cache.Set(map[string]interface{}{key:val},0)
	if returned.Err != nil {
		t.Error("Redis set test failed with",returned.Err)
	}
}

func TestRedis_Get(t *testing.T) {
	returned := <- Cache.Get([]string{key,})
	if returned.Err != nil {
		t.Error("Redis set test failed with",returned.Err)
	}
}

func TestConfigureApp(t *testing.T) {
	ConfigureApp(os.Getenv("GO_ENV"))

}
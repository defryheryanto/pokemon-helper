package redis

import (
	"testing"
	"time"
)

func TestGetRedisKey(t *testing.T) {
	redisKey := getRedisKey("Tyranitar")
	if redisKey != "pokemon:tyranitar" {
		t.Errorf("expected pokemon:tyranitar got %s", redisKey)
	}
}

func TestRedisExpiryTime(t *testing.T) {
	expiryTime := redisExpiryTime()
	if expiryTime != (5 * time.Minute) {
		t.Errorf("expected 5m0s got %s", expiryTime)
	}
}

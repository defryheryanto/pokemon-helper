package redis

import (
	"fmt"
	"strings"
	"time"
)

func getRedisKey(pokemonName string) string {
	return fmt.Sprintf("pokemon:%s", strings.ToLower(pokemonName))
}

func redisExpiryTime() time.Duration {
	return 5 * time.Minute
}

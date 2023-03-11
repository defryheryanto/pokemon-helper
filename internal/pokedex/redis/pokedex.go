package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/defry256/pokemon-helper/internal/pokedex"
	"github.com/defry256/pokemon-helper/internal/pokemon"
	"github.com/defry256/pokemon-helper/internal/pokemontype"
	queue "github.com/defryheryanto/job-queuer"
	"github.com/go-redis/redis/v8"
)

// RedisDecorator is a decorator that implements pokedex.IService
//
// Decorate the base service to using redis if possible
type RedisDecorator struct {
	pokedex.IService
	redisClient *redis.Client
	queuer      *queue.Queuer
}

func NewRedisDecorator(
	baseService pokedex.IService,
	redisClient *redis.Client,
	queuer *queue.Queuer,
) *RedisDecorator {
	return &RedisDecorator{baseService, redisClient, queuer}
}

func (s *RedisDecorator) GetPokedex(ctx context.Context, pokemonName string) *pokemon.PokemonData {
	pokemonData, err := s.getPokemonFromRedis(ctx, pokemonName)
	if err == nil && pokemonData != nil {
		fmt.Println("got from redis")
		return pokemonData
	}

	pokemonData = s.IService.GetPokedex(ctx, pokemonName)
	s.queuer.Push(newPokemonRedisRegistrar(
		[]*pokemon.PokemonData{pokemonData},
		s.redisClient,
	))
	return pokemonData
}

func (s *RedisDecorator) getPokemonFromRedis(ctx context.Context, pokemonName string) (*pokemon.PokemonData, error) {
	pokedexByte, err := s.redisClient.Get(ctx, getRedisKey(pokemonName)).Bytes()
	if err != nil {
		return nil, err
	}
	fmt.Printf("get with %s\n", getRedisKey(pokemonName))

	var pokemonMap map[string]interface{}
	err = json.Unmarshal(pokedexByte, &pokemonMap)
	if err != nil {
		return nil, err
	}

	var baseStatus *pokemon.Status
	switch stat := pokemonMap["base_status"].(type) {
	case *pokemon.Status:
		baseStatus = stat
	case map[string]interface{}:
		hp := int(stat["hp"].(float64))
		attack := int(stat["attack"].(float64))
		defense := int(stat["defense"].(float64))
		specialAttack := int(stat["special_attack"].(float64))
		spesialDefense := int(stat["special_defense"].(float64))
		speed := int(stat["speed"].(float64))
		total := int(stat["total"].(float64))
		baseStatus = &pokemon.Status{
			HP:             hp,
			Attack:         attack,
			Defense:        defense,
			SpecialAttack:  specialAttack,
			SpecialDefense: spesialDefense,
			Speed:          speed,
			Total:          total,
		}
	}

	pokemonTypes := []pokemontype.IType{}
	switch pokeTypes := pokemonMap["types"].(type) {
	case []pokemontype.IType:
		pokemonTypes = pokeTypes
	case []interface{}:
		for _, t := range pokeTypes {
			pokemonTypes = append(pokemonTypes, pokemontype.Type(fmt.Sprintf("%s", t)))
		}
	}

	return &pokemon.PokemonData{
		Name:       fmt.Sprintf("%s", pokemonMap["name"]),
		BaseStatus: baseStatus,
		Types:      pokemonTypes,
	}, nil
}

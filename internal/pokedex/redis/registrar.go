package redis

import (
	"context"
	"encoding/json"

	"github.com/defry256/pokemon-helper/internal/pokemon"
	"github.com/go-redis/redis/v8"
)

type pokemonRedisRegistrar struct {
	pokemons    []*pokemon.PokemonData
	redisClient *redis.Client
}

func newPokemonRedisRegistrar(
	pokemons []*pokemon.PokemonData,
	client *redis.Client,
) *pokemonRedisRegistrar {
	return &pokemonRedisRegistrar{pokemons, client}
}

func (r *pokemonRedisRegistrar) GetTitle() string {
	return "redis_register_pokemon"
}

func (r *pokemonRedisRegistrar) Do(ctx context.Context) error {
	for _, poke := range r.pokemons {
		err := r.setPokemon(ctx, poke)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *pokemonRedisRegistrar) setPokemon(ctx context.Context, pokemonData *pokemon.PokemonData) error {
	if pokemonData == nil {
		return nil
	}
	b, err := json.Marshal(pokemonData)
	if err != nil {
		return err
	}

	r.redisClient.Set(ctx, getRedisKey(pokemonData.Name), b, redisExpiryTime())
	return nil
}

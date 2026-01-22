package app

import (
	"database/sql"

	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	pokedex_service "github.com/defryheryanto/pokemon-helper/internal/pokedex/pokeapi"
	pokedex_redis "github.com/defryheryanto/pokemon-helper/internal/pokedex/redis"
	pokedex_traced "github.com/defryheryanto/pokemon-helper/internal/pokedex/traced"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	pokemon_repository "github.com/defryheryanto/pokemon-helper/internal/pokemon/repository"
	pokemon_service "github.com/defryheryanto/pokemon-helper/internal/pokemon/service"
	pokemon_traced "github.com/defryheryanto/pokemon-helper/internal/pokemon/traced"
	"github.com/defryheryanto/pokemon-helper/internal/teambuilder"
	teambuilder_traced "github.com/defryheryanto/pokemon-helper/internal/teambuilder/traced"
	teambuilder_service "github.com/defryheryanto/pokemon-helper/internal/teambuilder/v1"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/trace"
)

func BuildApp(redisClient *redis.Client, tracer trace.Tracer, db *sql.DB) *App {
	pokedexService := setupPokedex(redisClient, tracer)
	teamBuilderService := setupTeamBuilder(pokedexService, tracer)
	pokemonService := setupPokemon(db, tracer)

	return &App{
		Pokedex:     pokedexService,
		TeamBuilder: teamBuilderService,
		Pokemon:     pokemonService,
	}
}

func setupPokedex(redisClient *redis.Client, tracer trace.Tracer) pokedex.IService {
	var pokedexService pokedex.IService
	pokedexService = pokedex_service.NewPokeAPI()
	pokedexService = pokedex_redis.NewRedisDecorator(pokedexService, redisClient)
	pokedexService = pokedex_traced.NewTracedService(pokedexService, tracer)

	return pokedexService
}

func setupTeamBuilder(pokedexService pokedex.IService, tracer trace.Tracer) teambuilder.IService {
	var teamBuilderService teambuilder.IService
	teamBuilderService = teambuilder_service.NewService(pokedexService)
	teamBuilderService = teambuilder_traced.NewTracedService(teamBuilderService, tracer)

	return teamBuilderService
}

func setupPokemon(db *sql.DB, tracer trace.Tracer) pokemon.Service {
	var pokemonService pokemon.Service
	repository := pokemon_repository.NewRepository(db)
	pokemonService = pokemon_service.NewService(repository)
	pokemonService = pokemon_traced.NewTracedService(pokemonService, tracer)

	return pokemonService
}

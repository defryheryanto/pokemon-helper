package main

import (
	"github.com/defryheryanto/pokemon-helper/internal/app"
	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	pokedex_service "github.com/defryheryanto/pokemon-helper/internal/pokedex/pokeapi"
	pokedex_redis "github.com/defryheryanto/pokemon-helper/internal/pokedex/redis"
	"github.com/defryheryanto/pokemon-helper/internal/pokedex/traced"
	"github.com/defryheryanto/pokemon-helper/internal/teambuilder"
	teambuilder_traced "github.com/defryheryanto/pokemon-helper/internal/teambuilder/traced"
	teambuilder_service "github.com/defryheryanto/pokemon-helper/internal/teambuilder/v1"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/trace"
)

func BuildApp(redisClient *redis.Client, tracer trace.Tracer) *app.App {
	pokedexService := setupPokedex(redisClient, tracer)
	teamBuilderService := setupTeamBuilder(pokedexService, tracer)

	return &app.App{
		Pokedex:     pokedexService,
		TeamBuilder: teamBuilderService,
	}
}

func setupPokedex(redisClient *redis.Client, tracer trace.Tracer) pokedex.IService {
	var pokedexService pokedex.IService
	pokedexService = pokedex_service.NewPokeAPI()
	pokedexService = pokedex_redis.NewRedisDecorator(pokedexService, redisClient)
	pokedexService = traced.NewTracedService(pokedexService, tracer)

	return pokedexService
}

func setupTeamBuilder(pokedexService pokedex.IService, tracer trace.Tracer) teambuilder.IService {
	var teamBuilderService teambuilder.IService
	teamBuilderService = teambuilder_service.NewService(pokedexService)
	teamBuilderService = teambuilder_traced.NewTracedService(teamBuilderService, tracer)

	return teamBuilderService
}

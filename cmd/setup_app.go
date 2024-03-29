package main

import (
	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/pokedex"
	pokedex_redis "github.com/defry256/pokemon-helper/internal/pokedex/redis"
	"github.com/defry256/pokemon-helper/internal/pokedex/traced"
	pokedex_service "github.com/defry256/pokemon-helper/internal/pokedex/v1"
	"github.com/defry256/pokemon-helper/internal/teambuilder"
	teambuilder_traced "github.com/defry256/pokemon-helper/internal/teambuilder/traced"
	teambuilder_service "github.com/defry256/pokemon-helper/internal/teambuilder/v1"
	queue "github.com/defryheryanto/job-queuer"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel/trace"
)

func BuildApp(redisClient *redis.Client, queuer *queue.Queuer, tracer trace.Tracer) *app.App {
	pokedexService := setupPokedex(redisClient, queuer, tracer)
	teamBuilderService := setupTeamBuilder(pokedexService, tracer)

	return &app.App{
		Pokedex:     pokedexService,
		TeamBuilder: teamBuilderService,
	}
}

func setupPokedex(redisClient *redis.Client, queuer *queue.Queuer, tracer trace.Tracer) pokedex.IService {
	var pokedexService pokedex.IService
	pokedexService = pokedex_service.NewService()
	pokedexService = pokedex_redis.NewRedisDecorator(pokedexService, redisClient, queuer)
	pokedexService = traced.NewTracedService(pokedexService, tracer)

	return pokedexService
}

func setupTeamBuilder(pokedexService pokedex.IService, tracer trace.Tracer) teambuilder.IService {
	var teamBuilderService teambuilder.IService
	teamBuilderService = teambuilder_service.NewService(pokedexService)
	teamBuilderService = teambuilder_traced.NewTracedService(teamBuilderService, tracer)

	return teamBuilderService
}

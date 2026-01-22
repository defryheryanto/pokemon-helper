package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/defryheryanto/pokemon-helper/config"
	"github.com/defryheryanto/pokemon-helper/internal/logger"
	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	pokedex_service "github.com/defryheryanto/pokemon-helper/internal/pokedex/pokeapi"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	pokemon_repository "github.com/defryheryanto/pokemon-helper/internal/pokemon/repository"
	pokemon_service "github.com/defryheryanto/pokemon-helper/internal/pokemon/service"
	"github.com/go-co-op/gocron/v2"
	_ "github.com/lib/pq"
)

const (
	pokedexPageSize = 100
)

func main() {
	config.Load()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	db := setupDB()
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("error closing database connection", err)
		}
	}()

	pokedexService := pokedex_service.NewPokeAPI()
	pokemonService := setupPokemon(db)

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	scheduler, err := gocron.NewScheduler(gocron.WithLocation(loc))
	if err != nil {
		panic(err)
	}

	_, err = scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(3, 0, 0))),
		gocron.NewTask(func() {
			logger.Print("cron: starting pokemon population")
			if err := populatePokemonData(context.Background(), pokedexService, pokemonService); err != nil {
				logger.Error("cron: populate pokemon failed", err)
				return
			}
			logger.Print("cron: pokemon population finished")
		}),
	)
	if err != nil {
		panic(err)
	}

	scheduler.Start()
	logger.Print("cron: scheduler started")

	<-quit

	logger.Print("cron: shutting down scheduler")
	if err := scheduler.Shutdown(); err != nil {
		logger.Error("cron: scheduler shutdown failed", err)
	}
}

func populatePokemonData(ctx context.Context, pokedexService pokedex.IService, pokemonService pokemon.Service) error {
	for page := 1; ; page++ {
		list, err := pokedexService.GetAllPokedex(ctx, pokedex.GetAllPokedexFilter{
			Page:     page,
			PageSize: pokedexPageSize,
		})
		if err != nil {
			return err
		}
		if len(list) == 0 {
			return nil
		}

		for _, item := range list {
			detail, err := pokedexService.GetPokedex(ctx, strconv.Itoa(item.ID))
			if err != nil {
				logger.Error(fmt.Sprintf("cron: get pokedex detail failed for %s", item.Name), err)
				continue
			}
			if detail == nil {
				logger.Print(fmt.Sprintf("cron: empty pokedex detail for %s", item.Name))
				continue
			}

			if err := pokemonService.Upsert(ctx, detail); err != nil {
				logger.Error(fmt.Sprintf("cron: upsert pokemon failed for %s", detail.Name), err)
			}
		}
	}
}

func setupDB() *sql.DB {
	dsn := config.DatabaseConnectionString()
	if dsn == "" {
		panic("ENV DATABASE_CONNECTION_STRING is empty")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("error opening database connection", err)
		panic(err)
	}

	return db
}

func setupPokemon(db *sql.DB) pokemon.Service {
	repository := pokemon_repository.NewRepository(db)
	return pokemon_service.NewService(repository)
}

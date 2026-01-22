package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/defryheryanto/pokemon-helper/config"
	"github.com/defryheryanto/pokemon-helper/internal/logger"
	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	pokedex_service "github.com/defryheryanto/pokemon-helper/internal/pokedex/pokeapi"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	pokemon_repository "github.com/defryheryanto/pokemon-helper/internal/pokemon/repository"
	pokemon_service "github.com/defryheryanto/pokemon-helper/internal/pokemon/service"
	_ "github.com/lib/pq"
)

func main() {
	config.Load()

	db := setupDB()
	pokemonService := setupPokemon(db)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("error closing database connection", err)
		}
	}()

	pokedexService := pokedex_service.NewPokeAPI()
	if err := populatePokemonData(context.Background(), pokedexService, pokemonService); err != nil {
		logger.Error("populate pokemon data failed", err)
	}
}

func populatePokemonData(ctx context.Context, pokedexService pokedex.IService, pokemonService pokemon.Service) error {
	for page := 1; ; page++ {
		list, err := pokedexService.GetAllPokedex(ctx, pokedex.GetAllPokedexFilter{
			Page:     page,
			PageSize: 100,
		})
		if err != nil {
			return err
		}
		if len(list) == 0 {
			return nil
		}
		logger.Print(fmt.Sprintf("retrieved %d pokemon data", len(list)))

		for _, item := range list {
			logger.Print(fmt.Sprintf("processing pokemon %s", item.Name))
			detail, err := pokedexService.GetPokedex(ctx, strconv.Itoa(item.ID))
			if err != nil {
				logger.Error(fmt.Sprintf("cron: get pokedex detail failed for %s", item.Name), err)
				continue
			}
			if detail == nil {
				logger.Print(fmt.Sprintf("cron: empty pokedex detail for %s", item.Name))
				continue
			}
			logger.Print(fmt.Sprintf("pokemon %s detail retrieved", item.Name))

			if err := pokemonService.Upsert(ctx, detail); err != nil {
				logger.Error(fmt.Sprintf("cron: upsert pokemon failed for %s", detail.Name), err)
			}
			logger.Print(fmt.Sprintf("processing pokemon %s complete", item.Name))
		}
	}
}

func setupDB() *sql.DB {
	dsn := config.DatabaseConnectionString()
	log.Println(dsn)
	log.Println(config.LoggerFilepath())
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

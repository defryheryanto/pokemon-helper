package main

import (
	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/pokedex/v1"
)

func BuildApp() *app.App {
	pokedex := pokedex.NewService()

	return &app.App{
		Pokedex: pokedex,
	}
}

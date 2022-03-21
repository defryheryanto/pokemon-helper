package main

import (
	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/pokedex/v1"
	"github.com/defry256/pokemon-helper/internal/teambuilder/v1"
)

func BuildApp() *app.App {
	pokedex := pokedex.NewService()
	teambuilder := teambuilder.NewService(pokedex)

	return &app.App{
		Pokedex:     pokedex,
		TeamBuilder: teambuilder,
	}
}

package app

import (
	"github.com/defry256/pokemon-helper/internal/pokedex"
	"github.com/defry256/pokemon-helper/internal/teambuilder"
)

type App struct {
	Pokedex     pokedex.IService
	TeamBuilder teambuilder.IService
}

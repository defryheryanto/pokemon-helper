package app

import (
	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	"github.com/defryheryanto/pokemon-helper/internal/teambuilder"
)

type App struct {
	Pokedex     pokedex.IService
	TeamBuilder teambuilder.IService
	Pokemon     pokemon.Service
}

package app

import "github.com/defry256/pokemon-helper/internal/pokedex"

type App struct {
	Pokedex pokedex.IService
}

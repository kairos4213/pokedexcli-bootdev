package main

import (
	"github.com/kairos4213/pokedexcli-bootdev/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocAreasURL *string
	prevLocAreasURL *string
}

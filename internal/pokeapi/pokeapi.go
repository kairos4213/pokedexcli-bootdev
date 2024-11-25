package pokeapi

import (
	"net/http"
	"time"

	"github.com/kairos4213/pokedexcli-bootdev/internal/pokecache"
)

const apiURL = "https://pokeapi.co/api/v2/"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

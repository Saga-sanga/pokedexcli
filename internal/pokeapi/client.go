package pokeapi

import (
	"net/http"
	"time"

	"github.com/saga-sanga/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	pokecache  *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokecache: pokecache.NewCache(cacheInterval),
	}
}

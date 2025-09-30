package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zuluca/pokedexcli/internal/pokecache"
)

type Client struct {
	cache *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) FetchData(url string) ([]byte, error) {
	// 1. Check cache
	if data, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit:", url)
		return data, nil
	}

	// 2. If not in cache → fetch from API
	fmt.Println("Cache miss, fetching:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 3. Save in cache
	c.cache.Add(url, body)

	return body, nil
}

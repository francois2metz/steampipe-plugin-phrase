package phrase

import (
	"context"
	"errors"
	"os"

	phrase "github.com/phrase/phrase-go/v4"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const defaultDatacenter = "https://api.phrase.com/v2"

type phraseCache struct {
	client *phrase.APIClient
	apiKey phrase.APIKey
}

func toAuthContext(ctx context.Context, apiKey phrase.APIKey) context.Context {
	return context.WithValue(ctx, phrase.ContextAPIKey, apiKey)
}

func connect(ctx context.Context, d *plugin.QueryData) (*phrase.APIClient, *context.Context, error) {
	cacheKey := "phrase"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		cache := cachedData.(phraseCache)
		authContext := toAuthContext(ctx, cache.apiKey)
		return cache.client, &authContext, nil
	}

	access_token := os.Getenv("PHRASE_ACCESS_TOKEN")
	datacenter := os.Getenv("PHRASE_DATACENTER")

	phraseConfig := GetConfig(d.Connection)
	if phraseConfig.AccessToken != nil {
		access_token = *phraseConfig.AccessToken
	}
	if phraseConfig.Datacenter != nil {
		datacenter = *phraseConfig.Datacenter
	}

	if access_token == "" {
		return nil, nil, errors.New("'access_token' must be set in the connection configuration. Edit your connection configuration file or set the PHRASE_ACCESS_TOKEN environment variable and then restart Steampipe")
	}

	if datacenter == "" {
		datacenter = defaultDatacenter
	}

	cfg := phrase.NewConfiguration()
	client := phrase.NewAPIClient(cfg)
	client.ChangeBasePath(datacenter)

	apiKey := phrase.APIKey{
		Key:    access_token,
		Prefix: "token",
	}

	cache := phraseCache{
		client: client,
		apiKey: apiKey,
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, cache)

	authContext := toAuthContext(ctx, apiKey)

	return client, &authContext, nil
}

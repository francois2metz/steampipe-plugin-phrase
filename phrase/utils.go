package phrase

import (
	"context"
	"errors"
	"os"

	phrase "github.com/phrase/phrase-go/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*phrase.APIClient, *context.Context, error) {
	cfg := phrase.NewConfiguration()
	client := phrase.NewAPIClient(cfg)

	cacheKey := "phrase"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		authContext := context.WithValue(ctx, phrase.ContextAPIKey, cachedData.(phrase.APIKey))
		return client, &authContext, nil
	}

	access_token := os.Getenv("PHRASE_ACCESS_TOKEN")

	phraseConfig := GetConfig(d.Connection)
	if phraseConfig.AccessToken != nil {
		access_token = *phraseConfig.AccessToken
	}

	if access_token == "" {
		return nil, nil, errors.New("'access_token' must be set in the connection configuration. Edit your connection configuration file or set the PHRASE_ACCESS_TOKEN environment variable and then restart Steampipe")
	}

	apiKey := phrase.APIKey{
		Key:    access_token,
		Prefix: "token",
	}
	authContext := context.WithValue(ctx, phrase.ContextAPIKey, apiKey)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, apiKey)

	return client, &authContext, nil
}

package phrase

import (
	"context"
	"errors"
	"os"

	phrase "github.com/phrase/phrase-go/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*phrase.APIClient, *context.Context, error) {
	cacheKey := "phrase"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*phrase.APIClient), nil, nil
	}

	access_token := os.Getenv("PHRASE_ACCESS_TOKEN")

	phraseConfig := GetConfig(d.Connection)
	if phraseConfig.AccessToken != nil {
		access_token = *phraseConfig.AccessToken
	}

	if access_token == "" {
		return nil, nil, errors.New("'access_token' must be set in the connection configuration. Edit your connection configuration file or set the PHRASE_ACCESS_TOKEN environment variable and then restart Steampipe")
	}

	auth := context.WithValue(ctx, phrase.ContextAPIKey, phrase.APIKey{
		Key:    access_token,
		Prefix: "token",
	})

	cfg := phrase.NewConfiguration()
	client := phrase.NewAPIClient(cfg)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, &auth, nil
}

package phrase

import (
	"context"

	"github.com/antihax/optional"
	phrase "github.com/phrase/phrase-go/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePhraseLocale() *plugin.Table {
	return &plugin.Table{
		Name:        "phrase_locale",
		Description: "A locale defines a user language and contains language and country related parameters.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("project_id"),
			Hydrate:    listLocale,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project_id", "id"}),
			Hydrate:    getLocale,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
				Description: "Unique id of the locale.",
			},
			{
				Name:        "project_id",
				Type:        proto.ColumnType_STRING,
				Description: "Project id associated to the locale.",
				Transform:   transform.FromQual("project_id"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the locale.",
			},
			{
				Name:        "code",
				Type:        proto.ColumnType_STRING,
				Description: "Language code of the locale.",
			},
			{
				Name:        "default",
				Type:        proto.ColumnType_BOOL,
				Description: "Is this the default locale of the project.",
			},
			{
				Name:        "main",
				Type:        proto.ColumnType_BOOL,
				Description: "Is this the main locale of the project",
			},
			{
				Name:        "rtl",
				Type:        proto.ColumnType_BOOL,
				Description: "Right to left locale.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Creation date of the locale.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Update date of the locale.",
			},
			{
				Name:        "statictics_keys_total_count",
				Type:        proto.ColumnType_INT,
				Hydrate:     getLocaleDetail,
				Transform:   transform.FromField("Statistics.KeysTotalCount"),
				Description: "Number of keys.",
			},
			{
				Name:        "statictics_keys_untranslated_count",
				Type:        proto.ColumnType_INT,
				Hydrate:     getLocaleDetail,
				Transform:   transform.FromField("Statistics.KeysUntranslatedCount"),
				Description: "Number of untranslated keys.",
			},
			{
				Name:        "statictics_words_total_count",
				Type:        proto.ColumnType_INT,
				Hydrate:     getLocaleDetail,
				Transform:   transform.FromField("Statistics.WordsTotalCount"),
				Description: "Total number of words.",
			},
			{
				Name:        "statictics_translations_completed_count ",
				Type:        proto.ColumnType_INT,
				Hydrate:     getLocaleDetail,
				Transform:   transform.FromField("Statistics.TranslationsCompletedCount"),
				Description: "Number of completed translations.",
			},
			{
				Name:        "statictics_translations_unverified_count",
				Type:        proto.ColumnType_INT,
				Hydrate:     getLocaleDetail,
				Transform:   transform.FromField("Statistics.TranslationsUnverifiedCount"),
				Description: "Number of unverified translations.",
			},
			{
				Name:        "statictics_unverified_words_count",
				Type:        proto.ColumnType_INT,
				Hydrate:     getLocaleDetail,
				Transform:   transform.FromField("Statistics.UnverifiedWordsCount"),
				Description: "Number of unverified words.",
			},
			{
				Name:        "statictics_missing_words_count",
				Type:        proto.ColumnType_INT,
				Hydrate:     getLocaleDetail,
				Transform:   transform.FromField("Statistics.MissingWordsCount"),
				Description: "Number of missing words.",
			},
		},
	}
}

func listLocale(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_locale.listLocale", "connection_error", err)
		return nil, err
	}
	opts := &phrase.LocalesListOpts{
		Page:    optional.NewInt32(0),
		PerPage: optional.NewInt32(100),
	}
	project_id := d.EqualsQuals["project_id"].GetStringValue()
	for {
		locales, response, err := client.LocalesApi.LocalesList(*authContext, project_id, opts)
		if err != nil {
			plugin.Logger(*authContext).Error("phrase_locale.listLocale", err)
			return nil, err
		}
		for _, locale := range locales {
			d.StreamListItem(*authContext, locale)
		}
		if response.NextPage == 0 {
			break
		}
		opts.Page = optional.NewInt32(int32(response.NextPage))
		if d.RowsRemaining(ctx) <= 0 {
			break
		}
	}
	return nil, nil
}

func getLocale(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQuals["id"].GetStringValue()
	return getLocaleById(ctx, d, id)
}

func getLocaleDetail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	switch h.Item.(type) {
	case phrase.Locale:
		locale := h.Item.(phrase.Locale)
		id := locale.Id
		return getLocaleById(ctx, d, id)
	case phrase.LocaleDetails:
		return h.Item.(phrase.LocaleDetails), nil
	default:
		return nil, nil
	}
}

func getLocaleById(ctx context.Context, d *plugin.QueryData, id string) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_locale.getLocaleById", "connection_error", err)
		return nil, err
	}
	project_id := d.EqualsQuals["project_id"].GetStringValue()
	result, _, err := client.LocalesApi.LocaleShow(*authContext, project_id, id, nil)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_locale.getLocaleById", err)
		return nil, err
	}
	return result, nil
}

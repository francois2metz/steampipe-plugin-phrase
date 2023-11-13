package phrase

import (
	"context"

	"github.com/antihax/optional"
	phrase "github.com/phrase/phrase-go/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePhraseKey() *plugin.Table {
	return &plugin.Table{
		Name:        "phrase_key",
		Description: "Keys are used to identify translatable text strings within software code.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name: "project_id",
				},
				{
					Name:    "branch",
					Require: plugin.Optional,
				},
				{
					Name:    "q",
					Require: plugin.Optional,
				},
			},
			Hydrate: listKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project_id", "name"}),
			Hydrate:    getKey,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
				Description: "Unique id of the key.",
			},
			{
				Name:        "project_id",
				Type:        proto.ColumnType_STRING,
				Description: "Project id associated to the key.",
				Transform:   transform.FromQual("project_id"),
			},
			{
				Name:        "branch",
				Type:        proto.ColumnType_STRING,
				Description: "The branch to filter tag",
				Transform:   transform.FromQual("branch"),
			},
			{
				Name:        "q",
				Type:        proto.ColumnType_STRING,
				Description: "The query to do broad search. See the parameter key https://developers.phrase.com/api/#get-/projects/-project_id-/keys.",
				Transform:   transform.FromQual("q"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the key.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description of the key.",
			},
			{
				Name:        "name_hash",
				Type:        proto.ColumnType_STRING,
				Description: "Hash of the name.",
			},
			{
				Name:        "plural",
				Type:        proto.ColumnType_BOOL,
				Description: "Is the key a plural.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: "The associated tags to the key.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Creation date of the key.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Update date of the key.",
			},
		},
	}
}

func listKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_key.listKey", "connection_error", err)
		return nil, err
	}
	opts := &phrase.KeysListOpts{
		Page:    optional.NewInt32(0),
		PerPage: optional.NewInt32(100),
	}
	project_id := d.EqualsQuals["project_id"].GetStringValue()
	branch := d.EqualsQuals["branch"].GetStringValue()
	if branch != "" {
		opts.Branch = optional.NewString(branch)
	}
	q := d.EqualsQuals["q"].GetStringValue()
	if q != "" {
		opts.Q = optional.NewString(q)
	}
	for {
		locales, response, err := client.KeysApi.KeysList(*authContext, project_id, opts)
		if err != nil {
			plugin.Logger(*authContext).Error("phrase_key.listKey", err)
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

func getKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_key.getKey", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	project_id := d.EqualsQuals["project_id"].GetStringValue()
	result, _, err := client.KeysApi.KeyShow(*authContext, project_id, id, nil)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_key.getKey", err)
		return nil, err
	}
	return result, nil
}

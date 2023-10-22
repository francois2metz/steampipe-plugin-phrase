package phrase

import (
	"context"

	"github.com/antihax/optional"
	phrase "github.com/phrase/phrase-go/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePhraseTag() *plugin.Table {
	return &plugin.Table{
		Name:        "phrase_tag",
		Description: "",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("project_id"),
			Hydrate:    listTag,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project_id", "name"}),
			Hydrate:    getTag,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the tag.",
			},
			{
				Name:        "project_id",
				Type:        proto.ColumnType_STRING,
				Description: "Project id associated to the tag.",
				Transform:   transform.FromQual("project_id"),
			},
			{
				Name:        "keys_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of associated keys.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Creation date of the tag.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Update date of the tag.",
			},
		},
	}
}

func listTag(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_tag.listTag", "connection_error", err)
		return nil, err
	}
	opts := &phrase.TagsListOpts{
		Page:    optional.NewInt32(0),
		PerPage: optional.NewInt32(100),
	}
	project_id := d.EqualsQuals["project_id"].GetStringValue()
	for {
		locales, response, err := client.TagsApi.TagsList(*authContext, project_id, opts)
		if err != nil {
			plugin.Logger(*authContext).Error("phrase_tag.listTag", err)
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

func getTag(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_tag.getTag", "connection_error", err)
		return nil, err
	}
	name := d.EqualsQuals["name"].GetStringValue()
	project_id := d.EqualsQuals["project_id"].GetStringValue()
	result, _, err := client.TagsApi.TagShow(*authContext, project_id, name, nil)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_tag.getTag", err)
		return nil, err
	}
	return result, nil
}

package phrase

import (
	"context"

	"github.com/antihax/optional"
	phrase "github.com/phrase/phrase-go/v4"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePhraseProject() *plugin.Table {
	return &plugin.Table{
		Name:        "phrase_project",
		Description: "Projects the current user has access to.",
		List: &plugin.ListConfig{
			Hydrate:    listProject,
			KeyColumns: plugin.OptionalColumns([]string{"account_id"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProject,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
				Description: "Unique id of the project.",
			},
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Description: "Account id associated to the project.",
				Transform:   transform.FromQual("account_id"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the project.",
			},
			{
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "Slug of the project.",
			},
			{
				Name:        "main_format",
				Type:        proto.ColumnType_STRING,
				Description: "Main format of the project.",
			},
			{
				Name:        "project_image_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ProjectImageUrl"),
				Description: "Image URL of the project.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Creation date of the project.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Update date of the project.",
			},
		},
	}
}

func listProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_project.listProject", "connection_error", err)
		return nil, err
	}
	opts := &phrase.ProjectsListOpts{
		Page:    optional.NewInt32(0),
		PerPage: optional.NewInt32(100),
	}
	if d.EqualsQuals["account_id"] != nil {
		opts.AccountId = optional.NewString(d.EqualsQuals["account_id"].GetStringValue())
	}
	for {
		projects, response, err := client.ProjectsApi.ProjectsList(*authContext, opts)
		if err != nil {
			plugin.Logger(*authContext).Error("phrase_project.listProject", err)
			return nil, err
		}
		for _, project := range projects {
			d.StreamListItem(*authContext, project)
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

func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_project.getProject", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	result, _, err := client.ProjectsApi.ProjectShow(*authContext, id, nil)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_project.getProject", err)
		return nil, err
	}
	return result, nil
}

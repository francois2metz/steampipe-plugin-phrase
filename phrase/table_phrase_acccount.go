package phrase

import (
	"context"

	"github.com/antihax/optional"
	phrase "github.com/phrase/phrase-go/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePhraseAccount() *plugin.Table {
	return &plugin.Table{
		Name:        "phrase_account",
		Description: "Accounts the current user has access to.",
		List: &plugin.ListConfig{
			Hydrate: listAccount,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAccount,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
				Description: "Unique id of the account.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the account.",
			},
			{
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "Slug of the account.",
			},
			{
				Name:        "company",
				Type:        proto.ColumnType_STRING,
				Description: "Company of the account.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Creation date of the account.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Update date of the account.",
			},
			{
				Name:        "company_logo_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CompanyLogoUrl"),
				Description: "URL to the company logo.",
			},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_account.listAccount", "connection_error", err)
		return nil, err
	}
	opts := &phrase.AccountsListOpts{
		Page:    optional.NewInt32(0),
		PerPage: optional.NewInt32(100),
	}
	for {
		accounts, response, err := client.AccountsApi.AccountsList(*authContext, opts)
		if err != nil {
			plugin.Logger(*authContext).Error("phrase_account.listAccount", err)
			return nil, err
		}
		for _, account := range accounts {
			d.StreamListItem(*authContext, account)
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

func getAccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_account.getAccount", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	result, _, err := client.AccountsApi.AccountShow(*authContext, id, nil)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_account.getAccount", err)
		return nil, err
	}
	return result, nil
}

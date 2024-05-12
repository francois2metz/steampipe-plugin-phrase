package phrase

import (
	"context"

	"github.com/antihax/optional"
	phrase "github.com/phrase/phrase-go/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePhraseMember() *plugin.Table {
	return &plugin.Table{
		Name:        "phrase_member",
		Description: "Get all users active in the account.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("account_id"),
			Hydrate:    listMember,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"account_id", "id"}),
			Hydrate:    getMember,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
				Description: "Unique id of the user.",
			},
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Description: "Account id associated to the user.",
				Transform:   transform.FromQual("account_id"),
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "Email of the user.",
			},
			{
				Name:        "username",
				Type:        proto.ColumnType_STRING,
				Description: "Username of the user.",
			},
			{
				Name:        "role",
				Type:        proto.ColumnType_STRING,
				Description: "Role of the user.",
			},
			{
				Name:        "default_locale_codes",
				Type:        proto.ColumnType_JSON,
				Description: "Languages user has access to.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Creation date of the user.",
			},
			{
				Name:        "last_activity_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Last activity date of the user.",
			},
		},
	}
}

func listMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_member.listMember", "connection_error", err)
		return nil, err
	}
	opts := &phrase.MembersListOpts{
		Page:    optional.NewInt32(0),
		PerPage: optional.NewInt32(100),
	}
	account_id := d.EqualsQuals["account_id"].GetStringValue()
	for {
		members, response, err := client.MembersApi.MembersList(*authContext, account_id, opts)
		if err != nil {
			plugin.Logger(*authContext).Error("phrase_member.listMember", err)
			return nil, err
		}
		for _, member := range members {
			d.StreamListItem(*authContext, member)
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

func getMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, authContext, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_member.getMember", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	account_id := quals["account_id"].GetStringValue()
	id := quals["id"].GetStringValue()
	result, _, err := client.MembersApi.MemberShow(*authContext, account_id, id, nil)
	if err != nil {
		plugin.Logger(*authContext).Error("phrase_member.getMember", err)
		return nil, err
	}
	return result, nil
}

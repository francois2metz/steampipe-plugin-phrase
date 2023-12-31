package phrase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-phrase",
		DefaultTransform: transform.FromGo().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"phrase_account": tablePhraseAccount(),
			"phrase_key":     tablePhraseKey(),
			"phrase_locale":  tablePhraseLocale(),
			"phrase_member":  tablePhraseMember(),
			"phrase_project": tablePhraseProject(),
			"phrase_tag":     tablePhraseTag(),
		},
	}
	return p
}

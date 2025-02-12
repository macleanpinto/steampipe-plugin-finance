package finance

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-finance",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"finance_us_sec_public_company": tableFinanceUsSecPublicCompany(ctx),
			"finance_us_sec_filer":          tableFinanceUsSecFiler(ctx),
			"finance_quote":                 tableFinanceQuote(ctx),
			"finance_quote_daily":           tableFinanceQuoteDaily(ctx),
			"finance_quote_hourly":          tableFinanceQuoteHourly(ctx),
			"finance_eps_history":           tableFinanceEpsHistory(ctx),
		},
	}
	return p
}

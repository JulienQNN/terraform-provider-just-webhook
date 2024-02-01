package provider

import (
	"context"

	"github.com/JulienQNN/jwb-client-go"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &justWebhooksProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &justWebhooksProvider{
			version: version,
		}
	}
}

// justWebhooksProvider is the provider implementation.
type justWebhooksProvider struct {
	version string
}

// Metadata returns the provider type name.
func (p *justWebhooksProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "jwb"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *justWebhooksProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Configure client for resources.
func (p *justWebhooksProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Debug(ctx, "Creatin client")

	client, err := jwb.NewClient()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create API Client",
			"An unexpected error occurred when creating the API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Client Error: "+err.Error(),
		)
		return
	}
	resp.ResourceData = client

	tflog.Info(ctx, "Configured  client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *justWebhooksProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *justWebhooksProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTeamsWebhookResource,
		NewSlackWebhookMessageResource,
		NewSlackWebhookAttachmentResource,
	}
}

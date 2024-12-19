package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &ExampleProvider{}
)

// ExampleProvider defines the provider implementation
type ExampleProvider struct{}

// New creates a new provider instance
func New() provider.Provider {
	return &ExampleProvider{}
}

// Metadata provides provider metadata
func (p *ExampleProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "example"
}

// Schema defines the provider configuration schema
func (p *ExampleProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	// No configuration needed for this simple provider
	resp.Schema = schema.Schema{}
}

// Configure prepares the provider
func (p *ExampleProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// No configuration to process in this simple example
}

func (p *ExampleProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSiteResource,
		NewSiteV3Resource,
	}
}

func (p *ExampleProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

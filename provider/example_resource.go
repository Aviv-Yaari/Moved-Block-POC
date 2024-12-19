package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	//"go/types"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ resource.Resource              = &SiteResource{}
	_ resource.ResourceWithConfigure = &SiteResource{}
)

// SiteResourceModel describes the resource data model
type SiteResourceModel struct {
	ID          types.String `tfsdk:"id"`
	DisplayName types.String `tfsdk:"display_name"`
	Description types.String `tfsdk:"description"`
}

// SiteResource is the resource implementation
type SiteResource struct{}

func (r *SiteResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	//TODO implement me
}

// NewSiteResource creates a new resource
func NewSiteResource() resource.Resource {
	return &SiteResource{}
}

// Metadata returns the resource type name
func (r *SiteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	//resp.TypeName = req.ProviderTypeName + "_resource"
	resp.TypeName = "example_resource"
}

// Schema defines the schema for the resource
func (r *SiteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Example resource for demonstration",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "Unique identifier for the resource",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "Display name of the example resource",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Description of the example resource",
			},
		},
	}
}

// Create creates a new resource
func (r *SiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan SiteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Generate ID (in a real provider, this would be a real unique identifier)
	//plan.ID = types.StringValue(fmt.Sprintf("example-%s", plan.DisplayName.ValueString()))
	plan.ID = types.StringValue(generateRandomID())

	// Set state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the resource state
func (r *SiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Retrieve current state
	var state SiteResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Ensure ID is not empty
	if state.ID.IsNull() || state.ID.ValueString() == "" {
		state.ID = types.StringValue(generateRandomID())
	}
	// In a real provider, you would fetch the actual resource here
	// For this example, we’ll just set the state back
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// Update updates an existing resource
func (r *SiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan SiteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Ensure ID is set
	var currentState SiteResourceModel
	_ = req.State.Get(ctx, &currentState)
	plan.ID = currentState.ID

	// Set state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Delete deletes the resource
func (r *SiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve current state
	var state SiteResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// In a real provider, you would delete the actual resource here
	// For this example, we’ll just clear the state
	resp.State.RemoveResource(ctx)
}

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
	//"go/types"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ resource.Resource              = &SiteV3Resource{}
	_ resource.ResourceWithMoveState = &SiteV3Resource{}
)

// SiteV3ResourceModel describes the resource data model
type SiteV3ResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

// SiteV3Resource is the resource implementation
type SiteV3Resource struct{}

func (r *SiteV3Resource) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{
		{
			SourceSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id":           schema.StringAttribute{},
					"display_name": schema.StringAttribute{},
					"description":  schema.StringAttribute{},
				},
			},
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				// Always verify the expected source before working with the data.
				fmt.Printf("[DEBUG]: SourceTypeName is %s\n", req.SourceTypeName)
				if req.SourceTypeName != "example_resource" {
					return
				}

				fmt.Printf("[DEBUG]: SourceSchemaVersion is %d\n", req.SourceSchemaVersion)
				if req.SourceSchemaVersion != 0 {
					return
				}

				// This only checks the provider address namespace and type
				// since practitioners may use differing hostnames for the same
				// provider, such as a network mirror. If necessary though, the
				// hostname can be used for disambiguation.
				fmt.Printf("[DEBUG]: req.SourceProviderAddress is %s\n", req.SourceProviderAddress)
				if !strings.HasSuffix(req.SourceProviderAddress, "edu/example") {
					return
				}

				var sourceStateData SiteResourceModel

				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceStateData)...)

				if resp.Diagnostics.HasError() {
					return
				}

				fmt.Printf("[DEBUG]: SourceStateData is %v\n", sourceStateData)
				fmt.Printf("[DEBUG]: SourceStateData.ID is %s\n", sourceStateData.ID)

				targetStateData := SiteV3ResourceModel{
					ID:          sourceStateData.ID,
					Name:        sourceStateData.DisplayName,
					Description: sourceStateData.Description,
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetStateData)...)
			},
		},
	}
}

func (r *SiteV3Resource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	//TODO implement me
}

// NewSiteV3Resource creates a new resource
func NewSiteV3Resource() resource.Resource {
	return &SiteV3Resource{}
}

// Metadata returns the resource type name
func (r *SiteV3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	//resp.TypeName = req.ProviderTypeName + "_resource"
	resp.TypeName = "example_v3_resource"
}

// Schema defines the schema for the resource
func (r *SiteV3Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Example resource for demonstration",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the resource",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the example resource",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Description of the example resource",
			},
		},
	}
}

// Create creates a new resource
func (r *SiteV3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	fmt.Printf("[DEBUG]: Called create v3\n")

	// Retrieve values from plan
	var plan SiteV3ResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Generate ID (in a real provider, this would be a real unique identifier)
	//plan.ID = types.StringValue(fmt.Sprintf("example-%s", plan.Name.ValueString()))
	fmt.Printf("[DEBUG]: Called create v3 with params %v\n", plan)

	plan.ID = types.StringValue(generateRandomID())
	// Set state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the resource state
func (r *SiteV3Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	fmt.Printf("[DEBUG]: Called read v3\n")
	// Retrieve current state
	var state SiteV3ResourceModel
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
func (r *SiteV3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan SiteV3ResourceModel
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
func (r *SiteV3Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve current state
	var state SiteV3ResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// In a real provider, you would delete the actual resource here
	// For this example, we’ll just clear the state
	resp.State.RemoveResource(ctx)
}

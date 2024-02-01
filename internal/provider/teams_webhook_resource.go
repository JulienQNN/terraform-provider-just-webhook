package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/JulienQNN/jwb-client-go"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &TeamsWebhookResource{}
)

func NewTeamsWebhookResource() resource.Resource {
	return &TeamsWebhookResource{}
}

// TeamsWebhookResource is the resource implementation.
type TeamsWebhookResource struct {
	client *jwb.Client
}

type teamsWebhookResourceModel struct {
	LastUpdated     types.String                `tfsdk:"last_updated"`
	WebhookUrl      types.String                `tfsdk:"webhook_url"`
	ThemeColor      types.String                `tfsdk:"theme_color"`
	Section         []teamsSectionModel         `tfsdk:"sections"`
	PotentialAction []teamsPotentialActionModel `tfsdk:"potential_action"`
}

type teamsSectionModel struct {
	ActivityTitle    types.String     `tfsdk:"title"`
	ActivitySubtitle types.String     `tfsdk:"subtitle"`
	ActivityImage    types.String     `tfsdk:"image"`
	Text             types.String     `tfsdk:"text"`
	Facts            []teamsFactModel `tfsdk:"facts"`
	Markdown         types.Bool       `tfsdk:"markdown"`
}

type teamsFactModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type teamsPotentialActionModel struct {
	Name    types.String       `tfsdk:"name"`
	Targets []teamsTargetModel `tfsdk:"targets"`
}

type teamsTargetModel struct {
	Os  types.String `tfsdk:"os"`
	Uri types.String `tfsdk:"uri"`
}

// Metadata returns the resource type name.
func (r *TeamsWebhookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook_teams"
}

// Schema defines the schema for the resource.
func (r *TeamsWebhookResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"last_updated": schema.StringAttribute{
				MarkdownDescription: "The last time the webhook was updated",
				Computed:            true,
			},
			"webhook_url": schema.StringAttribute{
				MarkdownDescription: "The webhook url to send the message to",
				Required:            true,
			},
			"theme_color": schema.StringAttribute{
				MarkdownDescription: "The theme color of the message, in hex format without the #",
				Optional:            true,
			},
			"sections": schema.ListNestedAttribute{
				MarkdownDescription: "The section(s) of the message",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"title": schema.StringAttribute{
							MarkdownDescription: "The title of the section",
							Optional:            true,
						},
						"subtitle": schema.StringAttribute{
							MarkdownDescription: "The subtitle of the section",
							Optional:            true,
						},
						"image": schema.StringAttribute{
							MarkdownDescription: "The image of the section",
							Optional:            true,
						},
						"text": schema.StringAttribute{
							MarkdownDescription: "The text of the section",
							Optional:            true,
						},
						"markdown": schema.BoolAttribute{
							MarkdownDescription: "Whether the text of the section is markdown",
							Optional:            true,
						},
						"facts": schema.ListNestedAttribute{
							MarkdownDescription: "The fact(s) of the section",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "The name of the fact",
										Required:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "The value of the fact",
										Required:            true,
									},
								},
							},
						},
					},
				},
			},
			"potential_action": schema.ListNestedAttribute{
				MarkdownDescription: "The potential action(s) buttons of the message",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the potential action",
							Required:            true,
						},
						"targets": schema.ListNestedAttribute{
							MarkdownDescription: "The target(s) of the potential action",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"os": schema.StringAttribute{
										MarkdownDescription: "The os of the target, 'default' for OpenUri links",
										Optional:            true,
									},
									"uri": schema.StringAttribute{
										MarkdownDescription: "The uri of the target",
										Required:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *TeamsWebhookResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*jwb.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf(
				"Expected *jwb.Client, got: %T. Please report this issue to the provider developers.",
				req.ProviderData,
			),
		)

		return
	}

	r.client = client
}

func (r *TeamsWebhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan teamsWebhookResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	theme_color := plan.ThemeColor

	var sections []jwb.TeamsSection
	for _, section := range plan.Section {
		var facts []jwb.TeamsFact

		for _, fact := range section.Facts {
			facts = append(facts, jwb.TeamsFact{
				Name:  fact.Name.ValueString(),
				Value: fact.Value.ValueString(),
			})
		}

		sections = append(sections, jwb.TeamsSection{
			ActivityTitle:    section.ActivityTitle.ValueString(),
			ActivitySubtitle: section.ActivitySubtitle.ValueString(),
			ActivityImage:    section.ActivityImage.ValueString(),
			Text:             section.Text.ValueString(),
			Facts:            facts,
			Markdown:         section.Markdown.ValueBool(),
		})
	}

	var potentialActions []jwb.TeamsPotentialAction
	for _, potentialAction := range plan.PotentialAction {
		var targets []jwb.TeamsTarget

		for _, target := range potentialAction.Targets {
			targets = append(targets, jwb.TeamsTarget{
				Os:  target.Os.ValueString(),
				Uri: target.Uri.ValueString(),
			})
		}

		potentialActions = append(potentialActions, jwb.TeamsPotentialAction{
			Name:    potentialAction.Name.ValueString(),
			Targets: targets,
			Type:    "OpenUri",
		},
		)

	}

	item := jwb.TeamsPayloadWebhook{
		Type:                 "MessageCard",
		Context:              "http://schema.org/extensions",
		ThemeColor:           theme_color.ValueString(),
		Summary:              "Summary text",
		Sections:             sections,
		TeamsPotentialAction: potentialActions,
	}

	_, err := r.client.CreateTeamsWebhook(plan.WebhookUrl.ValueString(), item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating webhook",
			"Could not create webhook, unexpected error: "+err.Error(),
		)
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *TeamsWebhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *TeamsWebhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan teamsWebhookResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	theme_color := plan.ThemeColor

	var sections []jwb.TeamsSection
	for _, section := range plan.Section {
		var facts []jwb.TeamsFact

		for _, fact := range section.Facts {
			facts = append(facts, jwb.TeamsFact{
				Name:  fact.Name.ValueString(),
				Value: fact.Value.ValueString(),
			})
		}

		sections = append(sections, jwb.TeamsSection{
			ActivityTitle:    section.ActivityTitle.ValueString(),
			ActivitySubtitle: section.ActivitySubtitle.ValueString(),
			ActivityImage:    section.ActivityImage.ValueString(),
			Text:             section.Text.ValueString(),
			Facts:            facts,
			Markdown:         section.Markdown.ValueBool(),
		})
	}

	var potentialActions []jwb.TeamsPotentialAction
	for _, potentialAction := range plan.PotentialAction {
		var targets []jwb.TeamsTarget

		for _, target := range potentialAction.Targets {
			targets = append(targets, jwb.TeamsTarget{
				Os:  target.Os.ValueString(),
				Uri: target.Uri.ValueString(),
			})
		}

		potentialActions = append(potentialActions, jwb.TeamsPotentialAction{
			Name:    potentialAction.Name.ValueString(),
			Targets: targets,
			Type:    "OpenUri",
		},
		)

	}

	item := jwb.TeamsPayloadWebhook{
		Type:                 "MessageCard",
		Context:              "http://schema.org/extensions",
		ThemeColor:           theme_color.ValueString(),
		Summary:              "Summary text",
		Sections:             sections,
		TeamsPotentialAction: potentialActions,
	}

	_, err := r.client.CreateTeamsWebhook(plan.WebhookUrl.ValueString(), item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating webhook",
			"Could not create webhook, unexpected error: "+err.Error(),
		)
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *TeamsWebhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

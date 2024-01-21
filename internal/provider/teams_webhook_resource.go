package provider

import (
	"context"
	"fmt"

	"github.com/JulienQNN/jwb-client-go"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &TeamsWebhookResource{}
)

func NewWebhookResource() resource.Resource {
	return &TeamsWebhookResource{}
}

// TeamsWebhookResource is the resource implementation.
type TeamsWebhookResource struct {
	client *jwb.Client
}

type teamsWebhookResourceModel struct {
	WebhookUrl      string                 `tfsdk:"webhook_url"`
	ThemeColor      string                 `tfsdk:"theme_color"`
	Section         []sectionModel         `tfsdk:"sections"`
	PotentialAction []potentialActionModel `tfsdk:"potential_action"`
}

type sectionModel struct {
	ActivityTitle    *string     `tfsdk:"title"`
	ActivitySubtitle *string     `tfsdk:"subtitle"`
	ActivityImage    *string     `tfsdk:"image"`
	Text             *string     `tfsdk:"text"`
	Facts            []factModel `tfsdk:"facts"`
	Markdown         *bool       `tfsdk:"markdown"`
}

type factModel struct {
	Name  string `tfsdk:"name"`
	Value string `tfsdk:"value"`
}

type potentialActionModel struct {
	Name    string        `tfsdk:"name"`
	Targets []targetModel `tfsdk:"targets"`
}

type targetModel struct {
	Os  string `tfsdk:"os"`
	Uri string `tfsdk:"uri"`
}

// Metadata returns the resource type name.
func (r *TeamsWebhookResource) Metadata(
	_ context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_webhook_teams"
}

// Schema defines the schema for the resource.
func (r *TeamsWebhookResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"webhook_url": schema.StringAttribute{
				MarkdownDescription: "The webhook url to send the message to",
				Required:            true,
			},
			"theme_color": schema.StringAttribute{
				MarkdownDescription: "The theme color of the message",
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
							Optional:            true,
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
										Optional:            true,
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

func (r *TeamsWebhookResource) Configure(
	_ context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
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

func (r *TeamsWebhookResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	// Retrieve values from plan
	var plan teamsWebhookResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	theme_color := plan.ThemeColor

	var sections []jwb.Section
	for _, section := range plan.Section {
		var facts []jwb.Fact

		for _, fact := range section.Facts {
			facts = append(facts, jwb.Fact{
				Name:  fact.Name,
				Value: fact.Value,
			})
		}

		sections = append(sections, jwb.Section{
			ActivityTitle:    section.ActivityTitle,
			ActivitySubtitle: section.ActivitySubtitle,
			ActivityImage:    section.ActivityImage,
			Text:             section.Text,
			Facts:            facts,
			Markdown:         section.Markdown,
		})
	}

	var potentialActions []jwb.PotentialActionIntermediate
	for _, potentialAction := range plan.PotentialAction {
		var targets []jwb.Target

		for _, target := range potentialAction.Targets {
			targets = append(targets, jwb.Target{
				Os:  target.Os,
				Uri: target.Uri,
			})
		}

		potentialActions = append(potentialActions, jwb.PotentialActionIntermediate{
			Name:    potentialAction.Name,
			Targets: targets,
			Type:    "OpenUri",
		},
		)

	}

	item := jwb.TeamsPayloadWebhook{
		Type:            "MessageCard",
		Context:         "http://schema.org/extensions",
		ThemeColor:      theme_color,
		Summary:         "Summary text",
		Sections:        sections,
		PotentialAction: potentialActions,
	}

	_, err := r.client.CreateTeamsWebhook(plan.WebhookUrl, item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating webhook",
			"Could not create webhook, unexpected error: "+err.Error(),
		)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *TeamsWebhookResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *TeamsWebhookResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *TeamsWebhookResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
}

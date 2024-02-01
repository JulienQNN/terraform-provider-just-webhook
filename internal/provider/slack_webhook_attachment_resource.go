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
	_ resource.Resource = &SlackWebhookAttachmentResource{}
)

func NewSlackWebhookAttachmentResource() resource.Resource {
	return &SlackWebhookAttachmentResource{}
}

// SlackWebhookResource is the resource implementation.
type SlackWebhookAttachmentResource struct {
	client *jwb.Client
}

type slackWebhookAttachmentResourceModel struct {
	LastUpdated types.String                    `tfsdk:"last_updated"`
	WebhookUrl  types.String                    `tfsdk:"webhook_url"`
	Attachment  []*slackAttachmentResourceModel `tfsdk:"attachments"`
}

type slackAttachmentResourceModel struct {
	Block []*slackAttachmentBlockModel `tfsdk:"blocks"`
	Color types.String                 `tfsdk:"color"`
}

type slackAttachmentBlockModel struct {
	Type      types.String                   `tfsdk:"type"`
	Text      *slackAttachmentTextModel      `tfsdk:"text"`
	Title     *slackAttachmentTitleModel     `tfsdk:"title"`
	ImageUrl  types.String                   `tfsdk:"image_url"`
	AltText   types.String                   `tfsdk:"alt_text"`
	Accessory *slackAttachmentAccessoryModel `tfsdk:"accessory"`
	Fields    []*slackAttachmentFieldModel   `tfsdk:"fields"`
	Elements  []*slackAttachmentElementModel `tfsdk:"elements"`
}

type slackAttachmentTextModel struct {
	Type types.String `tfsdk:"type"`
	Text types.String `tfsdk:"text"`
}

type slackAttachmentTitleModel struct {
	Type  types.String `tfsdk:"type"`
	Text  types.String `tfsdk:"text"`
	Emoji types.Bool   `tfsdk:"emoji"`
}

type slackAttachmentFieldModel struct {
	Type types.String `tfsdk:"type"`
	Text types.String `tfsdk:"text"`
}

type slackAttachmentAccessoryModel struct {
	Type types.String                      `tfsdk:"type"`
	Text slackAttachmentAccessoryTextModel `tfsdk:"text"`
}

type slackAttachmentAccessoryTextModel struct {
	Type  types.String `tfsdk:"type"`
	Text  types.String `tfsdk:"text"`
	Emoji types.Bool   `tfsdk:"emoji"`
}

type slackAttachmentElementModel struct {
	Type  types.String                    `tfsdk:"type"`
	Text  slackAttachmentElementTextModel `tfsdk:"text"`
	Style types.String                    `tfsdk:"style"`
	Url   types.String                    `tfsdk:"url"`
}

type slackAttachmentElementTextModel struct {
	Type  types.String `tfsdk:"type"`
	Text  types.String `tfsdk:"text"`
	Emoji types.Bool   `tfsdk:"emoji"`
}

// Metadata returns the resource type name.
func (r *SlackWebhookAttachmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook_slack_attachment"
}

// Schema defines the schema for the resource.
func (r *SlackWebhookAttachmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"last_updated": schema.StringAttribute{
				MarkdownDescription: "the last time the webhook was updated",
				Computed:            true,
			},
			"webhook_url": schema.StringAttribute{
				MarkdownDescription: "the webhook url to send the message to",
				Required:            true,
			},
			"attachments": schema.ListNestedAttribute{
				MarkdownDescription: "the attachment(s) of the message",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"color": schema.StringAttribute{
							MarkdownDescription: "the color of the attachment",
							Optional:            true,
						},
						"blocks": schema.ListNestedAttribute{
							MarkdownDescription: "the block(s) of the attachment",
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "the type of the block (section, divider, image, actions)",
										Required:            true,
									},
									"text": schema.SingleNestedAttribute{
										MarkdownDescription: "the text of the block",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												MarkdownDescription: "the type of the text",
												Required:            true,
											},
											"text": schema.StringAttribute{
												MarkdownDescription: "the text of the text",
												Required:            true,
											},
										},
									},
									"title": schema.SingleNestedAttribute{
										MarkdownDescription: "the text of the block",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												MarkdownDescription: "the type of the text",
												Required:            true,
											},
											"text": schema.StringAttribute{
												MarkdownDescription: "the text of the text",
												Required:            true,
											},
											"emoji": schema.BoolAttribute{
												MarkdownDescription: "whether or not to use emoji in the text",
												Optional:            true,
											},
										},
									},
									// Neeed a type "image"
									"image_url": schema.StringAttribute{
										MarkdownDescription: "the url of the image",
										Optional:            true,
									},
									"alt_text": schema.StringAttribute{
										MarkdownDescription: "the alt text of the image",
										Optional:            true,
									},
									"fields": schema.ListNestedAttribute{
										MarkdownDescription: "the fields of the block",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "the type of the field",
													Required:            true,
												},
												"text": schema.StringAttribute{
													MarkdownDescription: "the text of the field",
													Required:            true,
												},
											},
										},
									},
									"accessory": schema.SingleNestedAttribute{
										MarkdownDescription: "the accessory of the block",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												MarkdownDescription: "the type of the accessory",
												Required:            true,
											},
											"text": schema.SingleNestedAttribute{
												MarkdownDescription: "the text of the accessory",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														MarkdownDescription: "the type of the text",
														Required:            true,
													},
													"text": schema.StringAttribute{
														MarkdownDescription: "the text of the text",
														Required:            true,
													},
													"emoji": schema.BoolAttribute{
														MarkdownDescription: "whether or not to use emoji in the text",
														Optional:            true,
													},
												},
											},
										},
									},
									"elements": schema.ListNestedAttribute{
										MarkdownDescription: "the elements of the block",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "the type of the element",
													Required:            true,
												},
												"text": schema.SingleNestedAttribute{
													MarkdownDescription: "the text of the element",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"type": schema.StringAttribute{
															MarkdownDescription: "the type of the text",
															Required:            true,
														},
														"text": schema.StringAttribute{
															MarkdownDescription: "the text of the element",
															Required:            true,
														},
														"emoji": schema.BoolAttribute{
															MarkdownDescription: "whether or not to use emoji in the text",
															Optional:            true,
														},
													},
												},
												"style": schema.StringAttribute{
													MarkdownDescription: "the style of the element (primary or danger)",
													Optional:            true,
													Computed:            true,
												},
												"url": schema.StringAttribute{
													MarkdownDescription: "the url of the element",
													Optional:            true,
												},
											},
										},
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

func (r *SlackWebhookAttachmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SlackWebhookAttachmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan slackWebhookAttachmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var attachments []jwb.SlackAttachment

	for _, attachment := range plan.Attachment {

		var blocks []jwb.SlackBlock

		for _, block := range attachment.Block {

			var text *jwb.SlackText
			if block.Text != nil {
				text = &jwb.SlackText{
					Type: block.Text.Type.ValueString(),
					Text: block.Text.Text.ValueString(),
				}
			}

			var title *jwb.SlackTitle
			if block.Title != nil {
				title = &jwb.SlackTitle{
					Type:  block.Title.Type.ValueString(),
					Text:  block.Title.Text.ValueString(),
					Emoji: block.Title.Emoji.ValueBool(),
				}
			}

			var fields []jwb.SlackField

			for _, field := range block.Fields {
				fields = append(fields, jwb.SlackField{
					Type: field.Type.ValueString(),
					Text: field.Text.ValueString(),
				})
			}

			var accessory *jwb.SlackAccessory
			if block.Accessory != nil {
				accessory = &jwb.SlackAccessory{
					Type: block.Accessory.Type.ValueString(),
					Text: &jwb.SlackAccessoryText{
						Type:  block.Accessory.Text.Type.ValueString(),
						Text:  block.Accessory.Text.Text.ValueString(),
						Emoji: block.Accessory.Text.Emoji.ValueBool(),
					},
				}
			}

			var elements []jwb.SlackElement

			for _, element := range block.Elements {
				elements = append(elements, jwb.SlackElement{
					Type: element.Type.ValueString(),
					Text: jwb.SlackElementText{
						Type:  element.Text.Type.ValueString(),
						Text:  element.Text.Text.ValueString(),
						Emoji: element.Text.Emoji.ValueBool(),
					},
					Style: element.Style.ValueString(),
					Url:   element.Url.ValueString(),
				})
			}

			validateBlock := ValidateSlackAttachmentBlock(*block)

			if validateBlock.HasError() {
				resp.Diagnostics.Append(validateBlock...)
				return
			}

			blocks = append(blocks, jwb.SlackBlock{
				Type:      block.Type.ValueString(),
				Text:      text,
				Title:     title,
				ImageUrl:  block.ImageUrl.ValueString(),
				AltText:   block.AltText.ValueString(),
				Fields:    fields,
				Accessory: accessory,
				Elements:  elements,
			})
		}

		validateAttachment := ValidateSlackAttachmentColor(*attachment)
		if validateAttachment.HasError() {
			resp.Diagnostics.Append(validateAttachment...)
			return
		}

		attachments = append(attachments, jwb.SlackAttachment{
			Color:  attachment.Color.ValueString(),
			Blocks: blocks,
		})
	}

	item := jwb.SlackAttachmentPayloadWebhook{
		Attachments: attachments,
	}

	_, err := r.client.CreateSlackAttachmentWebhook(plan.WebhookUrl.ValueString(), item)
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

func (r *SlackWebhookAttachmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *SlackWebhookAttachmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan slackWebhookAttachmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var attachments []jwb.SlackAttachment

	for _, attachment := range plan.Attachment {

		var blocks []jwb.SlackBlock

		for _, block := range attachment.Block {

			var text *jwb.SlackText
			if block.Text != nil {
				text = &jwb.SlackText{
					Type: block.Text.Type.ValueString(),
					Text: block.Text.Text.ValueString(),
				}
			}

			var title *jwb.SlackTitle
			if block.Title != nil {
				title = &jwb.SlackTitle{
					Type:  block.Title.Type.ValueString(),
					Text:  block.Title.Text.ValueString(),
					Emoji: block.Title.Emoji.ValueBool(),
				}
			}

			var fields []jwb.SlackField

			for _, field := range block.Fields {
				fields = append(fields, jwb.SlackField{
					Type: field.Type.ValueString(),
					Text: field.Text.ValueString(),
				})
			}

			var accessory *jwb.SlackAccessory
			if block.Accessory != nil {
				accessory = &jwb.SlackAccessory{
					Type: block.Accessory.Type.ValueString(),
					Text: &jwb.SlackAccessoryText{
						Type:  block.Accessory.Text.Type.ValueString(),
						Text:  block.Accessory.Text.Text.ValueString(),
						Emoji: block.Accessory.Text.Emoji.ValueBool(),
					},
				}
			}

			var elements []jwb.SlackElement

			for _, element := range block.Elements {
				elements = append(elements, jwb.SlackElement{
					Type: element.Type.ValueString(),
					Text: jwb.SlackElementText{
						Type:  element.Text.Type.ValueString(),
						Text:  element.Text.Text.ValueString(),
						Emoji: element.Text.Emoji.ValueBool(),
					},
					Style: element.Style.ValueString(),
					Url:   element.Url.ValueString(),
				})
			}

			validateBlock := ValidateSlackAttachmentBlock(*block)

			if validateBlock.HasError() {
				resp.Diagnostics.Append(validateBlock...)
				return
			}

			blocks = append(blocks, jwb.SlackBlock{
				Type:      block.Type.ValueString(),
				Text:      text,
				Title:     title,
				ImageUrl:  block.ImageUrl.ValueString(),
				AltText:   block.AltText.ValueString(),
				Fields:    fields,
				Accessory: accessory,
				Elements:  elements,
			})
		}

		validateAttachment := ValidateSlackAttachmentColor(*attachment)
		if validateAttachment.HasError() {
			resp.Diagnostics.Append(validateAttachment...)
			return
		}

		attachments = append(attachments, jwb.SlackAttachment{
			Color:  attachment.Color.ValueString(),
			Blocks: blocks,
		})
	}

	item := jwb.SlackAttachmentPayloadWebhook{
		Attachments: attachments,
	}

	_, err := r.client.CreateSlackAttachmentWebhook(plan.WebhookUrl.ValueString(), item)
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
func (r *SlackWebhookAttachmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

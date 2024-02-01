package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ValidateSlackMessageBlock(block slackMessageBlockModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if block.Fields != nil && block.Type.ValueString() != "section" {
		diags.AddError(
			"Error Schema Validation",
			fmt.Sprintf("If 'fields' is used, then 'type' of the block need to be 'section', but the provider received: %s", block.Type.ValueString()),
		)
	}

	if block.Type.ValueString() == "header" && block.Text == nil {
		diags.AddError(
			"Error Schema Validation",
			fmt.Sprintf("If 'type' is 'header', then 'block' need 'text' parameter."),
		)
	}

	if block.Type.ValueString() == "divider" && (block.Text != nil || block.Title != nil || block.ImageUrl.ValueString() != "" || block.AltText.ValueString() != "" || block.Accessory != nil || len(block.Fields) > 0 || len(block.Elements) > 0) {
		diags.AddError(
			"Error Schema Validation",
			fmt.Sprintf("If 'type' is 'divider', then it should be the only field in the 'block'."),
		)
	}

	if block.Type.ValueString() != "image" && (block.ImageUrl.ValueString() != "" || block.AltText.ValueString() != "") {
		diags.AddError(
			"Error Schema Validation",
			fmt.Sprintf("If 'image_url' is used, then 'type' of the block need to be 'image', but the provider received: %s", block.Type.ValueString()),
		)
	}

	if block.AltText.ValueString() != "" && block.ImageUrl.ValueString() == "" {
		diags.AddError(
			"Error Schema Validation",
			"If 'alt_text' is used, then 'image_url' must also be used.",
		)
	}

	return diags
}

func ValidateSlackAttachmentColor(attachment slackAttachmentResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if attachment.Color.ValueString() != "" && attachment.Color.ValueString()[0] != '#' {
		diags.AddError(
			"Error Schema Validation",
			"Color must be in hex format (e.g. #FF0000).",
		)
	}

	return diags
}

func ValidateSlackAttachmentBlock(block slackAttachmentBlockModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if block.Fields != nil && block.Type.ValueString() != "section" {
		diags.AddError(
			"Error Schema Validation",
			fmt.Sprintf("If 'fields' is used, then 'type' of the block need to be 'section', but the provider received: %s", block.Type.ValueString()),
		)
	}

	if block.Type.ValueString() == "header" && block.Text == nil {
		diags.AddError(
			"Error Schema Validation",
			fmt.Sprintf("If 'type' is 'header', then 'block' need 'text' parameter."),
		)
	}

	if block.Type.ValueString() == "divider" && (block.Text != nil || block.Title != nil || block.ImageUrl.ValueString() != "" || block.AltText.ValueString() != "" || block.Accessory != nil || len(block.Fields) > 0 || len(block.Elements) > 0) {
		diags.AddError(
			"Error Schema Validation",
			fmt.Sprintf("If 'type' is 'divider', then it should be the only field in the 'block'."),
		)
	}

	if block.Type.ValueString() != "image" && (block.ImageUrl.ValueString() != "" || block.AltText.ValueString() != "") {
		diags.AddError(
			"Error Schema Validation",
			fmt.Sprintf("If 'image_url' is used, then 'type' of the block need to be 'image', but the provider received: %s", block.Type.ValueString()),
		)
	}

	if block.AltText.ValueString() != "" && block.ImageUrl.ValueString() == "" {
		diags.AddError(
			"Error Schema Validation",
			"If 'alt_text' is used, then 'image_url' must also be used.",
		)
	}
	return diags
}

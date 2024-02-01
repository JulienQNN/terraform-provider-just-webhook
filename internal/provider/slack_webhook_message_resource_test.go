package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSlackWebhookMessageResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
resource "jwb_webhook_slack_message" "simple_with_image_and_link" {
  webhook_url = "https://webhook.site/f33cf6e1-32fe-4317-b60b-7c249b482bb2"
  blocks = [
    {
      type = "section"
      text = {
        type = "mrkdwn"
        text = "Hello, World!"
      }
      fields : [{
        "type" : "mrkdwn",
        "text" : "Some text for this field"
      }]
    },
    {
      type : "divider"
    },
    {
      type = "section"
      text = {
        type = "mrkdwn"
        text = "<https://google.com|Google is your friend, maybee>"
      },
    }
  ]
}`, Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "webhook_url", "https://webhook.site/f33cf6e1-32fe-4317-b60b-7c249b482bb2"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.0.type", "section"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.0.text.type", "mrkdwn"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.0.text.text", "Hello, World!"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.0.fields.0.type", "mrkdwn"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.0.fields.0.text", "Some text for this field"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.1.type", "divider"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.2.type", "section"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.2.text.type", "mrkdwn"),
					resource.TestCheckResourceAttr("jwb_webhook_slack_message.simple_with_image_and_link", "blocks.2.text.text", "<https://google.com|Google is your friend, maybee>"),

					// Verify dynamic values have any value set in the state.
					// resource.TestCheckResourceAttrSet("jwb_webhook_teams.test", "Type"),
					// resource.TestCheckResourceAttrSet("jwb_webhook_teams.test", "last_updated"),
				),
			},
		},
	})
}

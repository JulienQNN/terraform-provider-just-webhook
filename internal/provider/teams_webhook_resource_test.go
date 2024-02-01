package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTeamsWebhookResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
resource "jwb_webhook_teams" "test" {
  webhook_url = "https://hooks.teams.com/services/***********/*******/****"
  theme_color = "1BEF1D"
  sections = [{
    title    = "TITLE"
    subtitle = "Subtitle 1"
    image    = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/2560px-Amazon_Web_Services_Logo.svg.png"
    text     = "Text 1"
    facts = [
      {
        name  = "Hello name 1"
        value = "## Hello value 1"
      },
    ]
    markdown = true
    },
    {
      title    = "Title 2"
      image    = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/2560px-Amazon_Web_Services_Logo.svg.png"
      subtitle = "Subtitle 2"
      text     = "Text 2"
      facts = [
        {
          name  = "Hello name 3"
          value = "## Hello value 3"
        },
        {
          name  = "Hello name 4"
          value = "## Hello value 4"
        },
      ]
      markdown = true
    },
  ]
  potential_action = [
    {
      name = "View in Azure Portal"
      targets = [{
        os  = "default"
        uri = "https://azure.microsoft.com/fr-fr/get-started/azure-portal"
      }]
    },
    {
      name = "View in AWS Portal"
      targets = [{
        os  = "os"
        uri = "https://aws.amazon.com/fr/console/"
      }]
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("jwb_webhook_teams.test", "sections.#", "2"),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"webhook_url",
						"https://hooks.teams.com/services/***********/*******/****",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"theme_color",
						"1BEF1D",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.title",
						"TITLE",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.subtitle",
						"Subtitle 1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.image",
						"https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/2560px-Amazon_Web_Services_Logo.svg.png",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.text",
						"Text 1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.facts.#",
						"1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.facts.0.name",
						"Hello name 1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.facts.0.value",
						"## Hello value 1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.markdown",
						"true",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.title",
						"Title 2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.subtitle",
						"Subtitle 2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.image",
						"https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/2560px-Amazon_Web_Services_Logo.svg.png",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.text",
						"Text 2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.#",
						"2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.0.name",
						"Hello name 3",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.0.value",
						"## Hello value 3",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.1.name",
						"Hello name 4",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.1.value",
						"## Hello value 4",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.markdown",
						"true",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.#",
						"2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.0.name",
						"View in Azure Portal",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.0.targets.#",
						"1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.0.targets.0.os",
						"default",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.0.targets.0.uri",
						"https://azure.microsoft.com/fr-fr/get-started/azure-portal",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.1.name",
						"View in AWS Portal",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.1.targets.#",
						"1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.1.targets.0.os",
						"os",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.1.targets.0.uri",
						"https://aws.amazon.com/fr/console/",
					),

					// Verify dynamic values have any value set in the state.
					// resource.TestCheckResourceAttrSet("jwb_webhook_teams.test", "Type"),
					// resource.TestCheckResourceAttrSet("jwb_webhook_teams.test", "last_updated"),
				),
			},
			// Update and Read testing
			{
				Config: `
resource "jwb_webhook_teams" "test" {
  webhook_url = "https://hooks.teams.com/services/***********/*******/****""
  theme_color = "1BEF1D"
  sections = [{
    title    = "TITLE"
    subtitle = "Subtitle 1"
    image    = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/2560px-Amazon_Web_Services_Logo.svg.png"
    text     = "Text 1"
    facts = [
      {
        name  = "Hello name 1"
        value = "## Hello value 1"
      },
    ]
    markdown = true
    },
    {
      title    = "Title 2"
      image    = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/2560px-Amazon_Web_Services_Logo.svg.png"
      subtitle = "Subtitle 2"
      text     = "Text 2"
      facts = [
        {
          name  = "Hello name 3"
          value = "## Hello value 3"
        },
        {
          name  = "Hello name 4"
          value = "## Hello value 4"
        },
      ]
      markdown = true
    },
  ]
  potential_action = [
    {
      name = "View in Azure Portal"
      targets = [{
        os  = "default"
        uri = "https://azure.microsoft.com/fr-fr/get-started/azure-portal"
      }]
    },
    {
      name = "View in AWS Portal"
      targets = [{
        os  = "os"
        uri = "https://aws.amazon.com/fr/console/"
      }]
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of items
					resource.TestCheckResourceAttr("jwb_webhook_teams.test", "sections.#", "2"),
					// Verify first order item
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"webhook_url",
						"https://hooks.teams.com/services/***********/*******/****",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"theme_color",
						"1BEF1D",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.title",
						"TITLE",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.subtitle",
						"Subtitle 1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.image",
						"https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/2560px-Amazon_Web_Services_Logo.svg.png",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.text",
						"Text 1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.facts.#",
						"1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.facts.0.name",
						"Hello name 1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.facts.0.value",
						"## Hello value 1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.0.markdown",
						"true",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.title",
						"Title 2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.subtitle",
						"Subtitle 2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.image",
						"https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/2560px-Amazon_Web_Services_Logo.svg.png",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.text",
						"Text 2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.#",
						"2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.0.name",
						"Hello name 3",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.0.value",
						"## Hello value 3",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.1.name",
						"Hello name 4",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.facts.1.value",
						"## Hello value 4",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"sections.1.markdown",
						"true",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.#",
						"2",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.0.name",
						"View in Azure Portal",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.0.targets.#",
						"1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.0.targets.0.os",
						"default",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.0.targets.0.uri",
						"https://azure.microsoft.com/fr-fr/get-started/azure-portal",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.1.name",
						"View in AWS Portal",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.1.targets.#",
						"1",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.1.targets.0.os",
						"os",
					),
					resource.TestCheckResourceAttr(
						"jwb_webhook_teams.test",
						"potential_action.1.targets.0.uri",
						"https://aws.amazon.com/fr/console/",
					),

					// Verify dynamic values have any value set in the state.
					// resource.TestCheckResourceAttrSet("jwb_webhook_teams.test", "Type"),
					// resource.TestCheckResourceAttrSet("jwb_webhook_teams.test", "last_updated"),
				),
			},
		},
	})
}

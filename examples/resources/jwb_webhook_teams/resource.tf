terraform {
  required_providers {
    jwb = {
      source  = "JulienQNN/just-webhook"
      version = "0.1.5"
    }
  }
}

provider "jwb" {
}

resource "jwb_webhook_teams" "simple" {
  webhook_url = "https://123.webhook.office.com/webhookb2/****/IncomingWebhook/****/****"
  sections = [{
    title    = "The Hello World title"
    subtitle = "The Hello World subtitle"
    image    = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/Microsoft_logo.jpg/480px-Microsoft_logo.jpg"
    text     = "The Hello World text"
    facts = [
      {
        name  = "Fact name number 1"
        value = "## Fact value number 1 without markdown"
      },
    ]
  }]
}

resource "jwb_webhook_teams" "advanced" {
  webhook_url = "https://123.webhook.office.com/webhookb2/****/IncomingWebhook/****/****"
  theme_color = "1BEF1D"
  sections = [{
    title    = "The Hello World title"
    subtitle = "The Hello World subtitle"
    image    = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/Microsoft_logo.jpg/480px-Microsoft_logo.jpg"
    text     = "The Hello World text with **markdown**"
    markdown = true
  }]
  potential_action = [
    {
      name = "The action button name 1"
      targets = [{
        os  = "default"
        uri = "https://learn.microsoft.com/fr-fr/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook?tabs=javascript"
      }]
    },
    {
      name = "The action button name 2"
      targets = [{
        os  = "default"
        uri = "https://learn.microsoft.com/fr-fr/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook?tabs=javascript"
      }]
    }
  ]
}

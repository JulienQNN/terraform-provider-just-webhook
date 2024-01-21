resource "jwb_webhook_teams" "this" {
  webhook_url = "https://123.webhook.office.com/webhookb2/****/IncomingWebhook/****/****"
  theme_color = "1BEF1D"
  sections = [{
    title    = "Title 1"
    subtitle = " subtitle 1"
    image    = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/Microsoft_logo.jpg/480px-Microsoft_logo.jpg"
    text     = "text 1"
    facts = [
      {
        name  = "name 1"
        value = "## value 1"
      },
    ]
    markdown = true
    },
    {
      title    = "title 2"
      image    = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/Microsoft_logo.jpg/480px-Microsoft_logo.jpg"
      subtitle = "subtitle 2"
      text     = "text 2"
      facts = [
        {
          name  = "name 3"
          value = "##  value 3"
        },
        {
          name  = "name 4"
          value = "## value 4"
        },
      ]
    },
  ]
  potential_action = [
    {
      name = "View in Learn Microsoft Portal"
      targets = [{
        os  = "default"
        uri = "https://learn.microsoft.com/fr-fr/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook?tabs=javascript"
      }]
    },
    {
      name = "View in Learn Microsoft Portal"
      targets = [{
        os  = "default"
        uri = "https://learn.microsoft.com/fr-fr/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook?tabs=javascript"
      }]
    }
  ]
}

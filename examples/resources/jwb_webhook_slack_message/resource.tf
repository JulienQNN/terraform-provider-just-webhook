terraform {
  required_providers {
    jwb = {
      source  = "JulienQNN/just-webhook"
      version = "0.2.3"
    }
  }
}

resource "jwb_webhook_slack_message" "simple_with_image_and_link" {
  webhook_url = "https://hooks.slack.com/services/********/*******/********************"
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
      type : "image",
      title : {
        type : "plain_text",
        text : "Example Image",
        emoji : true
      },
      image_url : "https://api.slack.com/img/blocks/bkb_template_images/beagle.png",
      alt_text : "beagle",
    },
    {
      type : "divider"
    },
    {
      type = "section"
      text = {
        type = "mrkdwn"
        text = "<https://github.com/JulienQNN/terraform-provider-just-webhook|Yes button>"
      },
    }
  ]
}


resource "jwb_webhook_slack_message" "complexe_with_slack_buttons" {
  webhook_url = "https://hooks.slack.com/services/********/*******/********************"
  blocks = [
    {
      type = "section"
      text = {
        type = "mrkdwn"
        text = "Hello, World!"
      }
      fields : [
        {
          "type" : "mrkdwn",
          "text" : "Some text for this field"
        },
        {
          "type" : "mrkdwn",
          "text" : "Some text for this second field"
        }
      ]
    },
    {
      type : "divider"
    },
    {
      type : "actions",
      elements : [
        {
          type : "button",
          text : {
            type : "plain_text",
            emoji : true,
            text : "Approve"
          },
          style : "primary",
          url : "U need to config your Slack app to use this"
        },
        {
          type : "button",
          text : {
            type : "plain_text",
            emoji : true,
            text : "Deny"
          },
          style : "danger",
          url : "U need to config your Slack app to use this"
        }
      ]
    },
  ]
}

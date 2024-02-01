---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "jwb_webhook_slack_message Resource - terraform-provider-just-webhook"
subcategory: ""
description: |-
  
---

# jwb_webhook_slack_message (Resource)



## Example Usage

```terraform
terraform {
  required_providers {
    jwb = {
      source  = "JulienQNN/just-webhook"
      version = "0.1.5"
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `webhook_url` (String) the webhook url to send the message to

### Optional

- `blocks` (Attributes List) the block(s) of the message (see [below for nested schema](#nestedatt--blocks))

### Read-Only

- `last_updated` (String) the last time the webhook was updated

<a id="nestedatt--blocks"></a>
### Nested Schema for `blocks`

Required:

- `type` (String) the type of the block (section, divider, image, actions)

Optional:

- `accessory` (Attributes) the accessory of the block (see [below for nested schema](#nestedatt--blocks--accessory))
- `alt_text` (String) the alt text of the image
- `elements` (Attributes List) the elements of the block (see [below for nested schema](#nestedatt--blocks--elements))
- `fields` (Attributes List) the fields of the block (see [below for nested schema](#nestedatt--blocks--fields))
- `image_url` (String) the url of the image
- `text` (Attributes) the text of the block (see [below for nested schema](#nestedatt--blocks--text))
- `title` (Attributes) the text of the block (see [below for nested schema](#nestedatt--blocks--title))

<a id="nestedatt--blocks--accessory"></a>
### Nested Schema for `blocks.accessory`

Required:

- `type` (String) the type of the accessory

Optional:

- `text` (Attributes) the text of the accessory (see [below for nested schema](#nestedatt--blocks--accessory--text))

<a id="nestedatt--blocks--accessory--text"></a>
### Nested Schema for `blocks.accessory.text`

Required:

- `text` (String) the text of the text
- `type` (String) the type of the text

Optional:

- `emoji` (Boolean) whether or not to use emoji in the text



<a id="nestedatt--blocks--elements"></a>
### Nested Schema for `blocks.elements`

Required:

- `type` (String) the type of the element

Optional:

- `style` (String) the style of the element (primary or danger)
- `text` (Attributes) the text of the element (see [below for nested schema](#nestedatt--blocks--elements--text))
- `url` (String) the url of the element

<a id="nestedatt--blocks--elements--text"></a>
### Nested Schema for `blocks.elements.text`

Required:

- `text` (String) the text of the element
- `type` (String) the type of the text

Optional:

- `emoji` (Boolean) whether or not to use emoji in the text



<a id="nestedatt--blocks--fields"></a>
### Nested Schema for `blocks.fields`

Required:

- `text` (String) the text of the field
- `type` (String) the type of the field


<a id="nestedatt--blocks--text"></a>
### Nested Schema for `blocks.text`

Required:

- `text` (String) the text of the text
- `type` (String) the type of the text


<a id="nestedatt--blocks--title"></a>
### Nested Schema for `blocks.title`

Required:

- `text` (String) the text of the text
- `type` (String) the type of the text

Optional:

- `emoji` (Boolean) whether or not to use emoji in the text
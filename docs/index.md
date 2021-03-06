---
layout: "wallarm"
page_title: "Provider: Wallarm"
description: |-
  The Wallarm provider is used to interact with the Wallarm WAF resources. The provider needs to be configured with the proper authentication credentials before it can be used.
---

# Wallarm Provider

The Wallarm provider is used to interact with the [Wallarm WAF](https://docs.wallarm.com/) resources. The provider needs to be configured with the proper authentication credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Sets up Wallarm authentication credentials
# You can use environment variables instead
provider "wallarm" {
  api_uuid = "4f6b8ace-9502-48c5-b1ad-526a230c2203"
  api_secret = "e3f929a2eeb4f0fe6a453fee0f81e36d935b03a5f002d740ff8aa425f6cf4647"
  api_host = "https://api.wallarm.com"
  client_id = 6039
}

}

# Adds a domain to the Scanner scope
resource "wallarm_scanner" "scope" {
  # ...
}

# Creates a rule to block the requests
resource "wallarm_rule_vpatch" "vpatch" {
  # ...
}
```

## Argument Reference

The following arguments are supported in `provider "wallarm"`:

* `api_uuid` - (Required) Your Wallarm user UUID. Note that the most operations with Wallarm API are allowed only for the users with the **Administrator** role. To get UUID, please the [instruction](https://docs.wallarm.com/admin-en/api-en/#your-own-client). This can also be specified with the `WALLARM_API_UUID` shell environment variable.
* `api_secret` - (Required) Your Wallarm user secret. Note that the most operations with Wallarm API are allowed only for the users with the **Administrator** role. To get UUID, please the [instruction](https://docs.wallarm.com/admin-en/api-en/#your-own-client). This can also be specified with the `WALLARM_API_SECRET` shell environment variable.
* `api_host` - (Optional) Wallarm API URL. Can be: `https://us1.api.wallarm.com` for the [US cloud](https://docs.wallarm.com/quickstart-en/how-wallarm-works/qs-intro-en/#us-cloud), `https://api.wallarm.com` for the [EU cloud](https://docs.wallarm.com/quickstart-en/how-wallarm-works/qs-intro-en/#eu-cloud). This can also be specified with the `WALLARM_API_HOST` shell environment variable. Default: `https://api.wallarm.com`.
* `client_id` - (Optional) ID of the client to apply the rules to. The value is required for multi-tenant scenarios. This can also be specified with the `WALLARM_API_CLIENT_ID` shell environment variable. Default: client ID of the authenticated user defined by UUID and SECRET.
* `retries` - (Optional) Maximum number of retries to perform when an API request fails. Default: 3. This can also be specified with the `WALLARM_API_RETRIES` shell environment variable.
* `min_backoff` - (Optional) Minimum backoff period in seconds after failed API calls. Default: 1. This can also be specified with the `WALLARM_API_MIN_BACKOFF` shell environment variable.
* `max_backoff` - (Optional) Maximum backoff period in seconds after failed API calls Default: 30. This can also be specified with the `WALLARM_API_MAX_BACKOFF` shell environment variable.
* `api_client_logging` - (Optional) Whether to print logs from the API client (using the default log library logger). Default: false. This can also be specified with the `WALLARM_API_CLIENT_LOGGING` shell environment variable.

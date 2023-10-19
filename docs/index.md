---
organization: francois2metz
category: ["saas"]
brand_color: "#03eab3"
display_name: "Phrase"
short_name: "phrase"
description: "Steampipe plugin for querying projects, translations and more from Phrase."
og_description: "Query phrase with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/francois2metz/phrase-social-graphic.png"
icon_url: "/images/plugins/francois2metz/phrase.svg"
---

# Phrase + Steampipe

[Phrase](https://phrase.com/) is a translation tool.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  name,
  created_at
from
  phrase_project
```

```
+--------------------+-----------------------+
| name               | created_at            |
+--------------------+-----------------------+
| My Awesome Project | 2023-09-04T10:42:02Z  |
| Another Project    | 2023-10-16T07:31:34Z  |
+--------------------+-----------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/phrase/tables)**

## Get started

### Install

Download and install the latest Phrase plugin:

```bash
steampipe plugin install francois2metz/phrase
```

### Credentials

| Item        | Description                                                                                                                                                                      |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Phrase requires an [access token](https://app.phrase.com/settings/oauth_access_tokens).                                                                                          |
| Permissions | Token should have at least the `read` scope. To manage members, add `team.manage`                                                                                                |
| Radius      | Each connection represents a single Phrase account.                                                                                                                              |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/phrase.spc`)<br />2. Credentials specified in environment variables, e.g., `PHRASE_ACCESS_TOKEN`. |

### Configuration

Installing the latest Phrase plugin will create a config file (`~/.steampipe/config/phrase.spc`) with a single connection named `phrase`:

```hcl
connection "phrase" {
    plugin = "francois2metz/phrase"

    # The access token of your phrase account (required)
    # Get it on: https://app.phrase.com/settings/oauth_access_tokens
    # This can also be set via the `PHRASE_ACCESS_TOKEN` environment variable
    # access_token = "siFreKnogsObVeirlildOakmygDectItEejOathdajtivSeawdyearArEsfosDys"

    # Datacenter (optional)
    # By default EU datacenter: https://api.phrase.com/v2
    # US datacenter: https://api.us.app.phrase.com/v2/
    # datacenter = "https://api.phrase.com/v2"
}
```

### Credentials from Environment Variables

The Phrase plugin will use the following environment variables to obtain credentials **only if other argument (`access_token`) is not specified** in the connection:

```sh
export PHRASE_ACCESS_TOKEN=siFreKnogsObVeirlildOakmygDectItEejOathdajtivSeawdyearArEsfosDys
```

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-phrase

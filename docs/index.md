---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/reddit.svg"
brand_color: "#FF4500"
display_name: Reddit
name: reddit
description: Steampipe plugin to query Reddit users, posts, votes and more.
og_description: Query Reddit with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/reddit-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Reddit + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Reddit](https://reddit.com) is an American social news aggregation, web content rating, and discussion website.

Example query:
```sql
select
  subreddit_name_prefixed,
  name,
  title,
  score,
  upvote_ratio
from
  reddit_my_post
order by
  rank;
```

```
+-------------------------+-----------+---------------------------------------+-------+--------------------+
| subreddit_name_prefixed | name      | title                                 | score | upvote_ratio       |
+-------------------------+-----------+---------------------------------------+-------+--------------------+
| r/netsec                | t3_q7llj8 | Open source automated NIST SP 800-53… | 146   | 0.949999988079071  |
| r/aws                   | t3_q1nwoz | What's your oldest S3 bucket?         | 77    | 0.949999988079071  |
| r/netsec                | t3_ot1kwy | Use SQL to query AbuseIPDB deny list… | 55    | 0.8899999856948853 |
| r/gitlab                | t3_ot0tzw | New open source tool to query GitLab… | 31    | 0.9700000286102295 |
| r/blueteamsec           | t3_q7lo6l | Open source automated NIST SP 800-53… | 25    | 0.9200000166893005 |
+-------------------------+-----------+---------------------------------------+-------+--------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/reddit/tables)**

## Get started

### Install

Download and install the latest Reddit plugin:

```bash
steampipe plugin install reddit
```

### Configuration

Installing the latest reddit plugin will create a config file (`~/.steampipe/config/reddit.spc`) with a single connection named `reddit`:

```hcl
connection "reddit" {
  plugin        = "reddit"

  client_id     = "aoxJBaKh9W_wnNLKzJhxSw"
  client_secret = "P51-fNlprSkGcqdkQzogJ_noqcktis"
  username      = "jane"
  password      = "Pa$$w0rd"
}
```

[Create a Reddit app](https://github.com/reddit-archive/reddit/wiki/OAuth2-Quick-Start-Example#first-steps) to get your settings:
* `client_id` - Client ID for your Reddit app.
* `client_secret` - Client secret for your Reddit app.
* `username` - Your username.
* `password` - Your password.

Environment variables are also available as an alternate configuration method:
* `REDDIT_CLIENT_ID`
* `REDDIT_CLIENT_SECRET`
* `REDDIT_USERNAME`
* `REDDIT_PASSWORD`

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-reddit
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)

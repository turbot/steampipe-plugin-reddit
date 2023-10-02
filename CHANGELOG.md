## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#21](https://github.com/turbot/steampipe-plugin-reddit/pull/21))
- Recompiled plugin with Go version `1.21`. ([#21](https://github.com/turbot/steampipe-plugin-reddit/pull/21))

## v0.2.1 [2023-05-19]

_Bug fixes_

- Added the missing `plugin` config argument to the example connection config in `docs/index.md` file. ([#17](https://github.com/turbot/steampipe-plugin-reddit/pull/17))

## v0.2.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#14](https://github.com/turbot/steampipe-plugin-reddit/pull/14))

## v0.1.0 [2022-09-17]

_What's new?_

- Added support for temporary credentials to query tables. ([#11](https://github.com/turbot/steampipe-plugin-reddit/pull/11))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#8](https://github.com/turbot/steampipe-plugin-reddit/pull/8))
- Recompiled plugin with Go version `1.19`. ([#8](https://github.com/turbot/steampipe-plugin-reddit/pull/8))

## v0.0.1 [2022-06-09]

_What's new?_

- New tables added
  - [reddit_my_blocked](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_blocked)
  - [reddit_my_comment](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_comment)
  - [reddit_my_friend](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_blocked)
  - [reddit_my_info](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_info)
  - [reddit_my_moderated_subreddit](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_moderated_subreddit)
  - [reddit_my_post](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_post)
  - [reddit_my_subscribed_subreddit](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_subscribed_subreddit)
  - [reddit_popular_subreddit](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_popular_subreddit)
  - [reddit_subreddit_banned](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_subreddit_banned)
  - [reddit_subreddit_moderator](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_subreddit_moderator)
  - [reddit_subreddit_post_search](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_subreddit_post_search)
  - [reddit_subreddit_search](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_subreddit_search)
  - [reddit_user](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_user)
  - [reddit_user_search](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_user_search)

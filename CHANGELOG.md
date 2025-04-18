## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#48](https://github.com/turbot/steampipe-plugin-reddit/pull/48))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#48](https://github.com/turbot/steampipe-plugin-reddit/pull/48))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#45](https://github.com/turbot/steampipe-plugin-reddit/pull/45))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#45](https://github.com/turbot/steampipe-plugin-reddit/pull/45))

## v0.5.0 [2024-01-08]

_What's new?_

- New tables added
  - [reddit_my_saved_post](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_saved_post) (Thanks [@mkell43](https://github.com/mkell43) for the contribution!)
  - [reddit_my_saved_comment](https://hub.steampipe.io/plugins/turbot/reddit/tables/reddit_my_saved_comment) (Thanks [@mkell43](https://github.com/mkell43) for the contribution!)

## v0.4.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#33](https://github.com/turbot/steampipe-plugin-reddit/pull/33))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#33](https://github.com/turbot/steampipe-plugin-reddit/pull/33))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-reddit/blob/main/docs/LICENSE). ([#33](https://github.com/turbot/steampipe-plugin-reddit/pull/33))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#32](https://github.com/turbot/steampipe-plugin-reddit/pull/32))

## v0.3.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#23](https://github.com/turbot/steampipe-plugin-reddit/pull/23))

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

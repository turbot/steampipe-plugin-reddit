---
title: "Steampipe Table: reddit_subreddit_banned - Query Reddit Banned Subreddits using SQL"
description: "Allows users to query banned subreddits on Reddit, offering insights into moderation actions and community guidelines enforcement."
---

# Table: reddit_subreddit_banned - Query Reddit Banned Subreddits using SQL

Reddit is a social news aggregation, web content rating, and discussion website. It allows members to submit content to the site such as links, text posts, and images, which are then voted up or down by other members. A key component of Reddit's moderation system involves the banning of subreddits that violate the site's community guidelines.

## Table Usage Guide

The `reddit_subreddit_banned` table provides insights into banned subreddits on Reddit. As a moderator, researcher, or community manager, use this table to understand the enforcement of Reddit's community guidelines and moderation actions. It can be used to analyze patterns in banned content, identify common reasons for subreddit bans, and inform content moderation strategies.

**Important Notes**
- You must specify the `subreddit` in the `where` clause to query this table.
- Requires permission to read banned users from the subreddit.

## Examples

### List users banned from the subreddit "mysubreddit"
Determine the individuals who have been banned from a specific subreddit, along with the date of their ban, remaining ban duration, and any notes associated with the ban. This can be useful for subreddit moderators to track and manage user activity and behavior.

```sql+postgres
select
  name,
  date,
  days_left,
  note
from
  reddit_subreddit_banned
where
  subreddit = 'mysubreddit'
order by
  name;
```

```sql+sqlite
select
  name,
  date,
  days_left,
  note
from
  reddit_subreddit_banned
where
  subreddit = 'mysubreddit'
order by
  name;
```
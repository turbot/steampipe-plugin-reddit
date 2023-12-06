---
title: "Steampipe Table: reddit_subreddit_moderator - Query Reddit Subreddit Moderators using SQL"
description: "Allows users to query Reddit Subreddit Moderators, providing insights into the moderators of different subreddits and their permissions."
---

# Table: reddit_subreddit_moderator - Query Reddit Subreddit Moderators using SQL

Reddit is a network of communities based on people's interests. Subreddits are the communities within Reddit, each of which is managed by a team of moderators. These moderators are responsible for enforcing community rules, removing content that violates these rules, and generally maintaining the health of their community.

## Table Usage Guide

The `reddit_subreddit_moderator` table provides insights into the moderators of different subreddits on Reddit. As a community manager or researcher, explore moderator-specific details through this table, including their permissions, subreddit affiliations, and associated metadata. Utilize it to understand the distribution of moderation responsibilities, identify active and inactive moderators, and analyze the moderation structure of different subreddits.

**Important Notes**
- You must specify the `subreddit` in the `where` clause to query this table.

## Examples

### List moderators from the subreddit "aws"
Explore which users have moderation permissions in a specific subreddit to understand the distribution of moderation roles and responsibilities.

```sql
select
  name,
  mod_permissions
from
  reddit_subreddit_moderator
where
  subreddit = 'aws'
order by
  name;
```

### List moderators with "all" permission
Explore which moderators have comprehensive permissions within a specific online community, such as 'aws', to better manage user roles and access rights. This can be useful for assessing user privileges and maintaining security protocols.

```sql
select
  name,
  mod_permissions
from
  reddit_subreddit_moderator
where
  subreddit = 'aws'
  and mod_permissions ? 'all'
order by
  name;
```
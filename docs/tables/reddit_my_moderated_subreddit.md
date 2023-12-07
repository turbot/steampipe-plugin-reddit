---
title: "Steampipe Table: reddit_my_moderated_subreddit - Query Reddit Moderated Subreddits using SQL"
description: "Allows users to query Moderated Subreddits in Reddit, specifically the details of subreddits that the authenticated user moderates, providing insights into subreddit management and moderation activities."
---

# Table: reddit_my_moderated_subreddit - Query Reddit Moderated Subreddits using SQL

Reddit is a social news aggregation, web content rating, and discussion website. It allows registered members to submit content to the site such as links, text posts, and images, which are then voted up or down by other members. A moderated subreddit is a specific section of the site with its own set of rules, guidelines, and moderators.

## Table Usage Guide

The `reddit_my_moderated_subreddit` table provides insights into moderated subreddits within Reddit. As a subreddit moderator, explore subreddit-specific details through this table, including rules, guidelines, and moderation activities. Utilize it to uncover information about subreddits, such as those with specific rules, the relationships between moderators, and the verification of guidelines.

## Examples

### List moderated subreddits
Discover the segments that have been moderated on Reddit, sorted by their display name. This can help you assess the popularity of these segments based on the number of subscribers, and easily access them via their URLs.

```sql+postgres
select
  display_name_prefixed,
  subscribers,
  urle
from
  reddit_my_moderated_subreddit
order by
  display_name_prefixed;
```

```sql+sqlite
select
  display_name_prefixed,
  subscribers,
  urle
from
  reddit_my_moderated_subreddit
order by
  display_name_prefixed;
```

### Top 5 moderated subreddits by popularity
Discover the five most popular subreddits you moderate, ranked by the number of subscribers. This information can be useful to determine where your moderation efforts are having the most impact.

```sql+postgres
select
  display_name_prefixed,
  subscribers,
  url
from
  reddit_my_moderated_subreddit
order by
  subscribers desc
limit 5;
```

```sql+sqlite
select
  display_name_prefixed,
  subscribers,
  url
from
  reddit_my_moderated_subreddit
order by
  subscribers desc
limit 5;
```
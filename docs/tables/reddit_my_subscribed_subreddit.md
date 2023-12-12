---
title: "Steampipe Table: reddit_my_subscribed_subreddit - Query Reddit Subscribed Subreddits using SQL"
description: "Allows users to query Subscribed Subreddits in Reddit, specifically providing information about the subreddits a user is subscribed to."
---

# Table: reddit_my_subscribed_subreddit - Query Reddit Subscribed Subreddits using SQL

Reddit is a network of communities where people can dive into their interests, hobbies, and passions. It allows users to subscribe to various subreddits, which are essentially topic-based communities. Each subreddit is dedicated to a specific topic, and users can participate in these communities by posting, commenting, and voting on content.

## Table Usage Guide

The `reddit_my_subscribed_subreddit` table provides insights into the subreddits that a user is subscribed to on Reddit. As a Reddit user or a social media analyst, explore subreddit-specific details through this table, including subreddit names, descriptions, and subscriber counts. Utilize it to uncover information about your subscriptions, such as identifying popular subreddits, analyzing the type of content in different subreddits, and understanding the demographics of subscribers.

## Examples

### List subscribed subreddits
Explore your Reddit subscriptions to understand the popularity and reach of your subscribed subreddits. This can help identify which communities are largest and most active, providing insight into where your interests align with larger Reddit trends.

```sql+postgres
select
  display_name_prefixed,
  subscribers,
  url
from
  reddit_my_subscribed_subreddit
order by
  display_name_prefixed;
```

```sql+sqlite
select
  display_name_prefixed,
  subscribers,
  url
from
  reddit_my_subscribed_subreddit
order by
  display_name_prefixed;
```

### Top 5 subscribed subreddits by popularity
Discover the segments that are most popular among your subscribed subreddits. This helps in identifying the top 5 subreddits you are subscribed to, based on the number of subscribers, for a better understanding of popular trends.

```sql+postgres
select
  display_name_prefixed,
  subscribers,
  url
from
  reddit_my_subscribed_subreddit
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
  reddit_my_subscribed_subreddit
order by
  subscribers desc
limit 5;
```
---
title: "Steampipe Table: reddit_popular_subreddit - Query Reddit Popular Subreddits using SQL"
description: "Allows users to query Popular Subreddits on Reddit, providing insights into the most trending and active subreddit communities."
---

# Table: reddit_popular_subreddit - Query Reddit Popular Subreddits using SQL

Popular Subreddits on Reddit are the most active and trending communities that are categorized based on various topics of interest. These subreddits have a large number of active users and posts, making them a rich source of discussions and information. They cover a wide range of topics, from news and current events to specific interests like technology, movies, music, and more.

## Table Usage Guide

The `reddit_popular_subreddit` table provides insights into the most popular subreddits on Reddit. As a data analyst or social media marketer, you can explore subreddit-specific details through this table, including the number of active users, post frequency, and community engagement. Utilize it to uncover information about trending topics, user behavior, and the popularity of various interests, which can be valuable for market research and trend analysis.

**Important Notes**
- Limited to the top 500 results.

## Examples

### Top subreddits by popularity
Discover the most popular subreddits based on their subscriber count, which allows you to understand the trending topics and areas of interest among Reddit users.

```sql
select
  rank,
  display_name_prefixed,
  title,
  subscribers
from
  reddit_popular_subreddit
order by
  rank;
```
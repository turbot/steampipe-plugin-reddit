---
title: "Steampipe Table: reddit_subreddit_post_search - Query Reddit Subreddit Posts using SQL"
description: "Allows users to query Subreddit Posts in Reddit, specifically providing insights into the posts' content, comments, and metadata."
---

# Table: reddit_subreddit_post_search - Query Reddit Subreddit Posts using SQL

Reddit is a social news aggregation, web content rating, and discussion platform. Registered members submit content to the site such as links, text posts, and images, which are then voted up or down by other members. Posts are organized by subject into user-created boards called "subreddits", which cover a variety of topics.

## Table Usage Guide

The `reddit_subreddit_post_search` table provides insights into Subreddit Posts within Reddit. As a data analyst, explore post-specific details through this table, including content, comments, and associated metadata. Utilize it to uncover information about posts, such as trending topics, popular comments, and user engagement levels.

## Examples

### Search the "aws" subreddit for "steampipe"
Explore the popularity and relevance of Steampipe within the AWS community on Reddit. This query helps identify posts that mention 'steampipe', allowing users to understand its significance and usage trends among AWS users.

```sql
select
  rank,
  title,
  created_utc,
  score,
  url
from
  reddit_subreddit_post_search
where
  subreddit = 'aws'
  and query = 'steampipe'
order by
  rank;
```
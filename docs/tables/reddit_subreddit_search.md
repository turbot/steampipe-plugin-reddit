---
title: "Steampipe Table: reddit_subreddit_search - Query Reddit Subreddit Search using SQL"
description: "Allows users to query Subreddit Search in Reddit, providing insights into the search results of subreddits based on a specified query."
---

# Table: reddit_subreddit_search - Query Reddit Subreddit Search using SQL

Reddit Subreddit Search is a feature within Reddit that allows you to search for subreddits based on a specified query. It provides a way to discover new communities and topics within the Reddit platform. Subreddit Search helps you stay informed about the trending topics and discussions in your areas of interest.

## Table Usage Guide

The `reddit_subreddit_search` table provides insights into Subreddit Search within Reddit. As a data analyst, explore subreddit-specific details through this table, including subreddit names, descriptions, and associated metadata. Utilize it to uncover information about subreddits, such as those with a high number of subscribers, the most discussed topics, and the overall popularity of a subreddit.

**Important Notes**
- You must specify the `query` in the `where` clause to query this table.

## Examples

### Search subreddits with query "aws"
Explore various subreddits related to 'aws' to understand their popularity based on subscriber count, allowing you to identify the most engaged communities for specific topics.

```sql+postgres
select
  rank,
  display_name_prefixed,
  title,
  subscribers
from
  reddit_subreddit_search
where
  query = 'aws'
order by
  rank;
```

```sql+sqlite
select
  rank,
  display_name_prefixed,
  title,
  subscribers
from
  reddit_subreddit_search
where
  query = 'aws'
order by
  rank;
```
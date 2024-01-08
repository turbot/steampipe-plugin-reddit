---
title: "Table: reddit_my_saved_comment - Query Reddit Saved Comments using SQL"
description: "Allows users to query saved comments in Reddit, specifically user's saved comments, providing insights into user interactions and preferences."
---

# Table: reddit_my_saved_comment - Query Reddit Saved Comments using SQL

Reddit is a social news aggregation, web content rating, and discussion website. Registered members submit content to the site such as links, text posts, and images, which are then voted up or down by other members. The saved comments feature allows users to save specific comments for later reference.

## Table Usage Guide

The `reddit_my_saved_comment` table provides insights into user's saved comments within Reddit. As a data analyst, explore user-specific details through this table, including the content of the saved comments, the author of the comments, and the subreddit in which the comments were made. Utilize it to uncover information about user behavior and preferences, such as the topics they are interested in and the discussions they engage in. The schema presents a range of attributes of the saved comments for your analysis, like the comment body, creation date, author, and associated subreddit.

## Examples

### List five most recent comments
Explore the most recent discussions you've saved on Reddit. This query is useful for quickly accessing and reviewing your latest interactions, without having to manually sift through your comment history.

```sql+postgres
select
  created_utc,
  permalink,
  body
from
  reddit_my_saved_comment
order by
  created_utc desc
limit 5;
```

```sql+sqlite
select
  created_utc,
  permalink,
  body
from
  reddit_my_saved_comment
order by
  created_utc desc
limit 5;
```

### List top five comments by score
Discover the highest-rated comments from your saved Reddit comments. This can help you quickly identify popular opinions or trending topics within your saved content.

```sql+postgres
select
  score,
  permalink,
  body,
  replies
from
  reddit_my_saved_comment
order by
  score desc
limit 5;
```

```sql+sqlite
select
  score,
  permalink,
  body,
  replies
from
  reddit_my_saved_comment
order by
  score desc
limit 5;
```

### List comments by subreddit
Explore which subreddits your saved comments are most frequently associated with to understand your engagement patterns across different communities.

```sql+postgres
select
  subreddit_name_prefixed,
  count(*)
from
  reddit_my_saved_comment
group by
  subreddit_name_prefixed
order by
  count desc;
```

```sql+sqlite
select
  subreddit_name_prefixed,
  count(*)
from
  reddit_my_saved_comment
group by
  subreddit_name_prefixed
order by
  count(*) desc;
```

### List comments that contain the word "docs"
Explore comments that include the term 'docs'. This can be useful to identify discussions or references related to documentation, aiding in information gathering and community engagement.

```sql+postgres
select
  created_utc,
  permalink,
  body
from
  reddit_my_saved_comment
where
  body ilike '%docs%'
order by
  created_utc;
```

```sql+sqlite
select
  created_utc,
  permalink,
  body
from
  reddit_my_saved_comment
where
  body like '%docs%'
order by
  created_utc;
```
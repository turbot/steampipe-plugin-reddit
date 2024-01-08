---
title: "Table: reddit_my_saved_post - Query Reddit Saved Posts using SQL"
description: "Allows users to query their saved posts on Reddit, providing insights into post details such as title, author, subreddit, and more."
---

# Table: reddit_my_saved_post - Query Reddit Saved Posts using SQL

Reddit is a social media platform that allows users to discuss and vote on content shared by other users. Users can save posts for later reference; these saved posts can be from any subreddit and include various details such as the post's title, author, subreddit, and more. This functionality is part of the overall user interaction with the platform, contributing to the personalized user experience.

## Table Usage Guide

The `reddit_my_saved_post` table provides insights into a user's saved posts within Reddit. As a data analyst, explore post-specific details through this table, including title, author, subreddit, and associated metadata. Utilize it to uncover information about saved posts, such as those from specific subreddits, posts by certain authors, and the nature of the content saved by the user. The schema presents a range of attributes of the saved post for your analysis, like the post ID, title, author, subreddit, and more.

## Examples

### List five most recent posts
Explore the most recent activities on your Reddit account by identifying the five latest saved posts. This can help you keep track of your recent interactions and interests.

```sql+postgres
select
  created_utc,
  title,
  url
from
  reddit_my_saved_post
order by
  created_utc desc
limit 5;
```

```sql+sqlite
select
  created_utc,
  title,
  url
from
  reddit_my_saved_post
order by
  created_utc desc
limit 5;
```

### List top five posts by score
Gain insights into the most popular posts based on their score. This query helps you identify the top five posts, offering a quick overview of the most engaging content.

```sql+postgres
select
  score,
  upvote_ratio,
  title,
  url
from
  reddit_my_saved_post
order by
  score desc
limit 5;
```

```sql+sqlite
select
  score,
  upvote_ratio,
  title,
  url
from
  reddit_my_saved_post
order by
  score desc
limit 5;
```

### List posts by subreddit
Discover the segments that garner the most engagement on your saved Reddit posts. This allows you to focus your attention on the most active subreddits, thus optimizing your Reddit usage.

```sql+postgres
select
  subreddit_name_prefixed,
  count(*)
from
  reddit_my_saved_post
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
  reddit_my_saved_post
group by
  subreddit_name_prefixed
order by
  count(*) desc;
```

### List posts that contain the word "docs"
Discover the segments that include references to "docs" in your saved posts on Reddit. This can help you quickly locate posts that mention documentation or similar topics.

```sql+postgres
select
  created_utc,
  title,
  url,
  selftext
from
  reddit_my_saved_post
where
  selftext ilike '%docs%'
order by
  created_utc;
```

```sql+sqlite
select
  created_utc,
  title,
  url,
  selftext
from
  reddit_my_saved_post
where
  selftext like '%docs%'
order by
  created_utc;
```
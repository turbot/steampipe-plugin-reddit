---
title: "Steampipe Table: reddit_my_post - Query Reddit Posts using SQL"
description: "Allows users to query their Reddit Posts, specifically to retrieve details such as post title, content, subreddit, upvotes, and more, providing insights into user's post activity and engagement."
---

# Table: reddit_my_post - Query Reddit Posts using SQL

Reddit is a network of communities based on people's interests. It allows users to post content, including text posts, links, and images, which are then voted up or down by other members. Posts are organized by subject into user-created boards called "subreddits", which cover a variety of topics like news, science, movies, video games, music, books, fitness, food, and image-sharing.

## Table Usage Guide

The `reddit_my_post` table provides insights into a user's posts on Reddit. As a data analyst or social media manager, explore post-specific details through this table, including post title, content, subreddit, upvotes, and more. Utilize it to uncover information about user's post activity, engagement, and the popularity of posts across different subreddits.

## Examples

### 5 most recent posts
Discover the latest updates or additions to your Reddit posts. This query can be used to keep track of your most recent posts, helping you to maintain an active and timely presence on the platform.

```sql
select
  created_utc,
  title,
  url
from
  reddit_my_post
order by
  created_utc desc
limit 5;
```

### Top 5 posts by score
Discover the most popular posts based on their scores to understand what content resonates most with your audience. This can help you tailor future content to increase engagement and upvotes.

```sql
select
  score,
  upvote_ratio,
  title,
  url
from
  reddit_my_post
order by
  score desc
limit 5;
```

### Posts by subreddit
Explore which subreddits you are most active in by counting the number of posts you have made in each. This can help you understand your Reddit usage patterns and areas of interest.

```sql
select
  subreddit_name_prefixed,
  count(*)
from
  reddit_my_post
group by
  subreddit_name_prefixed
order by
  count desc;
```

### Posts containing the word "docs"
Discover the segments that include the word 'docs' within your Reddit posts. This can help you analyze the frequency and context of discussions about documentation, providing insights into user engagement and potential areas of improvement.

```sql
select
  created_utc,
  title,
  url,
  selftext
from
  reddit_my_post
where
  selftext ilike '%docs%'
order by
  created_utc;
```
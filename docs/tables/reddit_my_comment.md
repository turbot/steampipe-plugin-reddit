---
title: "Steampipe Table: reddit_my_comment - Query Reddit My Comments using SQL"
description: "Allows users to query their own comments on Reddit, specifically the text, score, and associated metadata, providing insights into user activity and engagement."
---

# Table: reddit_my_comment - Query Reddit My Comments using SQL

Reddit is a vast network of communities that are created, run, and populated by its users. Each user has the ability to comment on posts, providing their own insights, information, or humor. These comments become a part of the user's activity history and can be queried for analysis or review.

## Table Usage Guide

The `reddit_my_comment` table provides insights into a user's own comments within Reddit. As a data analyst, explore comment-specific details through this table, including the text, score, and associated metadata. Utilize it to uncover information about your comments, such as those with high engagement, the sentiment of your comments, and the overall activity pattern.

## Examples

### 5 most recent comments
Gain insights into your recent activity on Reddit by identifying the five most recent comments you've made. This can help you track your interactions and discussions on the platform.

```sql+postgres
select
  created_utc,
  permalink,
  body
from
  reddit_my_comment
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
  reddit_my_comment
order by
  created_utc desc
limit 5;
```

### Top 5 comments by score
Explore the most popular comments on your Reddit account, ranked by score. This can provide insight into which comments resonated most with other users, potentially informing future engagement strategies.

```sql+postgres
select
  score,
  permalink,
  body,
  replies
from
  reddit_my_comment
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
  reddit_my_comment
order by
  score desc
limit 5;
```

### Comments by subreddit
Explore which subreddits have the most user comments to understand where the most active discussions are taking place. This can help identify popular topics and trends within specific online communities.

```sql+postgres
select
  subreddit_name_prefixed,
  count(*)
from
  reddit_my_comment
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
  reddit_my_comment
group by
  subreddit_name_prefixed
order by
  count(*) desc;
```

### Comments containing the word "docs"
Explore comments that include the term 'docs', organized by the date they were created. This can help identify discussions or questions related to documentation, potentially highlighting areas needing improvement or clarification.

```sql+postgres
select
  created_utc,
  permalink,
  body
from
  reddit_my_comment
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
  reddit_my_comment
where
  body like '%docs%'
order by
  created_utc;
```
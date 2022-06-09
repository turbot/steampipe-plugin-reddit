# Table: reddit_my_post

Query your own posts.

## Examples

### 5 most recent posts

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

# Table: reddit_my_comment

Query your own comments.

## Examples

### 5 most recent comments

```sql
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

```sql
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

```sql
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

### Comments containing the word "docs"

```sql
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

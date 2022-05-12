# Table: reddit_my_subscribed_subreddit

Subreddits you are subscribed to.

## Examples

### List subscribed subreddits

```sql
select
  display_name_prefixed,
  subscribers,
  url
from
  reddit_my_subscribed_subreddit
order by
  display_name_prefixed
```

### Top 5 subscribed subreddits by popularity

```sql
select
  display_name_prefixed,
  subscribers,
  url
from
  reddit_my_subscribed_subreddit
order by
  subscribers desc
limit 5
```

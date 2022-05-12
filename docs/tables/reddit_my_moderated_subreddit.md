# Table: reddit_my_moderated_subreddit

Subreddits you are a moderator of.

## Examples

### List moderated subreddits

```sql
select
  display_name_prefixed,
  subscribers,
  urle
from
  reddit_my_moderated_subreddit
order by
  display_name_prefixed
```

### Top 5 moderated subreddits by popularity

```sql
select
  display_name_prefixed,
  subscribers,
  url
from
  reddit_my_moderated_subreddit
order by
  subscribers desc
limit 5
```

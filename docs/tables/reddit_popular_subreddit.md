# Table: reddit_popular_subreddit

List the most popular subreddits.

Notes:
* Limited to the top 500 results.

## Examples

### Top subreddits by popularity

```sql
select
  rank,
  display_name_prefixed,
  title,
  subscribers
from
  reddit_popular_subreddit
order by
  rank
```

# Table: reddit_subreddit_post_search

Search for posts within a subreddit that match the `query`.

Notes:
* `subreddit` and `query` must both be specified in the where clause.

## Examples

### Search the "aws" subreddit for "steampipe"

```sql
select
  rank,
  title,
  created_utc,
  score,
  url
from
  reddit_subreddit_post_search
where
  subreddit = 'aws'
  and query = 'steampipe'
order by
  rank;
```

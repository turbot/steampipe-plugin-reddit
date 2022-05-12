# Table: reddit_subreddit_search

Search for subreddits that match the `query`.

Notes:
* `query` must be specified in the where clause.

## Examples

### Search subreddits with query "aws"

```sql
select
  rank,
  display_name_prefixed,
  title,
  subscribers
from
  reddit_subreddit_search
where
  query = 'aws'
order by
  rank
```

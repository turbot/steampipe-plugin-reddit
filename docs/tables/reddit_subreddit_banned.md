# Table: reddit_subreddit_banned

Users banned from a given subreddit.

Notes:
* `subreddit` must be specified in the where clause.
* Requires permission to read banned users from the subreddit.

## Examples

### List users banned from the subreddit "mysubreddit"

```sql
select
  name,
  date,
  days_left,
  note
from
  reddit_subreddit_banned
where
  subreddit = 'mysubreddit'
order by
  name;
```

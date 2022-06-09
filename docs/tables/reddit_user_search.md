# Table: reddit_user_search

Search Reddit users.

Notes:
* `query` must be specified in the where clause.

## Examples

### Search users called "jane"

```sql
select
  *
from
  reddit_user_search
where
  query = 'jane'
order by
  rank;
```

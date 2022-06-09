# Table: reddit_user_search

Search Reddit users using `query`.

Notes:
* `query` must be specified in the where clause of queries.

## Examples

### Search for the user "steampipeio"

```sql
select
  *
from
  reddit_user_search
where
  query = 'steampipeio';
```

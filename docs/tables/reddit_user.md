# Table: reddit_user

Get information about a user.

Notes:
* `name` must be specified in the where clause of queries.

## Examples

### Get a user

```sql
select
  *
from
  reddit_user
where
  name = 'steampipeio';
```

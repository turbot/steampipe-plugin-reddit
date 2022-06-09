# Table: reddit_my_blocked

Query users you have blocked.

## Examples

### List all users you have blocked

```sql
select
  name,
  date
from
  reddit_my_blocked
order by
  name;
```

### Last 5 users you blocked

```sql
select
  name,
  date
from
  reddit_my_blocked
order by
  date desc
limit 5;
```

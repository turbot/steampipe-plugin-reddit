# Table: reddit_subreddit_moderator

Moderators for the given subreddit.

Notes:
* `subreddit` must be specified in the where clause.

## Examples

### List moderators from the subreddit "aws"

```sql
select
  name,
  mod_permissions
from
  reddit_subreddit_moderator
where
  subreddit = 'aws'
order by
  name
```

### List moderators with "all" permission

```sql
select
  name,
  mod_permissions
from
  reddit_subreddit_moderator
where
  subreddit = 'aws'
  and mod_permissions ? 'all'
order by
  name
```

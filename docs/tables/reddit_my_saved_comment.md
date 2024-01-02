# Table: reddit_my_saved_comment

Query your own saved comments.

Note:
When querying the `reddit_my_saved_post` or `reddit_my_saved_comment` tables, be aware that the underlying library fetches both saved posts and comments together from the Reddit API, with a default request of 100 items. Each query, however, will only return one type of item—either posts or comments—based on the table you are querying. Due to this behavior, the number of items returned in your query might be less than the total number of saved items you have, or the default request amount of 100, especially if there is a significant mix of saved posts and comments. The SQL `LIMIT` clause will only apply to the data returned by the library, not the data requested from the Reddit API.

## Examples

### 5 most recent comments

```sql
select
  created_utc,
  permalink,
  body
from
  reddit_my_saved_comment
order by
  created_utc desc
limit 5;
```

### Top 5 comments by score

```sql
select
  score,
  permalink,
  body,
  replies
from
  reddit_my_saved_comment
order by
  score desc
limit 5;
```

### Comments by subreddit

```sql
select
  subreddit_name_prefixed,
  count(*)
from
  reddit_my_saved_comment
group by
  subreddit_name_prefixed
order by
  count desc;
```

### Comments containing the word "docs"

```sql
select
  created_utc,
  permalink,
  body
from
  reddit_my_saved_comment
where
  body ilike '%docs%'
order by
  created_utc;
```

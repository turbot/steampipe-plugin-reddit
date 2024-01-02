# Table: reddit_my_saved_post

Query your saved posts.

Note:
When querying the `reddit_my_saved_post` or `reddit_my_saved_comment` tables, be aware that the underlying library fetches both saved posts and comments together from the Reddit API, with a default request of 100 items. Each query, however, will only return one type of item—either posts or comments—based on the table you are querying. Due to this behavior, the number of items returned in your query might be less than the total number of saved items you have, or the default request amount of 100, especially if there is a significant mix of saved posts and comments. The SQL `LIMIT` clause will only apply to the data returned by the library, not the data requested from the Reddit API.

## Examples

### List five most recent posts

```sql
select
  created_utc,
  title,
  url
from
  reddit_my_saved_post
order by
  created_utc desc
limit 5;
```

### List top five posts by score

```sql
select
  score,
  upvote_ratio,
  title,
  url
from
  reddit_my_saved_post
order by
  score desc
limit 5;
```

### List posts by subreddit

```sql
select
  subreddit_name_prefixed,
  count(*)
from
  reddit_my_saved_post
group by
  subreddit_name_prefixed
order by
  count desc;
```

### List posts that contain the word "docs"

```sql
select
  created_utc,
  title,
  url,
  selftext
from
  reddit_my_saved_post
where
  selftext ilike '%docs%'
order by
  created_utc;
```

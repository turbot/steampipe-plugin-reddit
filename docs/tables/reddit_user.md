---
title: "Steampipe Table: reddit_user - Query Reddit User using SQL"
description: "Allows users to query Reddit User profiles, providing insights into user details such as their id, name, and created timestamp."
---

# Table: reddit_user - Query Reddit User using SQL

Reddit is a social news aggregation, web content rating, and discussion website. Users register on the platform and then participate in various communities, known as 'subreddits'. These users can post content, comment, and vote on other user's posts.

## Table Usage Guide

The `reddit_user` table provides insights into user profiles within Reddit. As a data analyst, explore user-specific details through this table, including their id, name, and created timestamp. Utilize it to uncover information about users, such as their activity patterns, engagement levels, and overall contribution to the platform.

**Important Notes**
- You must specify the `name` in the `where` clause to query this table.

## Examples

### Get a user
Explore which Reddit user is associated with the username 'steampipeio'. This can help in identifying the user's activity and interactions on the platform.

```sql+postgres
select
  *
from
  reddit_user
where
  name = 'steampipeio';
```

```sql+sqlite
select
  *
from
  reddit_user
where
  name = 'steampipeio';
```
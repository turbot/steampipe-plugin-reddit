---
title: "Steampipe Table: reddit_my_info - Query Reddit User Information using SQL"
description: "Allows users to query User Information on Reddit, specifically the personal attributes and preferences, providing insights into user activity and behavior."
---

# Table: reddit_my_info - Query Reddit User Information using SQL

Reddit is a social news aggregation, web content rating, and discussion website. It allows registered members to submit content to the site such as links, text posts, and images, which are then voted up or down by other members. User information on Reddit includes personal attributes and preferences, providing insights into user activity and behavior.

## Table Usage Guide

The `reddit_my_info` table provides insights into user information on Reddit. As a data analyst or social media manager, explore user-specific details through this table, including personal attributes and preferences. Utilize it to uncover information about user activity and behavior, such as content submission, voting patterns, and interaction with other members.

## Examples

### Search for the user "steampipeio"
Explore which Reddit users have the username 'steampipeio'. This is useful for identifying specific users in the Reddit community for potential engagement or analysis.

```sql+postgres
select
  *
from
  reddit_user_search
where
  query = 'steampipeio';
```

```sql+sqlite
select
  *
from
  reddit_my_info;
```
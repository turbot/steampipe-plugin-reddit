---
title: "Steampipe Table: reddit_user_search - Query Reddit User Profiles using SQL"
description: "Allows users to query Reddit User Profiles, specifically the search results for users based on a query string, providing insights into user activities and preferences."
---

# Table: reddit_user_search - Query Reddit User Profiles using SQL

Reddit is a popular social media platform that allows users to discuss and vote on content that other users have submitted. It provides a platform for communities to discuss, connect, and share in an open environment, home to some of the most authentic content anywhere online. The nature of the platform allows users to remain anonymous and real-time, with the community deciding which discussions and topics get visibility.

## Table Usage Guide

The `reddit_user_search` table provides insights into user profiles within Reddit. As a data analyst or a social media marketer, explore user-specific details through this table, including their activities, preferences, and interactions. Utilize it to uncover information about users, such as their post history, comments, and upvotes, to understand user behavior and trends.

**Important Notes**
- You must specify the `query` in the `where` clause to query this table.

## Examples

### Search users called "jane"
Discover the segments that contain users named 'Jane' on Reddit, allowing you to analyze user data and trends associated with this name. This is particularly useful for market research, user analysis, and trend prediction.

```sql+postgres
select
  *
from
  reddit_user_search
where
  query = 'jane'
order by
  rank;
```

```sql+sqlite
select
  *
from
  reddit_user_search
where
  query = 'jane'
order by
  rank;
```
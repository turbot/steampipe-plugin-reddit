---
title: "Steampipe Table: reddit_my_blocked - Query Reddit Blocked Users using SQL"
description: "Allows users to query blocked users in Reddit, specifically the list of users that the authenticated user has blocked."
---

# Table: reddit_my_blocked - Query Reddit Blocked Users using SQL

Reddit Blocked Users is a feature within Reddit that allows users to block other users to prevent them from interacting with them on the platform. It provides a way to manage and maintain a list of users that the authenticated user does not want to interact with. Reddit Blocked Users helps to enhance the user experience by allowing users to control who can interact with them.

## Table Usage Guide

The `reddit_my_blocked` table provides insights into blocked users within Reddit. As a Reddit user, explore the list of blocked users through this table, including their usernames, the reasons for blocking, and the time they were blocked. Utilize it to manage and maintain your interactions on the platform, ensuring a more controlled and personalized Reddit experience.

## Examples

### List all users you have blocked
Discover the individuals you have chosen to block on Reddit and the dates these actions took place, allowing you to recall and manage your personal online interactions. This can be beneficial in maintaining a positive and controlled social media environment.

```sql+postgres
select
  name,
  date
from
  reddit_my_blocked
order by
  name;
```

```sql+sqlite
select
  name,
  date
from
  reddit_my_blocked
order by
  name;
```

### Last 5 users you blocked
Explore the most recent activities in your Reddit account by identifying the last five users you've blocked. This helps maintain a record of your interactions and monitor any recurring issues with specific users.

```sql+postgres
select
  name,
  date
from
  reddit_my_blocked
order by
  date desc
limit 5;
```

```sql+sqlite
select
  name,
  date
from
  reddit_my_blocked
order by
  date desc
limit 5;
```
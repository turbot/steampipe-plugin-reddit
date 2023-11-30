---
title: "Steampipe Table: reddit_my_friend - Query Reddit Friends using SQL"
description: "Allows users to query Friends on Reddit, specifically the details about the users who are friends with the authenticated user."
---

# Table: reddit_my_friend - Query Reddit Friends using SQL

Reddit is a social news aggregation, web content rating, and discussion website. Registered members submit content to the site such as links, text posts, and images, which are then voted up or down by other members. A unique feature of Reddit is the 'Friends' system, which allows users to follow and interact with other users on the platform.

## Table Usage Guide

The `reddit_my_friend` table provides insights into the 'Friends' feature on Reddit. As a social media analyst, explore details about the users who are friends with the authenticated user through this table, including their username, ID, and the date you became friends. Utilize it to uncover information about your Reddit network, such as who your oldest friend is, or if there are any patterns in the users you befriend.

## Examples

### List all your friends
Explore your social connections by listing all your friends in alphabetical order, useful for quickly identifying a specific friend or reviewing your social network.

```sql
select
  name,
  date
from
  reddit_my_friend
order by
  name;
```
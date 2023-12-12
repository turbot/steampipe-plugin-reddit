package reddit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-reddit",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"reddit_my_blocked":              tableRedditMyBlocked(ctx),
			"reddit_my_comment":              tableRedditMyComment(ctx),
			"reddit_my_friend":               tableRedditMyFriend(ctx),
			"reddit_my_info":                 tableRedditMyInfo(ctx),
			"reddit_my_post":                 tableRedditMyPost(ctx),
			"reddit_my_subscribed_subreddit": tableRedditMySubscribedSubreddit(ctx),
			"reddit_my_moderated_subreddit":  tableRedditMyModeratedSubreddit(ctx),
			"reddit_popular_subreddit":       tableRedditPopularSubreddit(ctx),
			"reddit_subreddit_banned":        tableRedditSubredditBanned(ctx),
			"reddit_subreddit_moderator":     tableRedditSubredditModerator(ctx),
			"reddit_subreddit_post_search":   tableRedditSubredditPostSearch(ctx),
			"reddit_subreddit_search":        tableRedditSubredditSearch(ctx),
			"reddit_user":                    tableRedditUser(ctx),
			"reddit_user_search":             tableRedditUserSearch(ctx),

			// my comment
			// my downvoted
			// my upvoted
			// my inbox (message)
			// my trophies
			// my settings
			// my subreddit moderated (ones I'm a moderator of)
			// moderator queue
			// moderator reported
			// moderator spam

			// user (get)
			// user downvoted (downvotedof)
			// user upvoted (upvotedof)
			// user comment (commentsof)
			// user post (postsof)
			// user trophy (trophiesof)
			// user hidden post
			// user search

			// post search https://pkg.go.dev/github.com/vartanbeno/go-reddit/v2@v2.0.1/reddit#SubredditService.SearchPosts

			// subreddit
			// subreddit_banned
			// subreddit_sticky
			// subreddit_contributor
			// subreddit_muted
			// subreddit_moderator
			// subreddit_post_requirement
			// subreddit_rule
			// subreddit_subscribed
			// subreddit_traffic
			// subreddit_settings
			// subreddit_controversial_post (how to limit / scope?)
			// subreddit_hot_post (how to limit / scope?)
			// subreddit_new_post (how to limit / scope?)
			// subreddit_top_post (how to limit / scope?)
			// subreddit_rising_post (how to limit / scope?)
		},
	}
	return p
}

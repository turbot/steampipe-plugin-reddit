package reddit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func tableRedditSubredditPostSearch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_subreddit_post_search",
		Description: "Search posts in a subreddit.",
		// Allow for 0 counts
		DefaultTransform: transform.FromJSONTag(),
		List: &plugin.ListConfig{
			Hydrate: listSubredditPostSearch,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "subreddit"},
				{Name: "query", CacheMatch: "exact"},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the post among the result rows, use for sorting."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.Title"), Description: "Title of the post."},
			{Name: "score", Type: proto.ColumnType_INT, Transform: transform.FromField("Post.Score"), Description: "Score of the post."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.URL"), Description: "URL the post links to, or of the post itself."},
			// Other columns
			{Name: "author", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.Author"), Description: "Author of the post."},
			{Name: "author_fullname", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.AuthorID"), Description: "Full name of the author for the post."},
			{Name: "created_utc", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Post.Created").Transform(timeToRfc3339), Description: "Time when the post was created."},
			{Name: "edited", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Post.Edited").Transform(timeToRfc3339), Description: "Time when the post was edited."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.ID"), Description: "ID of the post."},
			{Name: "is_self", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.IsSelfPost"), Description: ""},
			{Name: "likes", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Likes"), Description: "True if you've upvoted the post. False if you've downvoted it. Otherwise null."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Locked"), Description: "True if the post is locked."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.FullID"), Description: "Slug (full ID) of the post."},
			{Name: "num_comments", Type: proto.ColumnType_INT, Transform: transform.FromField("Post.NumberOfComments"), Description: "Number of comments on the post."},
			{Name: "over_18", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.NSFW"), Description: "True if the post is not safe for work (over 18)."},
			{Name: "permalink", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.Permalink"), Description: "Permalink (path only) to the post."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Search query string."},
			{Name: "saved", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Saved"), Description: "True if the post has been saved."},
			{Name: "selftext", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.Body"), Description: "Body of the post."},
			{Name: "spoiler", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Spoiler"), Description: "True if the post is a spoiler."},
			{Name: "stickied", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Stickied"), Description: "True if the post has been stickied."},
			{Name: "subreddit", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.SubredditName"), Description: "Name of the subreddit, e.g. aws."},
			{Name: "subreddit_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.SubredditID"), Description: "ID of the subreddit."},
			{Name: "subreddit_name_prefixed", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.SubredditNamePrefixed"), Description: "Prefixed name of the subreddit, e.g. /r/aws."},
			{Name: "subreddit_subscribers", Type: proto.ColumnType_INT, Transform: transform.FromField("Post.SubredditSubscribers"), Description: "Number of subscribers to the subreddit."},
			{Name: "upvote_ratio", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Post.UpvoteRatio"), Description: "Upvote ratio of the post."},
		}),
	}
}

type subredditPostSearchRow struct {
	Post *reddit.Post
	// Return rank to sort results
	Rank int `json:"rank"`
}

func listSubredditPostSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_subreddit_post_search.listSubredditPostSearch", "connection_error", err)
		return nil, err
	}

	subreddit := d.EqualsQuals["subreddit"].GetStringValue()
	query := d.EqualsQuals["query"].GetStringValue()

	opts := &reddit.ListPostSearchOptions{
		ListPostOptions: reddit.ListPostOptions{
			ListOptions: reddit.ListOptions{
				Limit: 100,
			},
		},
		Sort: "relevance",
	}

	count := 0
	for {
		posts, resp, err := conn.Subreddit.SearchPosts(ctx, query, subreddit, opts)
		if err != nil {
			plugin.Logger(ctx).Error("reddit_subreddit_post_search.listSubredditPostSearch", "query_error", err, "resp", resp, "query", query, "subreddit", subreddit)
			return nil, err
		}
		for _, i := range posts {
			count++
			row := subredditPostSearchRow{
				i,
				count,
			}
			d.StreamListItem(ctx, row)
		}
		if resp.After == "" {
			break
		}
		opts.After = resp.After
	}

	return nil, nil
}

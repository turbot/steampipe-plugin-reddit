package reddit

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableRedditMyComment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_my_comment",
		Description: "Your comments.",
		List: &plugin.ListConfig{
			Hydrate: listMyComment,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the comment among the result rows, use for sorting."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.ID"), Description: "ID of the comment."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.FullID"), Description: "Slug (full ID) of the comment."},
			{Name: "created_utc", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Comment.Created").Transform(timeToRfc3339), Description: "Time when the comment was created."},
			{Name: "edited", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Comment.Edited").Transform(timeToRfc3339), Description: "Time when the comment was edited."},
			{Name: "parent_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.ParentID"), Description: "Permalink (path only) to the comment."},
			{Name: "permalink", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.Permalink"), Description: "Permalink (path only) to the comment."},
			{Name: "body", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.Body"), Description: "Body of the comment."},
			{Name: "author", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.Author"), Description: "Author of the comment."},
			{Name: "author_fullname", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.AuthorID"), Description: "Full name of the author for the comment."},
			{Name: "author_flair_text", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.AuthorFlairText"), Description: "Flair text of the author for the comment."},
			{Name: "author_flair_template_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.AuthorFlairID"), Description: "ID of the flair template of the author for the comment."},
			{Name: "subreddit", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.SubredditName"), Description: "Name of the subreddit, e.g. aws."},
			{Name: "subreddit_name_prefixed", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.SubredditNamePrefixed"), Description: "Prefixed name of the subreddit, e.g. /r/aws."},
			{Name: "subreddit_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.SubredditID"), Description: "ID of the subreddit."},
			{Name: "likes", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Comment.Likes"), Description: "True if you've upvoted the comment. False if you've downvoted it. Otherwise null."},
			{Name: "score", Type: proto.ColumnType_INT, Transform: transform.FromField("Comment.Score"), Description: "Score of the comment."},
			{Name: "controversiality", Type: proto.ColumnType_INT, Transform: transform.FromField("Comment.Controversiality"), Description: "Controversiality score of the comment."},
			{Name: "link_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.PostID"), Description: ""},
			{Name: "link_title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.PostTitle"), Description: ""},
			{Name: "link_permalink", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.PostPermalink"), Description: ""},
			{Name: "link_author", Type: proto.ColumnType_STRING, Transform: transform.FromField("Comment.PostAuthor"), Description: ""},
			{Name: "num_comments", Type: proto.ColumnType_INT, Transform: transform.FromField("Comment.PostNumComments"), Description: ""},
			{Name: "is_submitter", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Comment.IsSubmitter"), Description: "True if the comment is a spoiler."},
			{Name: "score_hidden", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Comment.ScoreHidden"), Description: "True if the score is hidden on this comment."},
			{Name: "saved", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Comment.Saved"), Description: "True if the comment has been saved."},
			{Name: "stickied", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Comment.Stickied"), Description: "True if the comment has been stickied."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Comment.Locked"), Description: "True if the comment is locked."},
			{Name: "can_gild", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Comment.CanGild"), Description: ""},
			{Name: "over_18", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Comment.NSFW"), Description: "True if the comment is not safe for work (over 18)."},
			{Name: "replies", Type: proto.ColumnType_JSON, Transform: transform.FromField("Comment.Replies.Comments"), Description: "Replies to the comment."},
		},
	}
}

type myCommentRow struct {
	Comment *reddit.Comment
	// Return rank to sort results
	Rank int `json:"rank"`
}

func listMyComment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_comment.listMyComment", "connection_error", err)
		return nil, err
	}

	opts := &reddit.ListUserOverviewOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		// One of: hot, new, top, controversial.
		Sort: "top",
		// One of: hour, day, week, month, year, all.
		Time: "all",
	}

	count := 0
	for {
		items, resp, err := conn.User.Comments(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("reddit_my_comment.listMyComment", "query_error", err, "resp", resp, "opts", opts)
			return nil, err
		}
		for _, i := range items {
			count++
			row := myCommentRow{
				i,
				count,
			}
			d.StreamListItem(ctx, row)
		}
		// Stop if we've reached the end, or the max target limit
		if resp.After == "" {
			break
		}
		// Set the page and continue
		opts.After = resp.After
	}

	return nil, nil
}

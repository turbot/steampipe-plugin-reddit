package reddit

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "username",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getRedditAuthenticatedUser,
			Description: "The authorized username.",
			Transform:   transform.FromValue(),
		},
	}, c...)
}

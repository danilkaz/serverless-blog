package ydb

import (
	"context"
	"fmt"
	"path"
	"serverless-blog/internal/models"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

type Client struct {
	db *ydb.Driver
}

func New(endpoint, accessToken string) (*Client, error) {
	db, err := ydb.Open(context.Background(), endpoint, ydb.WithAccessTokenCredentials(accessToken))
	if err != nil {
		return nil, fmt.Errorf("open connection error: %w", err)
	}

	return &Client{db: db}, nil
}

func (c *Client) CreateTable(ctx context.Context) error {
	err := c.db.Table().Do(ctx,
		func(ctx context.Context, s table.Session) error {
			return s.CreateTable(ctx, path.Join(c.db.Name(), "posts"),
				options.WithColumn("name", types.TypeUTF8),
				options.WithColumn("text", types.TypeUTF8),
				options.WithColumn("created_at", types.TypeTimestamp),
				options.WithPrimaryKeyColumn("name"),
			)
		},
	)
	if err != nil {
		return fmt.Errorf("create table error: %w", err)
	}

	return nil
}

func (c *Client) CreatePost(ctx context.Context, post models.Post) error {
	err := c.db.Table().Do(ctx,
		func(ctx context.Context, s table.Session) error {
			_, res, err := s.Execute(ctx,
				table.TxControl(
					table.BeginTx(
						table.WithSerializableReadWrite(),
					),
					table.CommitTx(),
				),
				`
				DECLARE $name AS Utf8;
				DECLARE $text AS Utf8;
				DECLARE $created_at AS Timestamp;
				INSERT INTO posts (name, text, created_at) VALUES ($name, $text, $created_at);
				`,
				table.NewQueryParameters(
					table.ValueParam("$name", types.UTF8Value(post.Name)),
					table.ValueParam("$text", types.UTF8Value(post.Text)),
					table.ValueParam("$created_at", types.TimestampValueFromTime(post.CreatedAt)),
				),
			)
			if err != nil {
				return fmt.Errorf("execute insert error: %w", err)
			}

			defer res.Close()

			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("create post error: %w", err)
	}

	return nil
}

func (c *Client) Posts(ctx context.Context) ([]models.Post, error) {
	posts := make([]models.Post, 0)

	err := c.db.Table().Do(ctx,
		func(ctx context.Context, s table.Session) error {
			res, err := s.StreamExecuteScanQuery(ctx,
				`SELECT name, text, created_at FROM posts ORDER BY created_at DESC;`,
				table.NewQueryParameters(),
			)
			if err != nil {
				return fmt.Errorf("execute select error: %w", err)
			}

			defer res.Close()

			for res.NextResultSet(ctx) {
				if err = res.Err(); err != nil {
					return fmt.Errorf("next result error: %w", err)
				}

				for res.NextRow() {
					post := models.Post{}

					err = res.ScanNamed(
						named.Required("name", &post.Name),
						named.Required("text", &post.Text),
						named.Required("created_at", &post.CreatedAt),
					)
					if err != nil {
						return fmt.Errorf("scan error: %w", err)
					}

					posts = append(posts, post)
				}
			}

			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("get posts error: %w", err)
	}

	return posts, nil
}

package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShorturlModel = (*customShorturlModel)(nil)

type (
	// ShorturlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShorturlModel.
	ShorturlModel interface {
		shorturlModel
		FindOneByUrl(ctx context.Context, url string) (*Shorturl, error)
	}

	customShorturlModel struct {
		*defaultShorturlModel
	}
)

// NewShorturlModel returns a model for the database table.
func NewShorturlModel(conn sqlx.SqlConn, c cache.CacheConf) ShorturlModel {
	return &customShorturlModel{
		defaultShorturlModel: newShorturlModel(conn, c),
	}
}

func (m *customShorturlModel) FindOneByUrl(ctx context.Context, url string) (*Shorturl, error) {
	var resp Shorturl
	query := fmt.Sprintf("select %s from %s where `url` = ? limit 1", shorturlRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, url)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

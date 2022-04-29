package logic

import (
	"context"
	"fmt"

	"shorturl.com/shorturl/rpc/transform/internal/svc"
	"shorturl.com/shorturl/rpc/transform/model"
	"shorturl.com/shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	res, err := l.svcCtx.Model.FindOneByUrl(l.ctx, in.Url)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if res != nil {
		return &transform.ShortenResp{
			Shorten: res.Shorten,
		}, err
	}

	shorten := shortenUrl(in.Url)
	_, err = l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{Url: in.Url, Shorten: shorten})
	if err != nil {
		return nil, err
	}

	return &transform.ShortenResp{
		Shorten: shorten,
	}, nil
}
func shortenUrl(url string) string {
	return fmt.Sprintf("%x", hash.Md5([]byte(url)))
}

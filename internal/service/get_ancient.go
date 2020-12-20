package service

import (
	"context"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/go-rpc/rpc-ancient/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-ancient/internal/storage"
)

func (s *AncientService) GetAncient(ctx context.Context, req *api.GetAncientReq) (*api.Ancient, error) {
	shici := &storage.ShiCi{}
	if err := s.mysqlCli.Where("id=?", req.Id).First(shici).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, rpcx.NewErrorf(codes.NotFound, "NotFound", "shici [%v] not exist", req.Id)
		}
		return nil, errors.Wrapf(err, "mysql select shici [%v] failed", req.Id)
	}

	rpcx.CtxSet(ctx, "shici", shici)

	return &api.Ancient{
		Id:      int64(shici.ID),
		Title:   shici.Title,
		Author:  shici.Author,
		Dynasty: shici.Dynasty,
		Content: shici.Content,
	}, nil
}

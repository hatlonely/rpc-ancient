package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-ancient/api/gen/go/api"
	"github.com/hatlonely/rpc-ancient/internal/storage"
)

func (s *AncientService) UpdateAncient(ctx context.Context, req *api.UpdateAncientReq) (*empty.Empty, error) {
	shici := &storage.ShiCi{
		ID:      int(req.Ancient.Id),
		Title:   req.Ancient.Title,
		Author:  req.Ancient.Author,
		Dynasty: req.Ancient.Dynasty,
		Content: req.Ancient.Content,
	}

	rpcx.CtxSet(ctx, "shici", shici)

	if err := s.mysqlCli.Update(shici).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			if err := s.mysqlCli.Create(shici).Error; err != nil {
				return nil, errors.Wrap(err, "mysql update failed")
			}
		} else {
			return nil, errors.Wrap(err, "mysql update failed")
		}
	}

	return nil, nil
}

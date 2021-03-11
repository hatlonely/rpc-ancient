package service

import (
	"context"
	"time"

	"github.com/hatlonely/go-kit/wrap"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-ancient/api/gen/go/api"
	"github.com/hatlonely/rpc-ancient/internal/storage"
)

type AncientService struct {
	api.AncientServiceServer
	mysqlCli *wrap.GORMDBWrapper
	esCli    *wrap.ESClientWrapper
	options  *Options
}

func NewAncientServiceWithOptions(mysqlCli *wrap.GORMDBWrapper, esCli *wrap.ESClientWrapper, options *Options) (*AncientService, error) {
	if !mysqlCli.HasTable(&storage.ShiCi{}) {
		if err := mysqlCli.Set(context.Background(), "gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(context.Background(), &storage.ShiCi{}).Unwrap().Error; err != nil {
			return nil, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	exists, err := esCli.IndexExists(options.ElasticsearchIndex).Do(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		createIndex, err := esCli.CreateIndex(options.ElasticsearchIndex).Body(storage.AncientMappingForElasticsearch).Do(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "esCli.CreateIndex failed.")
		}
		if !createIndex.Acknowledged {
			return nil, errors.New("esCli.CreateIndex not acknowledged")
		}
	}

	return &AncientService{
		mysqlCli: mysqlCli,
		esCli:    esCli,
		options:  options,
	}, nil
}

type Options struct {
	ElasticsearchIndex string
}

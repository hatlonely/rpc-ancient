package service

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-ancient/internal/storage"
)

type AncientService struct {
	mysqlCli *gorm.DB
	esCli    *elastic.Client
	options  *Options
}

func NewAncientServiceWithOptions(mysqlCli *gorm.DB, esCli *elastic.Client, options *Options) (*AncientService, error) {
	if !mysqlCli.HasTable(&storage.ShiCi{}) {
		if err := mysqlCli.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&storage.ShiCi{}).Error; err != nil {
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

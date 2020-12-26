package service

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-ancient/api/gen/go/api"
)

func split(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		for _, s := range []rune("，。、？！； ,.") {
			if r == s {
				return true
			}
		}
		return false
	})
}

func (s *AncientService) SearchAncient(ctx context.Context, req *api.SearchAncientReq) (*api.SearchAncientRes, error) {
	if req.Limit == 0 {
		req.Limit = 10
	}

	query := elastic.NewBoolQuery()
	for _, val := range split(req.Keyword) {
		if len(val) == 0 {
			continue
		}
		q := elastic.NewBoolQuery()
		q.Should(elastic.NewTermQuery("title", val))
		q.Should(elastic.NewTermQuery("author", val))
		q.Should(elastic.NewTermQuery("dynasty", val))
		q.Should(elastic.NewTermQuery("content", val))
		query.Must(q)
	}

	esCtx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()
	res, err := s.esCli.Search().Index(s.options.ElasticsearchIndex).Query(query).From(int(req.Offset)).Size(int(req.Limit)).Do(esCtx)
	if err != nil {
		return nil, errors.Wrap(err, "EsCli.Search failed")
	}

	var ancients []*api.Ancient
	for _, item := range res.Each(reflect.TypeOf(api.Ancient{})) {
		if t, ok := item.(api.Ancient); ok {
			ancients = append(ancients, &t)
		}
	}

	return &api.SearchAncientRes{Ancients: ancients}, nil
}

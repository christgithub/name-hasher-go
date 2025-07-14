package indexer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type ESClient interface {
	Index() IndexService
}

type IndexService interface {
	Index(name string) IndexService
	BodyJson(body interface{}) IndexService
	Do(ctx context.Context) (*IndexResponse, error)
}

type IndexResponse struct {
	ID string
}

type ElasticAdapter struct {
	Client *elasticsearch.Client
}

func (e *ElasticAdapter) Index() IndexService {
	return &ElasticIndexService{client: e.Client}
}

type ElasticIndexService struct {
	client    *elasticsearch.Client
	indexName string
	body      interface{}
}

func (s *ElasticIndexService) Index(name string) IndexService {
	s.indexName = name
	return s
}

func (s *ElasticIndexService) BodyJson(body interface{}) IndexService {
	s.body = body
	return s
}

func (s *ElasticIndexService) Do(ctx context.Context) (*IndexResponse, error) {
	jsonData, err := json.Marshal(s.body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	res, err := s.client.Index(
		s.indexName,
		bytes.NewReader(jsonData),
		s.client.Index.WithContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("indexing failed: %w", err)
	}
	defer res.Body.Close()

	return &IndexResponse{ID: s.indexName}, nil
}

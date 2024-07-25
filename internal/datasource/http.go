// Package datasource - получение данных, которые надо отрепостить и отправка подтверждения репоста
package datasource

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/dnsoftware/socialposter/internal/config"
	"github.com/dnsoftware/socialposter/logger"
)

// DataSource работа с источником данных для постинга
type DataSource struct {
	cfg *config.Config
}

type Posts map[int]Item

type Item struct {
	Title   string `json:"title"`
	Short   string `json:"short"`
	Content string `json:"content"`
	Files   Files  `json:"files"`
}

type Files map[string]any // набор ссылок на картинки для какой-то платформы

type Images map[string]string

func NewDataSource(cfg *config.Config) *DataSource {
	return &DataSource{
		cfg: cfg,
	}
}

func (d *DataSource) GetData() (Posts, error) {
	buf := &bytes.Buffer{}
	ctx := context.Background()
	url := d.cfg.DataEndpoint

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, buf)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	request.Header.Set("Token", d.cfg.AuthToken)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	posts := Posts{}
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	return posts, nil
}

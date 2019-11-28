package internal

import (
	"context"
)

type CacheSaver interface {
	HasItem(context.Context, TimeRange) (bool, error)
	Save(context.Context, TimeRange) error
}

type FileCache struct {
	Path string
}

func NewFileCache(path string) *FileCache {
	return &FileCache{Path: path}
}

func (c *FileCache) HasItem(ctx context.Context, night TimeRange) (bool, error) {
	return false, nil
}

func (c *FileCache) Save(ctx context.Context, night TimeRange) error {
	return nil
}

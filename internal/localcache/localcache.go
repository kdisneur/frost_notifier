package localcache

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
)

type LocalCache struct {
	Path string
}

func New(path string) *LocalCache {
	return &LocalCache{Path: path}
}

func (l *LocalCache) HasItem(ctx context.Context, key string) (bool, error) {
	if _, err := os.Stat(l.Path); os.IsNotExist(err) {
		return false, nil
	}

	cache, err := ioutil.ReadFile(l.Path)
	if err != nil {
		return false, fmt.Errorf("can't read local cache file: %w", err)
	}

	return string(cache) == key, nil
}

func (l *LocalCache) Save(ctx context.Context, key string) error {
	f, err := os.OpenFile(l.Path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("can't open local cache file for writing: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(key)
	if err != nil {
		return fmt.Errorf("can't write in local cache file: %w", err)
	}

	return nil
}

package http_locks

import (
	"context"

	http_storage "github.com/storage-lock/go-http-storage"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

// LockFactory 基于 HTTP 存储的锁工厂
type LockFactory struct {
	*storage_lock_factory.StorageLockFactory[any]
}

var newHTTPStorage = http_storage.NewHTTPStorage

// NewLockFactory 创建锁工厂
func NewLockFactory() (*LockFactory, error) {
	s, err := newHTTPStorage(context.Background())
	if err != nil {
		return nil, err
	}
	factory := storage_lock_factory.NewStorageLockFactory[any](s, nil)
	return &LockFactory{
		StorageLockFactory: factory,
	}, nil
}

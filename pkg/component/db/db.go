package db

import (
	"context"
	"sync"
)

var _ Inter = (*DB)(nil)

type Inter interface {
	Upsert(ctx context.Context, key string, data []byte)
	Table(ctx context.Context) KeyDB
}

type DB struct {
	mu    sync.Mutex
	table KeyDB
}

type KeyDB map[string][]byte

func New(ctx context.Context) *DB {
	return &DB{
		mu:    sync.Mutex{},
		table: make(KeyDB),
	}
}

func (d *DB) Upsert(ctx context.Context, key string, data []byte) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.table[key] = data
}

// Table returns raw table unsafe method
func (d *DB) Table(ctx context.Context) KeyDB {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.table
}

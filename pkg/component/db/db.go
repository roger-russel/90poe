package db

import (
	"context"
)

var _ Inter = (*DB)(nil)

type Inter interface {
	Upsert(ctx context.Context, key string, data []byte)
}

type DB struct {
	table KeyDB
}

type KeyDB map[string][]byte

func New(ctx context.Context) *DB {
	return &DB{
		table: make(KeyDB),
	}
}

func (d *DB) Upsert(ctx context.Context, key string, data []byte) {
	d.table[key] = data
}

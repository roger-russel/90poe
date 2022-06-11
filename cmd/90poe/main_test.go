package main

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/roger.russel/90poe/internal/container"
	"github.com/roger.russel/90poe/internal/flags"
	"github.com/roger.russel/90poe/pkg/component/db"
)

const bddPath string = "../../_test/assets/bdd"

func minify(t *testing.T, d []byte) []byte {
	cp := bytes.NewBuffer([]byte{})
	if err := json.Compact(cp, d); err != nil {
		t.Errorf("unable to minimize json: %v", string(d))

		return nil
	}

	return cp.Bytes()
}

func Test_run(t *testing.T) {
	type args struct {
		flags flags.Flags
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantedTable db.KeyDB
	}{
		{
			name: "simple",
			args: args{
				flags: flags.Flags{
					File: bddPath + "/simple.json",
				},
			},
			wantErr: false,
			wantedTable: func() db.KeyDB {
				t := make(db.KeyDB)
				t["A"] = []byte(`{"foo":"boo A"}`)
				return t
			}(),
		},
		{
			name: "duplicated",
			args: args{
				flags: flags.Flags{
					File: bddPath + "/duplicated.json",
				},
			},
			wantErr: false,
			wantedTable: func() db.KeyDB {
				t := make(db.KeyDB)
				t["A"] = []byte(`{"foo":"boo A 2"}`)
				t["B"] = []byte(`{"foo":"boo B"}`)
				return t
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			ctx, dep, err := container.New(ctx, tt.args.flags)

			if err != nil {
				t.Error(err)
			}

			if err := run(ctx, cancel, dep); err != nil != tt.wantErr {
				t.Errorf("run got error: %v, differ front what is wanted, %v", err, tt.wantErr)
			}

			table := dep.Cmp.DB.Table(ctx)
			time.Sleep(100 * time.Nanosecond)

			if len(table) != len(tt.wantedTable) {
				t.Errorf("tables size are different wanted %v, got %v, diff %v", len(tt.wantedTable), len(table), cmp.Diff(table, tt.wantedTable))
			}

			for i, wt := range tt.wantedTable {
				tab := minify(t, table[i])
				if !cmp.Equal(wt, tab) {
					t.Errorf("data are different on key: %v, want %v, got %v", i, string(wt), string(tab))
				}
			}
		})
	}
}

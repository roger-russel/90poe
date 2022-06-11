package streamer

import (
	"bytes"
	"context"
	"encoding/json"
	"sync"
	"testing"
)

func minify(t *testing.T, d []byte) string {
	cp := bytes.NewBuffer([]byte{})
	if err := json.Compact(cp, d); err != nil {
		t.Errorf("unable to minimize json: %v", string(d))

		return ""
	}

	return cp.String()
}

func TestJson_Stream(t *testing.T) {
	type args struct {
		ctx      context.Context
		filePath string
	}

	type want struct {
		KeyName    string
		KeyContent string
	}

	tests := []struct {
		name string
		conf JSONConfig
		args args
		want []want
	}{
		{
			name: "object json",
			args: args{
				ctx:      context.Background(),
				filePath: "../../../_test/assets/json-streamer/object.json",
			},
			conf: JSONConfig{
				BufferParserSize: 100,
			},
			want: []want{
				{
					KeyName:    "A",
					KeyContent: `{"foo":"boo A"}`,
				},
				{
					KeyName:    "B",
					KeyContent: `{"foo":"boo B"}`,
				},
				{
					KeyName:    "C",
					KeyContent: `{"foo":"boo C"}`,
				},
				{
					KeyName:    "D{\\\"",
					KeyContent: `{"foo":"boo { \\{ \" D"}`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := NewJSON(tt.conf)

			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)

			wg := sync.WaitGroup{}
			wg.Add(len(tt.want))

			chDataOutput := make(chan Data)
			go func(ctx context.Context, t *testing.T, wg *sync.WaitGroup, ch chan Data, want []want) {
				i := 0
				for {
					select {
					case <-ctx.Done():
						return
					case d := <-ch:
						if want[i].KeyName != string(d.KeyName) {
							t.Errorf("KeyName: got %v, want %v", string(d.KeyName), want[i].KeyName)
						}
						content := minify(t, d.KeyContent)
						if want[i].KeyContent != content {
							t.Errorf("KeyContent: got: %v, want: %v, raw: %v", content, want[i].KeyContent, string(d.KeyContent))
						}
						wg.Done()
					}
					i++
				}
			}(ctx, t, &wg, chDataOutput, tt.want)

			if err := j.StreamFile(tt.args.ctx, tt.args.filePath, chDataOutput); err != nil {
				t.Errorf("error while reading file: %v", err)
			}

			wg.Wait()
			cancel()
			close(chDataOutput)
		})
	}
}

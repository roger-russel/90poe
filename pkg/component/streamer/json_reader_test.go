package streamer

import "testing"

func TestJson_handleBackslash(t *testing.T) {
	type args struct {
		c []byte
	}
	type want struct {
		skip   bool
		scaped bool
	}
	tests := []struct {
		name string
		args args
		want []want
	}{
		{
			name: "simple",
			args: args{c: []byte("a")},
			want: []want{
				{
					skip:   false,
					scaped: false,
				},
			},
		},
		{
			name: "backslash",
			args: args{c: []byte("\\")},
			want: []want{
				{
					skip:   true,
					scaped: false,
				},
			},
		},
		{
			name: "composite",
			args: args{c: []byte("aa")},
			want: []want{
				{
					skip:   false,
					scaped: false,
				},
				{
					skip:   false,
					scaped: false,
				},
			},
		},
		{
			name: "scaped",
			args: args{c: []byte("a\\a")},
			want: []want{
				{
					skip:   false,
					scaped: false,
				},
				{
					skip:   true,
					scaped: false,
				},
				{
					skip:   false,
					scaped: true,
				},
			},
		},
		{
			name: "scaped multiple times",
			args: args{c: []byte("a\\\\\\a\\\\")},
			want: []want{
				{ // a
					skip:   false,
					scaped: false,
				},
				{ // \
					skip:   true,
					scaped: false,
				},
				{ // \
					skip:   false,
					scaped: true,
				},
				{ // \
					skip:   true,
					scaped: false,
				},
				{ // a
					skip:   false,
					scaped: true,
				},
				{ // /
					skip:   true,
					scaped: false,
				},
				{ // /
					skip:   false,
					scaped: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSON{
				bufferParserSize: DefaultBufferParserSize,
			}

			for i, c := range tt.args.c {
				if got := j.handleBackslash(c); got != tt.want[i].skip {
					t.Errorf("Json.handleBackslash() skipe = %v, want %v, on %v, with %v", got, tt.want[i].skip, i, string(c))
				}

				if j.escapedChar != tt.want[i].scaped {
					t.Errorf("Json.handleBackslash() scaped char = %v, want %v, on %v, with %v", j.escapedChar, tt.want[i].scaped, i, string(c))
				}
			}
		})
	}
}

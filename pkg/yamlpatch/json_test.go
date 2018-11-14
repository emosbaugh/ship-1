package yamlpatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONGenerate(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				a: []byte(`{"a":100,"b":200,"c":"hello"}`),
				b: []byte(`{"b":200,"c":"goodbye","d":"hello again"}`),
			},
			want: []byte(`[` +
				`{"op":"remove","path":"/a"},` +
				`{"op":"replace","path":"/c","value":"goodbye"},` +
				`{"op":"add","path":"/d","value":"hello again"}` +
				`]`),
		},
		{
			name: "array",
			args: args{
				a: []byte(`{"a":[1,3]}`),
				b: []byte(`{"a":[1,2,3]}`),
			},
			want: []byte(`[` +
				`{"op":"add","path":"/a/1","value":2}` +
				`]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonGenerate(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.JSONEq(t, string(tt.want), string(got))
		})
	}
}

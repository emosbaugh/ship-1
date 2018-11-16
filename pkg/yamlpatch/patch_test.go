package yamlpatch

import (
	"reflect"
	"testing"

	yamlpatch "github.com/krishicks/yaml-patch"
	"github.com/replicatedhq/ship/pkg/testing/logger"
)

func TestPatch_Apply(t *testing.T) {
	type args struct {
		doc []byte
	}
	tests := []struct {
		name    string
		p       []byte
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "missing",
			p: []byte(`[{"op":"replace","path":"/key1","value":2},` +
				`{"op":"remove","path":"/key3"},` +
				`{"op":"add","path":"/key4/key5","value":4}]`),
			args: args{
				doc: []byte(`key1: 1`),
			},
			want: []byte("key1: 2\n"),
		},

		{
			name: "maintain order",
			p: []byte(`[{"op":"replace","path":"/key2","value":2},` +
				`{"op":"replace","path":"/key5","value":5},` +
				`{"op":"add","path":"/key1/key6","value":7},` +
				`{"op":"remove","path":"/key1/key4"}]`),
			args: args{
				doc: []byte("key5: 1\nkey1:\n  key4: 3\n  key3: 4\nkey2: 5"),
			},
			want: []byte("key5: 5\nkey1:\n  key3: 4\n  key6: 7\nkey2: 2\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := yamlpatch.DecodePatch(tt.p)
			if err != nil {
				t.Errorf("DecodePatch() error = %v", err)
				return
			}
			log := &logger.TestLogger{T: t}
			got, err := ApplyPatchNonStrict(p, tt.args.doc, log)
			if (err != nil) != tt.wantErr {
				t.Errorf("Patch.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Patch.Apply() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

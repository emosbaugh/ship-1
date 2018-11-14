package yamlpatch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

func TestGenerate(t *testing.T) {
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
				a: []byte(`a: 100
b: 200
c: hello`),
				b: []byte(`b: 200
c: goodbye
d: hello again`),
			},
			want: []byte(`[{"op":"remove","path":"/a"},` +
				`{"op":"replace","path":"/c","value":"goodbye"},` +
				`{"op":"add","path":"/d","value":"hello again"}` +
				`]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assertYAMLEq(t, string(tt.want), string(got))
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		doc       []byte
		patchJSON []byte
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
				doc: []byte(`a: 100
b: 200
c: hello`),
				patchJSON: []byte(`[{"op":"replace","path":"/c","value":"goodbye"},` +
					`{"op":"add","path":"/d","value":"hello again"},` +
					`{"op":"remove","path":"/a"}]`),
			},
			want: []byte(`b: 200
c: goodbye
d: hello again`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Apply(tt.args.doc, tt.args.patchJSON)
			if (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assertYAMLEq(t, string(tt.want), string(got))
		})
	}
}

func assertYAMLEq(t assert.TestingT, expected string, actual string, msgAndArgs ...interface{}) bool {
	var expectedYAMLAsInterface, actualYAMLAsInterface interface{}

	if err := yaml.Unmarshal([]byte(expected), &expectedYAMLAsInterface); err != nil {
		return assert.Fail(t, fmt.Sprintf("Expected value ('%s') is not valid yaml.\nYAML parsing error: '%s'", expected, err.Error()), msgAndArgs...)
	}

	if err := yaml.Unmarshal([]byte(actual), &actualYAMLAsInterface); err != nil {
		return assert.Fail(t, fmt.Sprintf("Input ('%s') needs to be valid yaml.\nYAML parsing error: '%s'", actual, err.Error()), msgAndArgs...)
	}

	return assert.Equal(t, expectedYAMLAsInterface, actualYAMLAsInterface, msgAndArgs...)
}

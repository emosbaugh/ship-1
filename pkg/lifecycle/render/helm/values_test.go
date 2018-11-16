package helm

import (
	"testing"

	"github.com/replicatedhq/ship/pkg/testing/logger"
	"github.com/stretchr/testify/require"
)

func TestMergeHelmValues(t *testing.T) {
	tests := []struct {
		name     string
		base     string
		vendor   string
		user     string
		expected string
	}{
		{
			name: "merge, vendor values only",
			base: "",
			vendor: `key1: 1
key2:
  - item1
deep_key:
  level1:
    level2:
      myvalue: 3
key3: a`,
			user: "",
			expected: `key1: 1
key2:
  - item1
deep_key:
  level1:
    level2:
      myvalue: 3
key3: a`,
		},

		{
			name: "merge, vendor and user values",
			base: `key1: 1
key2:
  - item1
deep_key:
  level1:
    level2:
      myvalue: 3
key3: a`,
			vendor: `key1: 1
key2:
  - item1
deep_key:
  level1:
    newkey: added-by-vendor
    level2:
      myvalue: 5
key3: modified-by-vendor`,
			user: `[` +
				`{"op":"add","path":"/key2/1","value":"item2_added_by_user"},` +
				`{"op":"replace","path":"/deep_key/level1/level2/myvalue","value":"modified-by-user-5"}` +
				`]`,
			expected: `key1: 1
key2:
- item1
- item2_added_by_user
deep_key:
  level1:
    newkey: added-by-vendor
    level2:
      myvalue: modified-by-user-5
key3: modified-by-vendor
`,
		},

		{
			name: "merge, vendor value no longer exists",
			base: `key1: 1
key2: 2`,
			vendor: `key1: 1`,
			user: `[` +
				`{"op":"replace","path":"/key2","value":"222"}` +
				`]`,
			expected: `key1: 1`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)
			merged, err := MergeHelmValues(test.base, test.vendor, test.user, &logger.TestLogger{T: t})
			req.NoError(err)
			req.Equal(test.expected, merged)
		})
	}
}

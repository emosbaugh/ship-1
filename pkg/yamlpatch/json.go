package yamlpatch

import (
	"encoding/json"
	"sort"

	"github.com/mattbaird/jsonpatch"
)

func jsonGenerate(a, b []byte) ([]byte, error) {
	patch, err := jsonpatch.CreatePatch(a, b)
	if err != nil {
		return nil, err
	}
	sort.Sort(jsonpatch.ByPath(patch))
	return json.Marshal(patch)
}

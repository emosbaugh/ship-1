package yamlpatch

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	yamlpatch "github.com/krishicks/yaml-patch"
)

// ApplyPatchNonStrict will apply ops one at a time ignoring errors
func ApplyPatchNonStrict(p yamlpatch.Patch, doc []byte, logger log.Logger) ([]byte, error) {
	for _, op := range p {
		subpatch := yamlpatch.Patch{op}
		d, err := subpatch.Apply(doc)
		if err != nil {
			level.Warn(logger).Log("op", op.Op, "path", op.Path, "err", err)
		} else {
			doc = d
		}
	}
	return doc, nil
}

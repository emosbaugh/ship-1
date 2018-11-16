package yamlpatch

import (
	"github.com/ghodss/yaml"
	"github.com/go-kit/kit/log"
	"github.com/krishicks/yaml-patch"
	"github.com/pkg/errors"
)

func Generate(a, b []byte) ([]byte, error) {
	aJSON, err := yaml.YAMLToJSON(a)
	if err != nil {
		return nil, errors.Wrap(err, "yaml to json a")
	}
	bJSON, err := yaml.YAMLToJSON(b)
	if err != nil {
		return nil, errors.Wrap(err, "yaml to json b")
	}
	return jsonGenerate(aJSON, bJSON)
}

func Apply(doc, patchJSON []byte, logger log.Logger) ([]byte, error) {
	patch, err := yaml.JSONToYAML(patchJSON)
	if err != nil {
		return nil, errors.Wrap(err, "patch json to yaml")
	}
	p, err := yamlpatch.DecodePatch(patch)
	if err != nil {
		return nil, errors.Wrap(err, "decode patch")
	}
	return ApplyPatchNonStrict(p, doc, logger)
}

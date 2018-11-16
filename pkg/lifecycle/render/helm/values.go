package helm

import (
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/replicatedhq/ship/pkg/yamlpatch"
)

// MergeHelmValues merges user edited values from state file and vendor values from upstream Helm
// repo.
// baseValues is the original config from state
// vendorValues is the new config from current chart
// userPatch is the jsonpatch from state
func MergeHelmValues(baseValues, vendorValues, userPatch string, logger log.Logger) (string, error) {
	var merged []byte
	if vendorValues != "" {
		merged = []byte(vendorValues)
	} else {
		merged = []byte(baseValues)
	}

	if userPatch != "" {
		var err error
		merged, err = yamlpatch.Apply(merged, []byte(userPatch), logger)
		if err != nil {
			return "", errors.Wrapf(err, "apply user json patch")
		}
	}

	return string(merged), nil
}

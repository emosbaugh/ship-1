package helm

import (
	"github.com/pkg/errors"
	"github.com/replicatedhq/ship/pkg/yamlpatch"
)

// Merges user edited values from state file and vendor values from upstream Helm repo.
// base is the original config from state
// user is the jsonpatch from state
// vendor is the new config from current chart
// Value priorities: user, vendor, base
func MergeHelmValues(baseValues, userPatch, vendorValues string) (string, error) {
	var merged []byte
	if vendorValues != "" {
		merged = []byte(vendorValues)
	} else {
		merged = []byte(baseValues)
	}

	if userPatch != "" {
		var err error
		merged, err = yamlpatch.Apply(merged, []byte(userPatch))
		if err != nil {
			return "", errors.Wrapf(err, "apply user json patch")
		}
	}

	return string(merged), nil
}

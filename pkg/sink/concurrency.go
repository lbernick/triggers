package sink

import (
	"fmt"
	"strings"

	triggersv1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1beta1"
)

// ResolveConcurrencyLabels TODO
func ResolveConcurrencyLabels(c *triggersv1.ConcurrencyControl, params []triggersv1.Param) map[string]string {
	if c == nil {
		return nil
	}
	out := c.Spec.Key
	for _, param := range params {
		// Assume the param is valid
		paramVariable := fmt.Sprintf("$(params.%s)", param.Name)
		out = strings.ReplaceAll(out, paramVariable, param.Value)
	}
	return map[string]string{"concurrency": out}
}

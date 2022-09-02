package sink_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	triggersv1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1beta1"
	"github.com/tektoncd/triggers/pkg/sink"
)

func TestGetConcurrencyLabels(t *testing.T) {
	cc := triggersv1.ConcurrencyControl{Spec: triggersv1.ConcurrencySpec{
		Key: "$(params.repo-full-name)",
	}}
	params := []triggersv1.Param{{
		Name:  "repo-full-name",
		Value: "tektoncd/pipeline",
	}}
	got := sink.ResolveConcurrencyLabels(&cc, params)
	want := map[string]string{"concurrency": "tektoncd/pipeline"}
	if d := cmp.Diff(want, got); d != "" {
		t.Error(d)
	}
}

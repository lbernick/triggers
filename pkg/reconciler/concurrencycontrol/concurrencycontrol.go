package concurrencycontrol

import (
	"context"

	pipelinelisters "github.com/tektoncd/pipeline/pkg/client/listers/pipeline/v1beta1"
	"github.com/tektoncd/triggers/pkg/apis/triggers/v1beta1"
	concurrencycontrolreconciler "github.com/tektoncd/triggers/pkg/client/injection/reconciler/triggers/v1beta1/concurrencycontrol"
	pkgreconciler "knative.dev/pkg/reconciler"
)

const ControllerName = "ConcurrencyControl"

type Reconciler struct {
	pipelinerunlister pipelinelisters.PipelineRunLister
}

var (
	_ concurrencycontrolreconciler.Interface = (*Reconciler)(nil)
)

func (r *Reconciler) ReconcileKind(ctx context.Context, cc *v1beta1.ConcurrencyControl) pkgreconciler.Event {
	//return nil
	// Need to get all PipelineRuns the concurrency control owns; filter to just running and pending ones
	// (might be a problem with how we implement retries since taskruns get marked as failed for a brief time-- prob not relevant to pipelineruns)
	// we could put references to pipelinerun in concurrencycontrol status
	//prs, err := r.pipelinerunlister.PipelineRuns("namespace").List() <- can list by label selector
	
}

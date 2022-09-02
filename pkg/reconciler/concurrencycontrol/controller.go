/*
Copyright 2021 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package concurrencycontrol

import (
	"context"

	pipelineruninformer "github.com/tektoncd/pipeline/pkg/client/injection/informers/pipeline/v1beta1/pipelinerun"
	"github.com/tektoncd/triggers/pkg/apis/triggers/v1beta1"
	concurrencycontrolinformer "github.com/tektoncd/triggers/pkg/client/injection/informers/triggers/v1beta1/concurrencycontrol"
	concurrencycontrolreconciler "github.com/tektoncd/triggers/pkg/client/injection/reconciler/triggers/v1beta1/concurrencycontrol"
	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
)

func NewController() func(context.Context, configmap.Watcher) *controller.Impl {
	return func(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
		concurrencyControlInformer := concurrencycontrolinformer.Get(ctx)
		reconciler := &Reconciler{}

		impl := concurrencycontrolreconciler.NewImpl(ctx, reconciler, func(impl *controller.Impl) controller.Options {
			return controller.Options{
				AgentName: ControllerName,
			}
		})

		concurrencyControlInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

		pipelineruninformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
			FilterFunc: controller.FilterController(&v1beta1.ConcurrencyControl{}),
			Handler:    controller.HandleAll(impl.EnqueueControllerOf),
		})

		return impl
	}
}

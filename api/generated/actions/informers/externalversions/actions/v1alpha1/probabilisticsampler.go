/*
Copyright 2022.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	actionsv1alpha1 "github.com/odigos-io/odigos/api/actions/v1alpha1"
	versioned "github.com/odigos-io/odigos/api/generated/actions/clientset/versioned"
	internalinterfaces "github.com/odigos-io/odigos/api/generated/actions/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/odigos-io/odigos/api/generated/actions/listers/actions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProbabilisticSamplerInformer provides access to a shared informer and lister for
// ProbabilisticSamplers.
type ProbabilisticSamplerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ProbabilisticSamplerLister
}

type probabilisticSamplerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProbabilisticSamplerInformer constructs a new informer for ProbabilisticSampler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProbabilisticSamplerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProbabilisticSamplerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProbabilisticSamplerInformer constructs a new informer for ProbabilisticSampler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProbabilisticSamplerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ActionsV1alpha1().ProbabilisticSamplers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ActionsV1alpha1().ProbabilisticSamplers(namespace).Watch(context.TODO(), options)
			},
		},
		&actionsv1alpha1.ProbabilisticSampler{},
		resyncPeriod,
		indexers,
	)
}

func (f *probabilisticSamplerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProbabilisticSamplerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *probabilisticSamplerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&actionsv1alpha1.ProbabilisticSampler{}, f.defaultInformer)
}

func (f *probabilisticSamplerInformer) Lister() v1alpha1.ProbabilisticSamplerLister {
	return v1alpha1.NewProbabilisticSamplerLister(f.Informer().GetIndexer())
}

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/odigos-io/odigos/api/actions/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ProbabilisticSamplerLister helps list ProbabilisticSamplers.
// All objects returned here must be treated as read-only.
type ProbabilisticSamplerLister interface {
	// List lists all ProbabilisticSamplers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProbabilisticSampler, err error)
	// ProbabilisticSamplers returns an object that can list and get ProbabilisticSamplers.
	ProbabilisticSamplers(namespace string) ProbabilisticSamplerNamespaceLister
	ProbabilisticSamplerListerExpansion
}

// probabilisticSamplerLister implements the ProbabilisticSamplerLister interface.
type probabilisticSamplerLister struct {
	listers.ResourceIndexer[*v1alpha1.ProbabilisticSampler]
}

// NewProbabilisticSamplerLister returns a new ProbabilisticSamplerLister.
func NewProbabilisticSamplerLister(indexer cache.Indexer) ProbabilisticSamplerLister {
	return &probabilisticSamplerLister{listers.New[*v1alpha1.ProbabilisticSampler](indexer, v1alpha1.Resource("probabilisticsampler"))}
}

// ProbabilisticSamplers returns an object that can list and get ProbabilisticSamplers.
func (s *probabilisticSamplerLister) ProbabilisticSamplers(namespace string) ProbabilisticSamplerNamespaceLister {
	return probabilisticSamplerNamespaceLister{listers.NewNamespaced[*v1alpha1.ProbabilisticSampler](s.ResourceIndexer, namespace)}
}

// ProbabilisticSamplerNamespaceLister helps list and get ProbabilisticSamplers.
// All objects returned here must be treated as read-only.
type ProbabilisticSamplerNamespaceLister interface {
	// List lists all ProbabilisticSamplers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProbabilisticSampler, err error)
	// Get retrieves the ProbabilisticSampler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProbabilisticSampler, error)
	ProbabilisticSamplerNamespaceListerExpansion
}

// probabilisticSamplerNamespaceLister implements the ProbabilisticSamplerNamespaceLister
// interface.
type probabilisticSamplerNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.ProbabilisticSampler]
}

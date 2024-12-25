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
	v1alpha1 "github.com/odigos-io/odigos/api/odigos/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// InstrumentedApplicationLister helps list InstrumentedApplications.
// All objects returned here must be treated as read-only.
type InstrumentedApplicationLister interface {
	// List lists all InstrumentedApplications in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstrumentedApplication, err error)
	// InstrumentedApplications returns an object that can list and get InstrumentedApplications.
	InstrumentedApplications(namespace string) InstrumentedApplicationNamespaceLister
	InstrumentedApplicationListerExpansion
}

// instrumentedApplicationLister implements the InstrumentedApplicationLister interface.
type instrumentedApplicationLister struct {
	listers.ResourceIndexer[*v1alpha1.InstrumentedApplication]
}

// NewInstrumentedApplicationLister returns a new InstrumentedApplicationLister.
func NewInstrumentedApplicationLister(indexer cache.Indexer) InstrumentedApplicationLister {
	return &instrumentedApplicationLister{listers.New[*v1alpha1.InstrumentedApplication](indexer, v1alpha1.Resource("instrumentedapplication"))}
}

// InstrumentedApplications returns an object that can list and get InstrumentedApplications.
func (s *instrumentedApplicationLister) InstrumentedApplications(namespace string) InstrumentedApplicationNamespaceLister {
	return instrumentedApplicationNamespaceLister{listers.NewNamespaced[*v1alpha1.InstrumentedApplication](s.ResourceIndexer, namespace)}
}

// InstrumentedApplicationNamespaceLister helps list and get InstrumentedApplications.
// All objects returned here must be treated as read-only.
type InstrumentedApplicationNamespaceLister interface {
	// List lists all InstrumentedApplications in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstrumentedApplication, err error)
	// Get retrieves the InstrumentedApplication from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstrumentedApplication, error)
	InstrumentedApplicationNamespaceListerExpansion
}

// instrumentedApplicationNamespaceLister implements the InstrumentedApplicationNamespaceLister
// interface.
type instrumentedApplicationNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.InstrumentedApplication]
}

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	odigosv1alpha1 "github.com/odigos-io/odigos/api/generated/odigos/applyconfiguration/odigos/v1alpha1"
	v1alpha1 "github.com/odigos-io/odigos/api/odigos/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInstrumentationConfigs implements InstrumentationConfigInterface
type FakeInstrumentationConfigs struct {
	Fake *FakeOdigosV1alpha1
	ns   string
}

var instrumentationconfigsResource = v1alpha1.SchemeGroupVersion.WithResource("instrumentationconfigs")

var instrumentationconfigsKind = v1alpha1.SchemeGroupVersion.WithKind("InstrumentationConfig")

// Get takes name of the instrumentationConfig, and returns the corresponding instrumentationConfig object, and an error if there is any.
func (c *FakeInstrumentationConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstrumentationConfig, err error) {
	emptyResult := &v1alpha1.InstrumentationConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(instrumentationconfigsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InstrumentationConfig), err
}

// List takes label and field selectors, and returns the list of InstrumentationConfigs that match those selectors.
func (c *FakeInstrumentationConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstrumentationConfigList, err error) {
	emptyResult := &v1alpha1.InstrumentationConfigList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(instrumentationconfigsResource, instrumentationconfigsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstrumentationConfigList{ListMeta: obj.(*v1alpha1.InstrumentationConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstrumentationConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instrumentationConfigs.
func (c *FakeInstrumentationConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(instrumentationconfigsResource, c.ns, opts))

}

// Create takes the representation of a instrumentationConfig and creates it.  Returns the server's representation of the instrumentationConfig, and an error, if there is any.
func (c *FakeInstrumentationConfigs) Create(ctx context.Context, instrumentationConfig *v1alpha1.InstrumentationConfig, opts v1.CreateOptions) (result *v1alpha1.InstrumentationConfig, err error) {
	emptyResult := &v1alpha1.InstrumentationConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(instrumentationconfigsResource, c.ns, instrumentationConfig, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InstrumentationConfig), err
}

// Update takes the representation of a instrumentationConfig and updates it. Returns the server's representation of the instrumentationConfig, and an error, if there is any.
func (c *FakeInstrumentationConfigs) Update(ctx context.Context, instrumentationConfig *v1alpha1.InstrumentationConfig, opts v1.UpdateOptions) (result *v1alpha1.InstrumentationConfig, err error) {
	emptyResult := &v1alpha1.InstrumentationConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(instrumentationconfigsResource, c.ns, instrumentationConfig, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InstrumentationConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstrumentationConfigs) UpdateStatus(ctx context.Context, instrumentationConfig *v1alpha1.InstrumentationConfig, opts v1.UpdateOptions) (result *v1alpha1.InstrumentationConfig, err error) {
	emptyResult := &v1alpha1.InstrumentationConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(instrumentationconfigsResource, "status", c.ns, instrumentationConfig, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InstrumentationConfig), err
}

// Delete takes name of the instrumentationConfig and deletes it. Returns an error if one occurs.
func (c *FakeInstrumentationConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(instrumentationconfigsResource, c.ns, name, opts), &v1alpha1.InstrumentationConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstrumentationConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(instrumentationconfigsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstrumentationConfigList{})
	return err
}

// Patch applies the patch and returns the patched instrumentationConfig.
func (c *FakeInstrumentationConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstrumentationConfig, err error) {
	emptyResult := &v1alpha1.InstrumentationConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(instrumentationconfigsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InstrumentationConfig), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied instrumentationConfig.
func (c *FakeInstrumentationConfigs) Apply(ctx context.Context, instrumentationConfig *odigosv1alpha1.InstrumentationConfigApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.InstrumentationConfig, err error) {
	if instrumentationConfig == nil {
		return nil, fmt.Errorf("instrumentationConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(instrumentationConfig)
	if err != nil {
		return nil, err
	}
	name := instrumentationConfig.Name
	if name == nil {
		return nil, fmt.Errorf("instrumentationConfig.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.InstrumentationConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(instrumentationconfigsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InstrumentationConfig), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeInstrumentationConfigs) ApplyStatus(ctx context.Context, instrumentationConfig *odigosv1alpha1.InstrumentationConfigApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.InstrumentationConfig, err error) {
	if instrumentationConfig == nil {
		return nil, fmt.Errorf("instrumentationConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(instrumentationConfig)
	if err != nil {
		return nil, err
	}
	name := instrumentationConfig.Name
	if name == nil {
		return nil, fmt.Errorf("instrumentationConfig.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.InstrumentationConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(instrumentationconfigsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InstrumentationConfig), err
}

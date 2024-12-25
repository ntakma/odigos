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

// FakeProcessors implements ProcessorInterface
type FakeProcessors struct {
	Fake *FakeOdigosV1alpha1
	ns   string
}

var processorsResource = v1alpha1.SchemeGroupVersion.WithResource("processors")

var processorsKind = v1alpha1.SchemeGroupVersion.WithKind("Processor")

// Get takes name of the processor, and returns the corresponding processor object, and an error if there is any.
func (c *FakeProcessors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Processor, err error) {
	emptyResult := &v1alpha1.Processor{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(processorsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Processor), err
}

// List takes label and field selectors, and returns the list of Processors that match those selectors.
func (c *FakeProcessors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProcessorList, err error) {
	emptyResult := &v1alpha1.ProcessorList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(processorsResource, processorsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProcessorList{ListMeta: obj.(*v1alpha1.ProcessorList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProcessorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested processors.
func (c *FakeProcessors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(processorsResource, c.ns, opts))

}

// Create takes the representation of a processor and creates it.  Returns the server's representation of the processor, and an error, if there is any.
func (c *FakeProcessors) Create(ctx context.Context, processor *v1alpha1.Processor, opts v1.CreateOptions) (result *v1alpha1.Processor, err error) {
	emptyResult := &v1alpha1.Processor{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(processorsResource, c.ns, processor, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Processor), err
}

// Update takes the representation of a processor and updates it. Returns the server's representation of the processor, and an error, if there is any.
func (c *FakeProcessors) Update(ctx context.Context, processor *v1alpha1.Processor, opts v1.UpdateOptions) (result *v1alpha1.Processor, err error) {
	emptyResult := &v1alpha1.Processor{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(processorsResource, c.ns, processor, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Processor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProcessors) UpdateStatus(ctx context.Context, processor *v1alpha1.Processor, opts v1.UpdateOptions) (result *v1alpha1.Processor, err error) {
	emptyResult := &v1alpha1.Processor{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(processorsResource, "status", c.ns, processor, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Processor), err
}

// Delete takes name of the processor and deletes it. Returns an error if one occurs.
func (c *FakeProcessors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(processorsResource, c.ns, name, opts), &v1alpha1.Processor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProcessors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(processorsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProcessorList{})
	return err
}

// Patch applies the patch and returns the patched processor.
func (c *FakeProcessors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Processor, err error) {
	emptyResult := &v1alpha1.Processor{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(processorsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Processor), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied processor.
func (c *FakeProcessors) Apply(ctx context.Context, processor *odigosv1alpha1.ProcessorApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Processor, err error) {
	if processor == nil {
		return nil, fmt.Errorf("processor provided to Apply must not be nil")
	}
	data, err := json.Marshal(processor)
	if err != nil {
		return nil, err
	}
	name := processor.Name
	if name == nil {
		return nil, fmt.Errorf("processor.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.Processor{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(processorsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Processor), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeProcessors) ApplyStatus(ctx context.Context, processor *odigosv1alpha1.ProcessorApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Processor, err error) {
	if processor == nil {
		return nil, fmt.Errorf("processor provided to Apply must not be nil")
	}
	data, err := json.Marshal(processor)
	if err != nil {
		return nil, err
	}
	name := processor.Name
	if name == nil {
		return nil, fmt.Errorf("processor.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.Processor{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(processorsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Processor), err
}

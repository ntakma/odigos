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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// InstrumentationLibraryOptionsApplyConfiguration represents an declarative configuration of the InstrumentationLibraryOptions type for use
// with apply.
type InstrumentationLibraryOptionsApplyConfiguration struct {
	LibraryName *string                          `json:"libraryName,omitempty"`
	Options     []ConfigOptionApplyConfiguration `json:"options,omitempty"`
}

// InstrumentationLibraryOptionsApplyConfiguration constructs an declarative configuration of the InstrumentationLibraryOptions type for use with
// apply.
func InstrumentationLibraryOptions() *InstrumentationLibraryOptionsApplyConfiguration {
	return &InstrumentationLibraryOptionsApplyConfiguration{}
}

// WithLibraryName sets the LibraryName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LibraryName field is set to the value of the last call.
func (b *InstrumentationLibraryOptionsApplyConfiguration) WithLibraryName(value string) *InstrumentationLibraryOptionsApplyConfiguration {
	b.LibraryName = &value
	return b
}

// WithOptions adds the given value to the Options field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Options field.
func (b *InstrumentationLibraryOptionsApplyConfiguration) WithOptions(values ...*ConfigOptionApplyConfiguration) *InstrumentationLibraryOptionsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOptions")
		}
		b.Options = append(b.Options, *values[i])
	}
	return b
}
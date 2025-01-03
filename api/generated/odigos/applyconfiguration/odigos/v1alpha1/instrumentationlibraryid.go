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

import (
	common "github.com/odigos-io/odigos/common"
)

// InstrumentationLibraryIdApplyConfiguration represents a declarative configuration of the InstrumentationLibraryId type for use
// with apply.
type InstrumentationLibraryIdApplyConfiguration struct {
	InstrumentationLibraryName *string          `json:"libraryName,omitempty"`
	SpanKind                   *common.SpanKind `json:"spanKind,omitempty"`
}

// InstrumentationLibraryIdApplyConfiguration constructs a declarative configuration of the InstrumentationLibraryId type for use with
// apply.
func InstrumentationLibraryId() *InstrumentationLibraryIdApplyConfiguration {
	return &InstrumentationLibraryIdApplyConfiguration{}
}

// WithInstrumentationLibraryName sets the InstrumentationLibraryName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InstrumentationLibraryName field is set to the value of the last call.
func (b *InstrumentationLibraryIdApplyConfiguration) WithInstrumentationLibraryName(value string) *InstrumentationLibraryIdApplyConfiguration {
	b.InstrumentationLibraryName = &value
	return b
}

// WithSpanKind sets the SpanKind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SpanKind field is set to the value of the last call.
func (b *InstrumentationLibraryIdApplyConfiguration) WithSpanKind(value common.SpanKind) *InstrumentationLibraryIdApplyConfiguration {
	b.SpanKind = &value
	return b
}

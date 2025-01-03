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

// ProbabilisticSamplerSpecApplyConfiguration represents a declarative configuration of the ProbabilisticSamplerSpec type for use
// with apply.
type ProbabilisticSamplerSpecApplyConfiguration struct {
	ActionName         *string                      `json:"actionName,omitempty"`
	Notes              *string                      `json:"notes,omitempty"`
	Disabled           *bool                        `json:"disabled,omitempty"`
	Signals            []common.ObservabilitySignal `json:"signals,omitempty"`
	SamplingPercentage *string                      `json:"sampling_percentage,omitempty"`
}

// ProbabilisticSamplerSpecApplyConfiguration constructs a declarative configuration of the ProbabilisticSamplerSpec type for use with
// apply.
func ProbabilisticSamplerSpec() *ProbabilisticSamplerSpecApplyConfiguration {
	return &ProbabilisticSamplerSpecApplyConfiguration{}
}

// WithActionName sets the ActionName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActionName field is set to the value of the last call.
func (b *ProbabilisticSamplerSpecApplyConfiguration) WithActionName(value string) *ProbabilisticSamplerSpecApplyConfiguration {
	b.ActionName = &value
	return b
}

// WithNotes sets the Notes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Notes field is set to the value of the last call.
func (b *ProbabilisticSamplerSpecApplyConfiguration) WithNotes(value string) *ProbabilisticSamplerSpecApplyConfiguration {
	b.Notes = &value
	return b
}

// WithDisabled sets the Disabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Disabled field is set to the value of the last call.
func (b *ProbabilisticSamplerSpecApplyConfiguration) WithDisabled(value bool) *ProbabilisticSamplerSpecApplyConfiguration {
	b.Disabled = &value
	return b
}

// WithSignals adds the given value to the Signals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Signals field.
func (b *ProbabilisticSamplerSpecApplyConfiguration) WithSignals(values ...common.ObservabilitySignal) *ProbabilisticSamplerSpecApplyConfiguration {
	for i := range values {
		b.Signals = append(b.Signals, values[i])
	}
	return b
}

// WithSamplingPercentage sets the SamplingPercentage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SamplingPercentage field is set to the value of the last call.
func (b *ProbabilisticSamplerSpecApplyConfiguration) WithSamplingPercentage(value string) *ProbabilisticSamplerSpecApplyConfiguration {
	b.SamplingPercentage = &value
	return b
}

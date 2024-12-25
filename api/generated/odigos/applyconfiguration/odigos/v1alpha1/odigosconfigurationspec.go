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

// OdigosConfigurationSpecApplyConfiguration represents a declarative configuration of the OdigosConfigurationSpec type for use
// with apply.
type OdigosConfigurationSpecApplyConfiguration struct {
	OdigosVersion               *string                                          `json:"odigosVersion,omitempty"`
	ConfigVersion               *int                                             `json:"configVersion,omitempty"`
	TelemetryEnabled            *bool                                            `json:"telemetryEnabled,omitempty"`
	OpenshiftEnabled            *bool                                            `json:"openshiftEnabled,omitempty"`
	IgnoredNamespaces           []string                                         `json:"ignoredNamespaces,omitempty"`
	IgnoredContainers           []string                                         `json:"ignoredContainers,omitempty"`
	Psp                         *bool                                            `json:"psp,omitempty"`
	ImagePrefix                 *string                                          `json:"imagePrefix,omitempty"`
	OdigletImage                *string                                          `json:"odigletImage,omitempty"`
	InstrumentorImage           *string                                          `json:"instrumentorImage,omitempty"`
	AutoscalerImage             *string                                          `json:"autoscalerImage,omitempty"`
	CollectorGateway            *CollectorGatewayConfigurationApplyConfiguration `json:"collectorGateway,omitempty"`
	GoAutoIncludeCodeAttributes *bool                                            `json:"goAutoIncludeCodeAttributes,omitempty"`
}

// OdigosConfigurationSpecApplyConfiguration constructs a declarative configuration of the OdigosConfigurationSpec type for use with
// apply.
func OdigosConfigurationSpec() *OdigosConfigurationSpecApplyConfiguration {
	return &OdigosConfigurationSpecApplyConfiguration{}
}

// WithOdigosVersion sets the OdigosVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OdigosVersion field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithOdigosVersion(value string) *OdigosConfigurationSpecApplyConfiguration {
	b.OdigosVersion = &value
	return b
}

// WithConfigVersion sets the ConfigVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigVersion field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithConfigVersion(value int) *OdigosConfigurationSpecApplyConfiguration {
	b.ConfigVersion = &value
	return b
}

// WithTelemetryEnabled sets the TelemetryEnabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TelemetryEnabled field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithTelemetryEnabled(value bool) *OdigosConfigurationSpecApplyConfiguration {
	b.TelemetryEnabled = &value
	return b
}

// WithOpenshiftEnabled sets the OpenshiftEnabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpenshiftEnabled field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithOpenshiftEnabled(value bool) *OdigosConfigurationSpecApplyConfiguration {
	b.OpenshiftEnabled = &value
	return b
}

// WithIgnoredNamespaces adds the given value to the IgnoredNamespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IgnoredNamespaces field.
func (b *OdigosConfigurationSpecApplyConfiguration) WithIgnoredNamespaces(values ...string) *OdigosConfigurationSpecApplyConfiguration {
	for i := range values {
		b.IgnoredNamespaces = append(b.IgnoredNamespaces, values[i])
	}
	return b
}

// WithIgnoredContainers adds the given value to the IgnoredContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IgnoredContainers field.
func (b *OdigosConfigurationSpecApplyConfiguration) WithIgnoredContainers(values ...string) *OdigosConfigurationSpecApplyConfiguration {
	for i := range values {
		b.IgnoredContainers = append(b.IgnoredContainers, values[i])
	}
	return b
}

// WithPsp sets the Psp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Psp field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithPsp(value bool) *OdigosConfigurationSpecApplyConfiguration {
	b.Psp = &value
	return b
}

// WithImagePrefix sets the ImagePrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePrefix field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithImagePrefix(value string) *OdigosConfigurationSpecApplyConfiguration {
	b.ImagePrefix = &value
	return b
}

// WithOdigletImage sets the OdigletImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OdigletImage field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithOdigletImage(value string) *OdigosConfigurationSpecApplyConfiguration {
	b.OdigletImage = &value
	return b
}

// WithInstrumentorImage sets the InstrumentorImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InstrumentorImage field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithInstrumentorImage(value string) *OdigosConfigurationSpecApplyConfiguration {
	b.InstrumentorImage = &value
	return b
}

// WithAutoscalerImage sets the AutoscalerImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AutoscalerImage field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithAutoscalerImage(value string) *OdigosConfigurationSpecApplyConfiguration {
	b.AutoscalerImage = &value
	return b
}

// WithCollectorGateway sets the CollectorGateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CollectorGateway field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithCollectorGateway(value *CollectorGatewayConfigurationApplyConfiguration) *OdigosConfigurationSpecApplyConfiguration {
	b.CollectorGateway = value
	return b
}

// WithGoAutoIncludeCodeAttributes sets the GoAutoIncludeCodeAttributes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GoAutoIncludeCodeAttributes field is set to the value of the last call.
func (b *OdigosConfigurationSpecApplyConfiguration) WithGoAutoIncludeCodeAttributes(value bool) *OdigosConfigurationSpecApplyConfiguration {
	b.GoAutoIncludeCodeAttributes = &value
	return b
}

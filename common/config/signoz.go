package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/odigos-io/odigos/common"
)

const (
	signozUrlKey = "SIGNOZ_URL"
)

type Signoz struct{}

func (s *Signoz) DestType() common.DestinationType {
	return common.SignozDestinationType
}

func (s *Signoz) ModifyConfig(dest ExporterConfigurer, currentConfig *Config) error {
	url, exists := dest.GetConfig()[signozUrlKey]
	if !exists {
		return errors.New("Signoz url not specified, gateway will not be configured for Signoz")
	}

	if strings.HasPrefix(url, "https://") {
		return errors.New("Signoz does not currently supports tls export, gateway will not be configured for Signoz")
	}

	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimSuffix(url, ":4317")
	signozExporterName := "otlp/signoz-" + dest.GetID()
	currentConfig.Exporters[signozExporterName] = GenericMap{
		"endpoint": fmt.Sprintf("%s:4317", url),
		"tls": GenericMap{
			"insecure": true,
		},
	}

	if isTracingEnabled(dest) {
		tracesPipelineName := "traces/signoz-" + dest.GetID()
		currentConfig.Service.Pipelines[tracesPipelineName] = Pipeline{
			Exporters: []string{signozExporterName},
		}
	}

	if isMetricsEnabled(dest) {
		metricsPipelineName := "metrics/signoz-" + dest.GetID()
		currentConfig.Service.Pipelines[metricsPipelineName] = Pipeline{
			Exporters: []string{signozExporterName},
		}
	}

	if isLoggingEnabled(dest) {
		logsPipelineName := "logs/signoz-" + dest.GetID()
		currentConfig.Service.Pipelines[logsPipelineName] = Pipeline{
			Exporters: []string{signozExporterName},
		}
	}

	return nil
}

package config

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/odigos-io/odigos/common"
)

const (
	otlpHttpAuthHeaderEndpointKey = "OTLP_HTTP_ENDPOINT_AUTH_HEADER_KEY"
	otlpHttpAuthHeaderKey         = "OTLP_HTTP_AUTH_HEADER_KEY"
)

type OTLPHttpAuthHeaderKey struct{}

func (g *OTLPHttpAuthHeaderKey) DestType() common.DestinationType {
	return common.OtlpHttpHeaderAuthDestinationType
}

func (g *OTLPHttpAuthHeaderKey) ModifyConfig(dest ExporterConfigurer, currentConfig *Config) error {

	url, exists := dest.GetConfig()[otlpHttpAuthHeaderEndpointKey]
	if !exists {
		return errors.New("OTLP http endpoint not specified, gateway will not be configured for otlp http")
	}

	parsedUrl, err := parseOtlpHttpAuthHeaderEndpoint(url)
	if err != nil {
		return errors.Join(err, errors.New("otlp http endpoint invalid, gateway will not be configured for otlp http"))
	}
	headerKey, exists := dest.GetConfig()[otlpHttpAuthHeaderKey]
	otlpHttpHeaderAuthExporterName := "otlphttpHeaderAuth/generic-" + dest.GetID()
	currentConfig.Exporters[otlpHttpHeaderAuthExporterName] = GenericMap{
		"endpoint": parsedUrl,
		"headers": GenericMap{
			"x-api-key": headerKey,
		},
	}
	currentConfig.Exporters[otlpHttpHeaderAuthExporterName] = otlpHttpHeaderAuthExporterName

	if isTracingEnabled(dest) {
		tracesPipelineName := "traces/otlphttpheaderkey-" + dest.GetID()
		currentConfig.Service.Pipelines[tracesPipelineName] = Pipeline{
			Exporters: []string{otlpHttpHeaderAuthExporterName},
		}
	}

	if isMetricsEnabled(dest) {
		metricsPipelineName := "metrics/otlphttpheaderkey-" + dest.GetID()
		currentConfig.Service.Pipelines[metricsPipelineName] = Pipeline{
			Exporters: []string{otlpHttpHeaderAuthExporterName},
		}
	}

	if isLoggingEnabled(dest) {
		logsPipelineName := "logs/otlphttpheaderkey-" + dest.GetID()
		currentConfig.Service.Pipelines[logsPipelineName] = Pipeline{
			Exporters: []string{otlpHttpHeaderAuthExporterName},
		}
	}

	return nil
}

func parseOtlpHttpAuthHeaderEndpoint(rawUrl string) (string, error) {
	noWhiteSpaces := strings.TrimSpace(rawUrl)
	parsedUrl, err := url.Parse(noWhiteSpaces)
	if err != nil {
		return "", fmt.Errorf("failed to parse otlp http endpoint: %w", err)
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return "", fmt.Errorf("invalid otlp http endpoint scheme: %s", parsedUrl.Scheme)
	}

	return noWhiteSpaces, nil
}

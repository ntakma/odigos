package config

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/odigos-io/odigos/common"
)

const (
	otlpHttpEndpointKey          = "OTLP_HTTP_ENDPOINT"
	otlpHttpBasicAuthUsernameKey = "OTLP_HTTP_BASIC_AUTH_USERNAME"
	otlpHttpBasicAuthPasswordKey = "OTLP_HTTP_BASIC_AUTH_PASSWORD"
	OtlpHttpAuthHeaderKey        = "OTLP_HTTP_AUTH_HEADER_KEY"
)

type OTLPHttp struct{}

func (g *OTLPHttp) DestType() common.DestinationType {
	return common.OtlpHttpDestinationType
}

func (g *OTLPHttp) ModifyConfig(dest ExporterConfigurer, currentConfig *Config) error {

	url, exists := dest.GetConfig()[otlpHttpEndpointKey]
	if !exists {
		return errors.New("OTLP http endpoint not specified, gateway will not be configured for otlp http")
	}

	parsedUrl, err := parseOtlpHttpEndpoint(url)
	if err != nil {
		return errors.Join(err, errors.New("otlp http endpoint invalid, gateway will not be configured for otlp http"))
	}
	//parsedUrlEndpoint
	exporterConf := GenericMap{
		"endpoint": parsedUrl,
	}
	basicAuthExtensionName, basicAuthExtensionConf, err := applyBasicAuth(dest)
	if err != nil {
		return errors.Join(err, errors.New("failed to apply basic auth to otlp http exporter"))
	}

	// add authenticator extension if use basic auth
	if basicAuthExtensionName != "" && basicAuthExtensionConf != nil {
		currentConfig.Extensions[basicAuthExtensionName] = *basicAuthExtensionConf
		currentConfig.Service.Extensions = append(currentConfig.Service.Extensions, basicAuthExtensionName)
		exporterConf["auth"] = GenericMap{
			"authenticator": basicAuthExtensionName,
		}
	}
	// add authentication with header key
	headerAuthExtentionName, headerAuthExtentionConf, err := applyHeaderAuth(dest)
	if err != nil {
		return errors.Join(err, errors.New("failed to apply header auth to otlp http exporter"))
	}
	if headerAuthExtentionName != "" && headerAuthExtentionConf != nil {
		currentConfig.Extensions[headerAuthExtentionName] = *headerAuthExtentionConf
		currentConfig.Service.Extensions = append(currentConfig.Service.Extensions, headerAuthExtentionName)
		exporterConf["headers"] = exporterConf
	}
	otlpHttpExporterName := "otlphttp/generic-" + dest.GetID()

	currentConfig.Exporters[otlpHttpExporterName] = exporterConf

	if isTracingEnabled(dest) {
		tracesPipelineName := "traces/otlphttp-" + dest.GetID()
		currentConfig.Service.Pipelines[tracesPipelineName] = Pipeline{
			Exporters: []string{otlpHttpExporterName},
		}
	}

	if isMetricsEnabled(dest) {
		metricsPipelineName := "metrics/otlphttp-" + dest.GetID()
		currentConfig.Service.Pipelines[metricsPipelineName] = Pipeline{
			Exporters: []string{otlpHttpExporterName},
		}
	}

	if isLoggingEnabled(dest) {
		logsPipelineName := "logs/otlphttp-" + dest.GetID()
		currentConfig.Service.Pipelines[logsPipelineName] = Pipeline{
			Exporters: []string{otlpHttpExporterName},
		}
	}

	return nil
}

func parseOtlpHttpEndpoint(rawUrl string) (string, error) {
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

func applyBasicAuth(dest ExporterConfigurer) (extensionName string, extensionConf *GenericMap, err error) {

	username := dest.GetConfig()[otlpHttpBasicAuthUsernameKey]
	if username == "" {
		return "", nil, nil
	}

	extensionName = "basicauth/otlphttp-" + dest.GetID()
	extensionConf = &GenericMap{
		"client_auth": GenericMap{
			"username": username,
			"password": fmt.Sprintf("${%s}", otlpHttpBasicAuthPasswordKey),
		},
	}

	return extensionName, extensionConf, nil
}
func applyHeaderAuth(dest ExporterConfigurer) (extensionName string, extensionConf *GenericMap, err error) {

	authHeaderKey := dest.GetConfig()[OtlpHttpAuthHeaderKey]
	if authHeaderKey == "" {
		return "", nil, nil
	}

	extensionName = "headerAuth/otlphttp-" + dest.GetID()
	extensionConf = &GenericMap{
		"header": GenericMap{
			"x-api-key": fmt.Sprintf("${%s}", OtlpHttpAuthHeaderKey),
		},
	}

	return extensionName, extensionConf, nil
}

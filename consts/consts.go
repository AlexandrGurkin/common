// Package consts contains constants for connections.
package consts

// SubsystemName is const string for metrics.
const (
	SubsystemName = "gph"
)

// RequestID is const string for correlation id header.
const (
	CorrelationID = `X-Correlation-ID`
)

// Log field's names
const (
	FieldModule        = `module`
	FieldAction        = `action`
	FieldCorrelationID = `correlation_id`
	FieldURI           = `api_url`
	FieldParams        = `params`
	FieldHttpCode      = `http_code`
	FieldDuration      = `duration`
)

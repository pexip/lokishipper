package lokishipper

import (
	"time"
	"net/url"

	"github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
)

// Provide the types for convenience and forward compatibility by reducing
// direct dependencies.
type HTTPClientConfig = config.HTTPClientConfig

// Config describes configuration for a HTTP pusher client.
type Config struct {
	URL       *url.URL
	BatchWait time.Duration
	BatchSize int

	Client HTTPClientConfig

	// The labels to add to any time series or alerts when communicating with loki
	ExternalLabels model.LabelSet
	Timeout        time.Duration
}

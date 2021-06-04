package lokishipper

import (
	"time"
	"net/http"
	"net/url"

	"github.com/prometheus/common/model"
)

// Config describes configuration for a HTTP pusher client.
type Config struct {
	URL       *url.URL
	BatchWait time.Duration
	BatchSize int

	client *http.Client

	// The labels to add to any time series or alerts when communicating with loki
	ExternalLabels model.LabelSet
	Timeout        time.Duration
}

package lokishipper

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/common/model"
)

// The MultiError type implements the error interface, and contains the
// Errors used to construct it.
type MultiError []error

// Returns a concatenated string of the contained errors
func (es MultiError) Error() string {
	var buf bytes.Buffer

	if len(es) > 1 {
		_, _ = fmt.Fprintf(&buf, "%d errors: ", len(es))
	}

	for i, err := range es {
		if i != 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(err.Error())
	}

	return buf.String()
}

// Add adds the error to the error list if it is not nil.
func (es *MultiError) Add(err error) {
	if err == nil {
		return
	}
	if merr, ok := err.(MultiError); ok {
		*es = append(*es, merr...)
	} else {
		*es = append(*es, err)
	}
}

// Err returns the error list as an error or nil if it is empty.
func (es MultiError) Err() error {
	if len(es) == 0 {
		return nil
	}
	return es
}

// MultiClient is client pushing to one or more loki instances.
type MultiClient []Client

// NewMulti creates a new client
func NewMulti(logger log.Logger, cfgs ...Config) (Client, error) {
	if len(cfgs) == 0 {
		return nil, errors.New("at least one client config should be provided")
	}

	clients := make([]Client, 0, len(cfgs))
	for _, cfg := range cfgs {
		client, err := New(cfg, logger)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return MultiClient(clients), nil
}

// Handle Implements api.EntryHandler
func (m MultiClient) Handle(labels model.LabelSet, time time.Time, entry string) error {
	var result MultiError
	for _, client := range m {
		if err := client.Handle(labels, time, entry); err != nil {
			result.Add(err)
		}
	}
	return result.Err()
}

// Stop implements Client
func (m MultiClient) Stop() {
	for _, c := range m {
		c.Stop()
	}
}

package sdxx

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicSdxXAPI provides the sdxX RPC service that can be
// use publicly without security implications.
type PublicSdxXAPI struct {
	t        *SdxX
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicSdxXAPI create a new RPC sdxX service.
func NewPublicSdxXAPI(t *SdxX) *PublicSdxXAPI {
	api := &PublicSdxXAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the SdxX sub-protocol version.
func (api *PublicSdxXAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}

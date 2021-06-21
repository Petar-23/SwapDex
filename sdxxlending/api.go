package sdxxlending

import (
	"context"
	"errors"
	"sync"
	"time"
)

// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicSdxXLendingAPI provides the sdxX RPC service that can be
// use publicly without security implications.
type PublicSdxXLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicSdxXLendingAPI create a new RPC sdxX service.
func NewPublicSdxXLendingAPI(t *Lending) *PublicSdxXLendingAPI {
	api := &PublicSdxXLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicSdxXLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}

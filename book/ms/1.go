package ms

import (
	"context"
	"fmt"
	"sync"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/pkg/errors"
)

// FIND Problem in Code
var (
	cachedProvider      *oidc.Provider
	cachedProviderMutex = &sync.RWMutex{}
)

func getProviderFromCache() (*oidc.Provider, bool) {
	cachedProviderMutex.RLock()
	defer cachedProviderMutex.RUnlock()
	fmt.Println("Getting Provider from Cache")
	if cachedProvider == nil {
		return nil, false
	}
	return cachedProvider, true
}

func GetOIDCIssuerProvider() (*oidc.Provider, error) {
	// fast path: read from cache
	if cached, ok := getProviderFromCache(); ok {
		return cached, nil
	}

	// slow path: construct from remote
	cachedProviderMutex.Lock()
	defer cachedProviderMutex.Unlock()

	// check again after acquiring the lock to avoid race condition
	if cached, ok := getProviderFromCache(); ok {
		return cached, nil
	}

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://my-oidc-provider/.openid-connect")
	if err != nil {
		// failed in this attempt, let other attempts retry
		return nil, errors.Wrap(err, "failed to create provider for azure")
	}

	cachedProvider = provider

	return provider, nil
}

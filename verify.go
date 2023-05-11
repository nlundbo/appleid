package appleid

import (
	"context"
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

const (
	appleKeysURL = "https://appleid.apple.com/auth/keys"
)

type Verifier interface {
	AutoRefresh(ctx context.Context) error
	Verify(ctx context.Context, idToken string) error
}

func New() Verifier {
	return &verifier{}
}

type verifier struct {
	keySet *jwk.Cache
}

func (v *verifier) AutoRefresh(ctx context.Context) error {
	c := jwk.NewCache(ctx)

	c.Register(appleKeysURL, jwk.WithMinRefreshInterval(15*time.Minute))

	// Refresh the JWKS once before getting into the main loop.
	// This allows you to check if the JWKS is available before we start
	// a long-running program
	set, err := c.Refresh(ctx, appleKeysURL)
	if err != nil {
		return err
	}
	fmt.Println(set)
	v.keySet = c
	return nil
}

func (v *verifier) Verify(ctx context.Context, idToken string) error {
	set, err := v.keySet.Get(ctx, appleKeysURL)
	if err != nil {
		return err
	}
	token, err := jwt.Parse([]byte(idToken), jwt.WithKeySet(set) /*jwt.WithAudience("your clientID com.something.somethingapp")*/)
	if err != nil {
		return err
	}

	fmt.Println(token)
	return nil
}

package client

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWaitForRewardIssuedRejectsInvalidDurations(t *testing.T) {
	client := &Client{}

	_, err := client.WaitForRewardIssued(t.Context(), 1, 0, time.Second)
	if err == nil || !errors.Is(err, ErrInvalidWaitTimeout) {
		t.Fatalf("expected timeout validation error, got %v", err)
	}

	_, err = client.WaitForRewardIssued(t.Context(), 1, time.Second, 0)
	if err == nil || !errors.Is(err, ErrInvalidWaitInterval) {
		t.Fatalf("expected interval validation error, got %v", err)
	}
}

func TestExtractClaimAbilityAllowsUnavailableWithoutReason(t *testing.T) {
	resp := &GraphQLResponse[OctoplusOfferBySlugData]{
		Data: OctoplusOfferBySlugData{
			OctoplusOffer: operationsOctoplusOffer("caffe-nero", false, ""),
		},
	}

	canClaim, reason, err := ExtractClaimAbility(resp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if canClaim {
		t.Fatalf("expected canClaim=false")
	}
	if reason != "" {
		t.Fatalf("expected empty reason, got %q", reason)
	}
}

func TestEnsureFreshTokenSkipsRefreshForFreshToken(t *testing.T) {
	client := &Client{ExpiresAt: time.Now().Add(30 * time.Minute)}

	if err := client.EnsureFreshToken(t.Context(), 5*time.Minute); err != nil {
		t.Fatalf("expected no refresh for fresh token, got %v", err)
	}
}

func TestWaitForRewardIssuedReturnsParentCancelCause(t *testing.T) {
	ctx, cancel := context.WithCancelCause(t.Context())
	cause := errors.New("manual stop")
	cancel(cause)

	client := &Client{ExpiresAt: time.Now().Add(30 * time.Minute)}
	_, err := client.WaitForRewardIssued(ctx, 123, 5*time.Second, 10*time.Millisecond)
	if err == nil {
		t.Fatal("expected cancellation error, got nil")
	}
	if !errors.Is(err, cause) {
		t.Fatalf("expected error to wrap cancel cause, got %v", err)
	}
}

func operationsOctoplusOffer(slug string, canClaim bool, reason string) OctoplusOffer {
	return OctoplusOffer{
		Slug: slug,
		ClaimAbility: ClaimAbility{
			CanClaimOffer:     canClaim,
			CannotClaimReason: reason,
		},
	}
}

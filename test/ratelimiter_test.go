package test

import (
	ratelimiter "RateLimiter/RateLimiter"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tb := ratelimiter.NewTokenBucket(5, 2) // 5 tokens with a refill rate of 2 tokens per second

	// Initially, the bucket should have 5 tokens.
	if tb.Tokens != 5 {
		t.Errorf("Expected initial token count to be 5, but got %d", tb.Tokens)
	}

	// Check that the token bucket allows the first 5 requests.
	for i := 0; i < 5; i++ {
		if !tb.Take() {
			t.Errorf("Request %d: Expected to be allowed, but denied", i+1)
		}
	}

	// Check that the token bucket denies the next 5 requests due to the rate limit.
	for i := 5; i < 10; i++ {
		if tb.Take() {
			t.Errorf("Request %d: Expected to be denied, but allowed", i+1)
		}
	}

	// Wait for the bucket to refill.
	time.Sleep(3 * time.Second)

	// After waiting, the token bucket should have refilled to 5 tokens.
	if tb.Tokens != 5 {
		t.Errorf("Expected token count after refill to be 5, but got %d", tb.Tokens)
	}

	// Check that the token bucket allows the next 5 requests.
	for i := 0; i < 5; i++ {
		if !tb.Take() {
			t.Errorf("Request %d: Expected to be allowed, but denied", i+1)
		}
	}
}

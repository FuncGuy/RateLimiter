package ratelimiter

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	Tokens     int
	capacity   int
	refillRate int
	mu         sync.Mutex
	lastRefill time.Time
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	tb := &TokenBucket{
		Tokens:     capacity,
		capacity:   capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
	return tb
}

func (tb *TokenBucket) refill() {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	tokensToAdd := int(elapsed.Seconds()) * tb.refillRate
	tb.Tokens += int(math.Min(float64(tokensToAdd), float64(tb.capacity-tb.Tokens)))
	tb.lastRefill = now
}

func (tb *TokenBucket) Take() bool {
	tb.refill()

	tb.mu.Lock()
	defer tb.mu.Unlock()

	if tb.Tokens > 0 {
		tb.Tokens--
		return true
	}

	return false
}

package rate_limiter

import (
	"log"
	"net/http"
	"time"
)

type RateLimiter struct {
	priority   int
	identifier func(r *http.Request) string
	config     RateLimitingConfig
	info       *log.Logger
	warn       *log.Logger
	error      *log.Logger

	tokenBucketLastRequest     map[string]*time.Time // time of last request of the client
	tokenBucketAvailableTokens map[string]*int       // number of tokens of the client
}

type RateLimitingConfig struct {
	requestNumber int
	interval      time.Duration
}

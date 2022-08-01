package rate_limiter

import (
	"log"
	"net/http"
	"time"
)

func (rl *RateLimiter) SetInfoLogger(logger *log.Logger) {

	rl.info = logger
}
func (rl *RateLimiter) SetWarnLogger(logger *log.Logger) {
	rl.warn = logger
}
func (rl *RateLimiter) SetErrorLogger(logger *log.Logger) {
	rl.error = logger
}

func (rl *RateLimiter) Priority() int {
	return rl.priority
}
func (rl *RateLimiter) SetPriority(priority int) {
	rl.priority = priority

}

func (rl *RateLimiter) SetConfig(config RateLimitingConfig) {

	rl.tokenBucketLastRequest = map[string]*time.Time{}
	rl.tokenBucketAvailableTokens = map[string]*int{}

	rl.config = config

}

func (rl *RateLimiter) Name() string {
	return "rate-limiter"

}

func (rl *RateLimiter) SetRequestIdentifierFunc(identifier func(request *http.Request) string) {

	rl.identifier = identifier

}

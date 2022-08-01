package rate_limiter

import (
	"net/http"
	"time"
)

func (rl *RateLimiter) Handle(r *http.Request, w http.ResponseWriter) (ok bool, err error) {

	id := rl.identifier(r)
	config := rl.config
	timeOfLastRequest := rl.tokenBucketLastRequest[id]
	tokens := rl.tokenBucketAvailableTokens[id]

	updateAvailableTokens := func() {

		/* Adds tokens proportionally to elapsed time since last request,
		e.g. with a configured interval of 5 seconds, if 3 seconds elapsed since last request,
		3/5 (60%) of max-token number will be added to counter.

		The casting to int works as math.Floor()
		*/

		newNumberOfTokens := *tokens + int(timeOfLastRequest.Sub(time.Now()).Seconds()/config.interval.Seconds())

		if newNumberOfTokens > config.requestNumber {

			*tokens = config.requestNumber

		} else {

			*tokens = newNumberOfTokens

		}

	}

	if config.requestNumber <= 0 {

		return

	}

	if timeOfLastRequest.IsZero() {

		*tokens = config.requestNumber

		*timeOfLastRequest = time.Now()
		*tokens--

	} else {

		updateAvailableTokens()

		if *tokens > 1 {

			*timeOfLastRequest = time.Now()
			*tokens--

			ok = true
			return

		}

	}

	rl.info.Printf("%s rate limited: out of tokens", id)

	w.WriteHeader(http.StatusTooManyRequests)
	w.Write([]byte("Too many requests"))
	return

}

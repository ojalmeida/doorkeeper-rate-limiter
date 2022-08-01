<img src="https://github.com/ojalmeida/doorkeeper-core/blob/main/logo.png?raw=true" alt="drawing" width="400"/>
<br>
<br>

**doorkeeper-rate-limiter** middleware limits the per-client incoming requests making use of Token Bucket algorithm.

It can be used to limit client requests by geolocation, IP address, Session_ID, or any other method of identification.

## Getting started

```bash
go get -u github.com/ojalmeida/doorkeeper-rate-limiter
```

```go
package main

import (
	core "github.com/ojalmeida/doorkeeper-core"
	rl "github.com/ojalmeida/doorkeeper-rate-limiter"
	"os"
)

var rateLimiter = rl.RateLimiter{}

func main() {

	file, err := os.Open("/tmp/config.json")

	if err != nil {
		panic(err)

	}

	core.SetConfigFile(file)
	core.BindMiddleware("^/api/v1/", &rateLimiter)

	// 5 reqs/s by client
	rateLimiter.SetConfig(RateLimitingConfig{
		requestNumber: 5,
		interval:      time.Second,
	})
	
	// identify clients by its ip address
	rateLimiter.SetRequestIdentifierFunc(func(request *http.Request) string {

		return strings.Split(request.RemoteAddr, ":")[0]

	})

	core.Start()

}
```

Any doubt about configuration files refer to [doorkeeper-core](https://github.com/ojalmeida/doorkeeper-core)

### Starting application

```bash

go run *.go

```


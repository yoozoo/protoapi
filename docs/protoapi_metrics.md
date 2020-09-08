# protoapi metrics

## Usage

Import middleware and use in echo.

example:
```go
package main

import (
	"github.com/labstack/echo"
	"github.com/yoozoo/protoapi/protoapigo/metrics"
	"code/generated/by/protoapi/demo"
)

type Service struct{}

var _ demo.DemoService = &Service{}

func main() {
    e := echo.New()

    service := &Service{}
	calsvr.RegisterCalcService(e, c)

    // use metrics middleware
    m := metrics.MetricsFunc()
    e.Use(m)

	e.Logger.Fatal(e.Start(":8080"))
}
```

## Metrics endpoint

By default, echo server exports metrics under the `/metrics` path on its client port.

The metrics can be fetched with `curl`:

```sh
$ curl -L http://localhost:8080/metrics
```
## More Options
Function `metrics.MetricsFunc()` can have `Option` parameters.
```go
package metrics

func Registry(r prometheus.Registerer) Option {}

func MetricsPath(v string) Option {}

func Namespace(v string) Option {}

func Subsystem(v string) Option {}
```

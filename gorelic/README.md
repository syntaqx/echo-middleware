# Gorelic

NewRelic middleware for [yvasiyarov/gorelic](https://github.com/yvasiyarov/gorelic)

## Installation

```shell
go get github.com/syntaqx/echo-middleware/gorelic
```

## Usage

```go
package main

import (
    "github.com/labstack/echo"
    "github.com/syntaqx/echo-middleware/gorelic"
)

func main() {
    e := echo.New()

    // Attach middleware
    gorelic.InitNewRelicAgent("YOUR_LICENSE_KEY", "YOUR_APPLICATION_NAME", true)
    e.Use(gorelic.Handler())

    e.Run(":8080")
}
```

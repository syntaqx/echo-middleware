# RemoteAddr Middleware

RemoteAddr provides middleware for sanitizing the `RemoteAddr` based on other
possible headers.

## Install

```sh
go get github.com/syntaqx/echo-middleware # or, you could specify only remoteaddr
```

## Getting Started

Here's a simple example application using remoteaddr and Echo. Save this file as
`main.go`

```go
package main

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/syntaqx/echo-middleware/remoteaddr"
)

func main() {
    e := echo.New()

    e.Use(remoteaddr.New().Handler)

    e.Get("/", func(c *echo.Context) error {
        return c.HTML(http.StatusOK, c.Request().RemoteAddr)
    })

    e.Run(":8080")
}

```

Then, run your server:

```sh
go run main.go
```

The server now runs on `localhost:8080`

```sh
$ curl -D - -H 'Origin: http://localhost' http://localhost:8080/
```

You should recieve a response giving you back your IP address (which, in this
case will either be `127.0.0.1` or `::1`). Now you can use that value safely,
knowing it's being looked after!

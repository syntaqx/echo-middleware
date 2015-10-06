# Session

Middleware support for echo by gorilla/session

## Installation

```shell
go get github.com/syntaqx/echo-middleware/session
```

## Usage

```go
package main

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/syntaqx/echo-middleware/session"
)

func index(c *echo.Context) error {
    session := session.Default(c)

    var count int
    v := session.Get("count")

    if v == nil {
        count = 0
    } else {
        count = v.(int)
        count += 1
    }

    session.Set("count", count)
    session.Save()

    data := struct {
        Visit int
    }{
        Visit: count,
    }

    return c.JSON(http.StatusOK, data)
}

func main() {
    store := session.NewCookieStore([]byte("secret-key"))
    // store, err := session.NewRedisStore(32, "tcp", "localhost:6379", "", []byte("secret-key"))
    // if err != nil {
    //     panic(err)
    // }

    e := echo.New()

    // Attach middleware
    e.Use(session.Sessions("ESESSION", store))

    // Routes
    e.Get("/", index)

    e.Run(":8080")
}
```

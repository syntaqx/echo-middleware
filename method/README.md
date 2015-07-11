# Method

Middleware/handler for handling http method overrides. This checks for the
`X-HTTP-Method-Override` header and uses it if the original reqeust method is
`POST`. `GET/HEAD` methods shouldn't be overriden, hence they can't be
overriden.

This is useful for REST APIs and services making use of many HTTP verbs, and
when the HTTP clients don't support all of them.

## Installation

```shell
go get github.com/syntaqx/echo-middleware/method
``

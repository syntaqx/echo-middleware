// Method implements the ability to override HTTP methods
// This package was created from github.com/codegangsta/martini-contrib/method
// and modified to work with the echo framework
package method

import (
	"errors"

	"github.com/labstack/echo"
)

// HeaderHTTPMethodOVerride is a common HTTP header used to override the HTTP
// method
const HeaderHTTPMethodOverride = "X-HTTP-Method-Override"

// ParamHTTPMethodOverride is a common used HTTML form parameter used to
// override the HTTP method
const ParamHTTPMethodOverride = "_method"

// HTTP methods overridden by this package
var httpMethods = []string{"PUT", "PATCH", "DELETE"}

// ErrInvalidOverrideMethod is returned when
// an invalid http method was given to OverrideRequestMethod.
var ErrInvalidOverrideMethod = errors.New("invalid override method")

// isValidOverrideMethod determines if the method is a valid override method
func isValidOverrideMethod(method string) bool {
	for _, m := range httpMethods {
		if m == method {
			return true
		}
	}
	return false
}

// Override checks for the X-HTTP-Method-Override header
// or the HTML for parameter, `_method`
// and uses (if valid) the http method instead of
// Request.Method.
// This is especially useful for http clients
// that don't support many http verbs.
// It isn't secure to override e.g a GET to a POST,
// so only Request.Method which are POSTs are considered.
func Override() echo.HandlerFunc {
	return func(c *echo.Context) error {
		if c.Request().Method == "POST" {
			m := c.Form(ParamHTTPMethodOverride)
			if isValidOverrideMethod(m) {
				OverrideRequestMethod(c, m)
			}
			m = c.Request().Header.Get(HeaderHTTPMethodOverride)
			if isValidOverrideMethod(m) {
				c.Request().Method = m
			}
		}
		return nil
	}
}

// OverrideRequestMethod overrides the http
// request's method with the specified method.
func OverrideRequestMethod(c *echo.Context, method string) error {
	if !isValidOverrideMethod(method) {
		return ErrInvalidOverrideMethod
	}
	c.Request().Header.Set(HeaderHTTPMethodOverride, method)
	return nil
}

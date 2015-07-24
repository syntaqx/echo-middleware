package gorelic

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
)

var agent *gorelic.Agent

// InitNewRelicAgent initializes a new gorelic agent for usage in Handler
func InitNewRelicAgent(license string, appname string, verbose bool) error {
	if license == "" {
		return fmt.Errorf("Please specify a NewRelic license")
	}

	agent = gorelic.NewAgent()

	agent.NewrelicLicense = license
	agent.NewrelicName = appname
	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose

	agent.Run()

	return nil
}

// Gorelic returns a middleware function that attaches a gorelic agent
func Handler() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			startTime := time.Now()
			err := h(c)
			agent.HTTPTimer.UpdateSince(startTime)
			return err
		}
	}
}

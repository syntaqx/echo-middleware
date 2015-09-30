package gorelic

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
)

var agent *gorelic.Agent

// Gorelic returns a middleware function that attaches a gorelic agent
func Handler() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			startTime := time.Now()
			err := h(c)

			if agent != nil {
				agent.HTTPTimer.UpdateSince(startTime)
			}

			return err
		}
	}
}

// InitNewRelicAgent initializes a new gorelic agent for usage in Handler
func InitNewRelicAgent(license string, appname string, verbose bool) (*gorelic.Agent, error) {
	agent = gorelic.NewAgent()
	if license == "" {
		return agent, fmt.Errorf("Please specify a NewRelic license")
	}

	agent.NewrelicLicense = license
	agent.NewrelicName = appname
	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose

	agent.Run()

	return agent, nil
}

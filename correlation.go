package cfecho

import (
	"context"

	"github.com/labstack/echo"
	"github.com/m4rw3r/uuid"
)

var (
	HeaderVCapRequestID = "X-Vcap-Request-Id"
	HeaderCorrelationID = "X-Correlation-Id"
	KeyCorrelationID    = "correlationid"
)

func CorrelationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			id := req.Header.Get(HeaderCorrelationID)
			if id == "" {
				id = req.Header.Get(HeaderVCapRequestID)
				if id == "" { // Generate if not there
					generated, _ := uuid.V4()
					id = generated.String()
				}
			}
			c.Response().Header().Set(HeaderCorrelationID, id)
			newreq := req.WithContext(context.WithValue(req.Context(), KeyCorrelationID, id))
			c.SetRequest(newreq)
			return next(c)
		}
	}
}

package access

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sendya/pkg/log"
)

// ignored: SkipPaths
func New(ignored ...string) gin.HandlerFunc {
	// cached ignored request path
	skipPaths := make(map[string]bool, len(ignored))
	for _, path := range ignored {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		t := time.Now()
		path := c.Request.RequestURI

		c.Next()

		if _, ok := skipPaths[path]; !ok {
			latency := time.Since(t)

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					log.WithCtx(c.Request.Context()).Error(e)
				}
			} else {
				fields := []zap.Field{
					log.Int("status", c.Writer.Status()),
					log.String("method", c.Request.Method),
					log.String("path", path),
					log.String("ip", c.ClientIP()),
					log.String("ua", c.Request.UserAgent()),
					log.Duration("latency", latency),
				}
				log.WithCtx(c.Request.Context()).Info("access", fields...)
			}
		}

	}
}

func NewRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.WithCtx(c.Request.Context()).
						Error(c.Request.URL.Path,
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
						)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					log.WithCtx(c.Request.Context()).
						Error("[Recovery from panic]",
							log.Time("time", time.Now()),
							log.Any("error", err),
							log.String("request", string(httpRequest)),
							log.String("stack", string(debug.Stack())),
						)
				} else {
					log.WithCtx(c.Request.Context()).
						Error("[Recovery from panic]",
							log.Time("time", time.Now()),
							log.Any("error", err),
							log.String("request", string(httpRequest)),
						)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}

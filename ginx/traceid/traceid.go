package traceid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/sendya/pkg/log"
)

const traceIDKey = "traceId"

func New() gin.HandlerFunc {

	return func(c *gin.Context) {
		// append traceId to log fields
		ctx, _ := log.WithFields(c.Request.Context(), log.String(traceIDKey, uuid.New().String()))
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

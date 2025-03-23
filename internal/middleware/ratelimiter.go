package middleware

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth/v7"
	"github.com/gin-gonic/gin"
)

/*
-----------------------------------------------------------------------------------------

	  Function Name : RateLimitMiddleware
	  Purpose       : RateLimitMiddleware func is used to checek the rate limit
	   -----------------------------------------------------------------------------------------


		Returns:
	   -----------------------------------------------------------------------------------------
		gin.HandlerFunc

		Success Response:
		-----------------------------------------------------------------------------------------
		In case of successful execution, this function will return Expected success result.



		Author        : LOGESHKUMAR P
		Created Date  : 22-03-2025

-----------------------------------------------------------------------------------------
*/
func RateLimitMiddleware() gin.HandlerFunc {
	limiter := tollbooth.NewLimiter(60, nil)

	limiter.SetTokenBucketExpirationTTL(1 * time.Minute)

	limiter.SetMessage("Too many requests. Please try again later.")
	limiter.SetMessageContentType("application/json")

	limiter.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error": "Rate limit exceeded. Please wait."}`))
	})

	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
		if httpError != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":     "Rate limit exceeded",
				"max":       limiter.GetMax(),
				"remaining": limiter.GetMax() - 1,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

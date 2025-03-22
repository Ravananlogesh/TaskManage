package middleware

import (
	"time"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth_gin"
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

	return tollbooth_gin.LimitHandler(limiter)

}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"

	"backend/middlewares"
	"backend/routes"
)

var cb *gobreaker.CircuitBreaker

func init() {
	cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "backend-circuit-breaker",
		MaxRequests: 3,
		Interval:    10 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 10 && failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			fmt.Printf("Circuit breaker: %s -> %s\n", from, to)
		},
	})
}

func circuitBreakerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := cb.Execute(func() (interface{}, error) {
			c.Next()
			return nil, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"error": "circuit breaker open",
			})
			return
		}
	}
}

func rateLimitMiddleware(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.AbortWithStatusJSON(httpError.StatusCode, gin.H{
				"error": httpError.Message,
			})
			return
		}
		c.Next()
	}
}

func main() {
	fmt.Println("Starting...")

	r := gin.Default()

	lmt := tollbooth.NewLimiter(5, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Minute,
	})
	lmt.SetIPLookups([]string{"X-Forwarded-For", "X-Real-IP", "RemoteAddr"})

	r.Use(rateLimitMiddleware(lmt))
	r.Use(circuitBreakerMiddleware())
	r.Use(middlewares.Logging())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, You are at root!",
		})
	})

	r.GET("/dominican", func(c *gin.Context) {
		fmt.Println(" I no black, I dominican papi!")
	}, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You dominican papi!",
		})
	})

	routes.UserRoute(r)

	r.GET("/magnacarta", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"lines": []string{
				"John, by the grace of God King of England, Lord of Ireland, and Duke of Normandy.",
				"To all his faithful and loving people, greeting.",
				"Know that we, of our pure and earnest heart, for the honour of God and the amendment of our kingdom...",
			},
		})
	})

	r.Run(":3333")
}
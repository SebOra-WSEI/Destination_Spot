package route

import (
	 "github.com/SebOra-WSEI/Destination_spot/core/internal/handler/auth"
	 "github.com/SebOra-WSEI/Destination_spot/core/internal/handler/reservation"
	 "github.com/SebOra-WSEI/Destination_spot/core/internal/handler/spot"
	 "github.com/SebOra-WSEI/Destination_spot/core/internal/handler/user"
	 "github.com/gin-gonic/gin"
	 "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Init(r *gin.Engine) {
	 r.Use(getCors())
	 r.Use(otelgin.Middleware("my-server"))

	 // Auth service endpoints
	 r.POST("/sign-in", auth.SignIn)
	 r.POST("/sign-up", auth.SignUp)
	 r.PUT("/reset-password/:id", auth.ResetPassword)
	 r.PUT("/access-control/:id", auth.AccessControl)

	 // Users endpoints
	 r.GET("/users", user.GetAll)
	 r.GET("/users/:id", user.GetById)
	 r.DELETE("/users/:id", user.Delete)

	 // Spots endpoints
	 r.GET("/spots", spot.GetAll)
	 r.POST("/spots", spot.Create)
	 r.GET("/spots/:id", spot.GetById)
	 r.DELETE("/spots/:id", spot.Delete)

	 // Reservations endpoints
	 r.GET("/reservations", reservation.GetAll)
	 r.POST("/reservations", reservation.Create)
	 r.GET("/reservations/:id", reservation.GetById)
	 r.DELETE("/reservations/:id", reservation.Delete)
	 r.PUT("/reservations/:id", reservation.Update)
}

func getCors() gin.HandlerFunc {
	 return func(c *gin.Context) {
		  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		  c.Writer.Header().Set(
				"Access-Control-Allow-Headers",
				"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		  )
		  c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		  if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
		  }

		  c.Next()
	 }
}

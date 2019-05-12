package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func BindRoutes(r *gin.Engine) gin.IRoutes {
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// TODO Replace with JWT
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"mobile": "a_nil_trap", // user:mobile password:a_nil_trap
		"web":    "a_nil_trap", // user:web password:a_nil_trap
	}))

	authorized.POST("register", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

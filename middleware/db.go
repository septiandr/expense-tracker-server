// middleware/db.go
package middleware

import (
	"expense-tracker/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", config.DB)
		c.Next()
	}
}

// Helper ambil DB dari context
func GetDB(c *gin.Context) *gorm.DB {
	return c.MustGet("db").(*gorm.DB)
}

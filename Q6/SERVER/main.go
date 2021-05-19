package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	outfile, _ = os.Create("server.log") // update path for your needs
	l          = log.New(outfile, "", 0)
)

type Data struct {
	Counter string `json:"counter"`
}

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var json Data
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		xrandom := c.Request.Header.Get("X-RANDOM")

		c.JSON(201, gin.H{"counter": json.Counter, "X-RANDOM": xrandom})
		l.Printf("[%s] %s %s %s", time.Now().Format(time.RFC1123), c.Request.Method, c.ClientIP(), gin.H{"counter": json.Counter, "X-RANDOM": xrandom})
	})
	r.Run(":9999")
}

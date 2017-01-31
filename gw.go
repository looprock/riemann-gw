package main

import (
  "gopkg.in/gin-gonic/gin.v1"
  "strconv"
  "github.com/amir/raidman"
)

// Binding from JSON
type Metrics struct {
    State string `form:"state" json:"state" binding:"required"`
    Host string `form:"host" json:"host" binding:"required"`
    Service string `form:"service" json:"service" binding:"required"`
    Metric string `form:"metric" json:"metric" binding:"required"`
    // Ttl int `form:"ttl" json:"ttl" binding:"required"`
}

func main() {
    router := gin.Default()

    // forward post data to riemann
    router.POST("/riemann", func(c *gin.Context) {
      var json Metrics
      riemann, err := raidman.Dial("tcp", "localhost:5555")
      if err != nil {
              panic(err)
      }
      if c.BindJSON(&json) == nil {
        var m, _ = strconv.ParseFloat(json.Metric,10)
        var event = &raidman.Event{
          State:   json.State,
          Host:    json.Host,
          Service: json.Service,
          Metric:  int(m),
          Ttl:     10,
        }
        err = riemann.Send(event)
        if err != nil {
          c.JSON(500, gin.H{"status": "error", "message": err})
        } else {
          c.JSON(200, gin.H{"status": "success", "message": "OK"})
        }
      }
      riemann.Close()
    })

    router.Run(":18001")
}

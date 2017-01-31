package main

import (
  "gopkg.in/gin-gonic/gin.v1"
  "strconv"
  "github.com/amir/raidman"
  "flag"
  "fmt"
)


// Binding from JSON
type Metrics struct {
    State string `form:"state" json:"state" binding:"required"`
    Host string `form:"host" json:"host" binding:"required"`
    Service string `form:"service" json:"service" binding:"required"`
    Metric string `form:"metric" json:"metric" binding:"required"`
    Ttl string `form:"ttl" json:"ttl,omitempty"`
    Tags []string `form:"tags" json:"tags,omitempty"`
}

func main() {
    riemannserver := flag.String("riemannserver", "localhost", "Riemann server IP/hostname")
    riemannport := flag.String("riemannport", "5555", "Riemann server port")
    listenaddr := flag.String("l", "", "IP to listen on")
    listenport := flag.String("p", "18001", "Port to listen on")
    flag.Parse()
    
    var riemannhost string
    riemannhost = *riemannserver + ":" + *riemannport
    fmt.Printf("Reporting to Riemann host: %s\n", riemannhost)
    var listenhost string
    listenhost = *listenaddr + ":" + *listenport
    fmt.Printf("Listening at: %s\n", listenhost)
    router := gin.Default()

    // forward post data to riemann
    router.POST("/riemann", func(c *gin.Context) {
      var json Metrics
      riemann, err := raidman.Dial("tcp", riemannhost)
      if err != nil {
              panic(err)
      }
      if c.BindJSON(&json) == nil {
        var t float64
        if json.Ttl != "" {
           t, _ = strconv.ParseFloat(json.Ttl,10)
         } else {
           t, _ = strconv.ParseFloat("10",10)
         }
        var m, _ = strconv.ParseFloat(json.Metric, 64)
        var event = &raidman.Event{
          State:   json.State,
          Host:    json.Host,
          Service: json.Service,
          Metric:  m,
          Ttl:     float32(t),
          Tags:    json.Tags,
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

    router.Run(listenhost)
}

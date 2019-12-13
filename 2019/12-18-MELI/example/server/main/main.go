package main

import (
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var counter uint64
var cMux = &sync.Mutex{}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	rand.Seed(42)

	r.GET("/ping", pingHandler)

	ginpprof.Wrap(r)

	return r
}

func main() {
	// ejecutar la aplicación por 15 segundos
	runtimeTicker := time.NewTicker(15 * time.Second)

	go func() {
		for {
			select {
			case <-runtimeTicker.C:
				log.WithFields(log.Fields{
					"time": time.Now().String(),
					"counter": counter,
				}).Info("closing server")
				os.Exit(0)
			}
		}
	}()

	r := setupRouter()
	r.Run(":5050")
}

func pingHandler(c *gin.Context) {
	// logs
	log.WithFields(log.Fields{
		"time": time.Now().String(),
	}).Info("ping received")
	log.WithFields(log.Fields{
		"context": c,
	}).Info("request received")
	log.WithFields(log.Fields{
		"context": c,
	}).Info("request received")

	// counter lock
	cMux.Lock()
	defer cMux.Unlock()
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Nanosecond)
	counter = counter + 1
	
	// response
	c.String(http.StatusOK, "pong")
}
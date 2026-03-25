package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"vpn-manager/internal/db"
	"vpn-manager/internal/metrics"
	"vpn-manager/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	metrics.Init()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	serverPublicKey := "5ABgAyy7PLlR+dw971B2mwP4eiKIgdfKd+rfW7dmIlY="
	serverIP := "127.0.0.1"

	// CREATE DEVICE
	router.POST("/devices", func(c *gin.Context) {

		start := time.Now()

		metrics.RequestCounter.WithLabelValues("POST", "/devices").Inc()

		device, err := service.CreateDevice(database, serverPublicKey, serverIP)
		if err != nil {
			metrics.ErrorCounter.Inc()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		metrics.RequestDuration.WithLabelValues("/devices").Observe(time.Since(start).Seconds())

		c.JSON(http.StatusOK, device)
	})

	// GET DEVICES
	router.GET("/devices", func(c *gin.Context) {

		start := time.Now()

		metrics.RequestCounter.WithLabelValues("GET", "/devices").Inc()

		devices, err := db.GetDevices(database)
		if err != nil {
			metrics.ErrorCounter.Inc()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		metrics.RequestDuration.WithLabelValues("/devices").Observe(time.Since(start).Seconds())

		c.JSON(http.StatusOK, devices)
	})

	// DELETE /devices/:id
	router.DELETE("/devices/:id", func(c *gin.Context) {

		start := time.Now()
		metrics.RequestCounter.WithLabelValues("DELETE", "/devices").Inc()

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			metrics.ErrorCounter.Inc()
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		err = service.DeleteDevice(database, id)
		if err != nil {

			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
				return
			}

			metrics.ErrorCounter.Inc()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		metrics.RequestDuration.WithLabelValues("/devices").Observe(time.Since(start).Seconds())

		c.JSON(http.StatusOK, gin.H{"message": "device deleted"})
	})

	router.Run(":8080")
}

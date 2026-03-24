package main

import (
	"net/http"
	"strconv"

	"vpn-manager/internal/db"
	"vpn-manager/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	database, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	serverPublicKey := "5ABgAyy7PLlR+dw971B2mwP4eiKIgdfKd+rfW7dmIlY="
	serverIP := "127.0.0.1"

	// CREATE DEVICE
	router.POST("/devices", func(c *gin.Context) {

		device, err := service.CreateDevice(database, serverPublicKey, serverIP)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, device)
	})

	// GET DEVICES
	router.GET("/devices", func(c *gin.Context) {

		devices, err := db.GetDevices(database)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, devices)
	})

	// DELETE DEVICE
	router.DELETE("/devices/:id", func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		err = service.DeleteDevice(database, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "device deleted"})
	})

	router.Run(":8080")
}

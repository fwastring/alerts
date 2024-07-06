package main

import (
	// "errors"
	// "fmt"
	"log"
	"net/http"

	// badger "github.com/dgraph-io/badger/v4"
	db "github.com/fwastring/alerts/database"
	"github.com/gin-gonic/gin"
)

type Alert struct {
    Name     string  `json:"name"`
    Instance  string  `json:"instance"`
}

func resolveAlert(c *gin.Context) {
	name := c.Param("name")

	err := db.Delete(name)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	c.IndentedJSON(http.StatusOK, nil)
}


func getAlerts(c *gin.Context) {
	alerts, err := db.GetAll()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	c.IndentedJSON(http.StatusOK, alerts)
}

func getAlertByName(c *gin.Context) {
    name := c.Param("name")

	alert, err := db.Get(name)
	if err != nil {
		log.Fatalf("error: %v\n", err)
		c.IndentedJSON(http.StatusNotFound, nil)
	}
	c.IndentedJSON(http.StatusOK, Alert{Name: name, Instance: alert})
}

func postAlert(c *gin.Context) {
    var newAlert Alert

    if err := c.BindJSON(&newAlert); err != nil {
        return
    }
	err := db.Set(newAlert.Name, newAlert.Instance)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	c.IndentedJSON(http.StatusOK, newAlert)

}

func main() {
	router := gin.Default()
    router.GET("/alerts", getAlerts)
	router.GET("/alerts/:name", getAlertByName)
    router.POST("/alerts", postAlert)
	router.DELETE("/alerts/:name", resolveAlert)

    router.Run("localhost:8080")
}

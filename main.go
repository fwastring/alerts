package main

import (
	// "errors"
	// "fmt"
	"log"
	"net/http"

	// badger "github.com/dgraph-io/badger/v4"
	// "fwastring/database"
	"github.com/gin-gonic/gin"
)

type Alert struct {
    Name     string  `json:"name"`
    Instance  string  `json:"instance"`
}


func getAlerts(c *gin.Context) {
	alerts, err := getAll()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	c.IndentedJSON(http.StatusOK, alerts)

}

func getAlertByName(c *gin.Context) {
    name := c.Param("name")

	alert, err := get(name)
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
	err := set(newAlert.Name, newAlert.Instance)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	c.IndentedJSON(http.StatusOK, newAlert)

}

func main() {
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
    router.GET("/alerts", getAlerts)
	router.GET("/alerts/:name", getAlertByName)
    router.POST("/alerts", postAlert)

    router.Run("localhost:8080")
}

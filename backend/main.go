package main

import (
	_ "database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
	"io"
	"log"
	"net/http"
	"pala/backend/models"
)

func deleteVoterByID(c *gin.Context) {
	// vid := c.Param("id")
	// id, err := models.DeleteVoter(&vid)

	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No Voters Of That ID"})
	// 	return
	// } else {
	// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted voter"})
	// }

}

func getVoters(c *gin.Context) {
	voters, err := models.QueryVoters()
	if err != nil {
		log.Fatal(err)
		return
	}

	if voters == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No Voters Have Voted Yet"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, voters)
	}
}

func getVoterByID(c *gin.Context) {
	id := c.Param("id")
	voter, err := models.GetVoterByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "voter not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, voter)
}

func updateVoters(c *gin.Context) {
	var newVoter models.Voter

	err := c.BindJSON(&newVoter)
	if err != nil {
		return
	}

	_, err = models.AddVoter(&newVoter)

	if err != nil {
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, &newVoter)
}

func main() {
	err := models.ConnectDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/voters/", getVoters)
	r.POST("/voters/", updateVoters)
	r.GET("/voters/:id", getVoterByID)
	// r.POST("/voters/:id", deleteVoterByID)
	r.Run("localhost:8080")
}

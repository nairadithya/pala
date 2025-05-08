package app

import (
	"log"
	"net/http"
	"pala/backend/pkg/api"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "voter API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreateVoter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newVoter api.NewVoterRequest
		if err := c.ShouldBindJSON(&newVoter); err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload", "error": err.Error()})
			return
		}

		voterID, err := s.voterService.New(newVoter)

		if err != nil {
			log.Printf("service error: %v", err)
			statusCode := http.StatusInternalServerError
			errorMessage := "Failed to create voter"
			c.JSON(statusCode, gin.H{"message": errorMessage, "error": err.Error()})
			return
		}

		response := gin.H{
			"status":   "success",
			"message":  "New voter created successfully",
			"voter_id": voterID,
		}

		c.JSON(http.StatusOK, response)
	}
}

// func (s *Server) CreateVoter() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Header("Content-Type", "application/json")

// 		var newVoter api.NewVoterRequest
// 		err := c.ShouldBindJSON(&newVoter)

// 		if err != nil {
// 			log.Printf("handler error: %v", err)
// 			c.JSON(http.StatusBadRequest, nil)
// 			return
// 		}

// 		err = s.voterService.New(newVoter)

// 		if err != nil {
// 			log.Printf("service error: %v", err)
// 			c.JSON(http.StatusInternalServerError, nil)
// 			return
// 		}

// 		response := map[string]string{
// 			"status": "success",
// 			"data":   "new voter created",
// 		}

// 		c.JSON(http.StatusOK, response)
// 	}
// }

func (s *Server) CreateVote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newVote api.NewVoteRequest
		err := c.ShouldBindJSON(&newVote)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.voteService.New(newVote)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new vote created",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetVoterInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var voterID int
		idStr := c.Param("id")
		if idStr == "" {
			log.Printf("handler error: missing voter ID in path")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing voter ID in path"})
			return
		}

		voterID, err := strconv.Atoi(idStr)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		voterInfo, err := s.voterService.GetVoterInfo(voterID)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   voterInfo,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetAllCandidateInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		candidateInfo, err := s.candidateService.GetAllCandidates()

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   candidateInfo,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetAllParties() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("content-type", "application/json")

		parties, err := s.partyService.GetAllParties()

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   parties,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) VoterHasVoted() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		idStr := c.Param("id")
		if idStr == "" {
			log.Printf("handler error: missing voter ID in path")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing voter ID in path"})
			return
		}

		voterID, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid voter ID format"})
			return
		}

		hasVotedResult, err := s.voterService.HasVoted(voterID)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check voting status"})
			return
		}

		response := map[string]interface{}{
			"status":   "success",
			"hasVoted": hasVotedResult,
		}

		c.JSON(http.StatusOK, response)
	}
}

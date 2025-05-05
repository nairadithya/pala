package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"pala/backend/pkg/api"
)

type Server struct {
	router           *gin.Engine
	voterService     api.VoterService
	voteService      api.VoteService
	partyService     api.PartyService
	candidateService api.CandidateService
	electionService  api.ElectionService
}

func NewServer(router *gin.Engine, voterService api.VoterService, voteService api.VoteService, partyService api.PartyService, candidateService api.CandidateService, electionService api.ElectionService) *Server {
	return &Server{
		router:           router,
		voterService:     voterService,
		voteService:      voteService,
		partyService:     partyService,
		candidateService: candidateService,
		electionService:  electionService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}

package api

import "errors"

type CandidateService interface {
	New(request NewCandidateRequest) error
	GetCandidateInfo(candidateID int) (Candidate, error)
	GetAllCandidates() ([]Candidate, error)
}

type CandidateRepository interface {
	CreateCandidate(request NewCandidateRequest) error
	GetCandidate(candidateID int) (Candidate, error)
	GetAllCandidates() ([]Candidate, error)
}

type candidateService struct {
	storage CandidateRepository
}

func NewCandidateService(repo CandidateRepository) CandidateService {
	return &candidateService{storage: repo}
}

func (c *candidateService) New(request NewCandidateRequest) error {
	if request.FirstName == "" || request.LastName == "" {
		return errors.New("candidate: first and last name are required")
	}

	return c.storage.CreateCandidate(request)
}

func (c *candidateService) GetCandidateInfo(candidateID int) (Candidate, error) {
	return c.storage.GetCandidate(candidateID)
}

func (c *candidateService) GetAllCandidates() ([]Candidate, error) {
	return c.storage.GetAllCandidates()
}

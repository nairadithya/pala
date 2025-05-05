package api

import "time"

type Party struct {
	PartyID   int       `json:"party_id"`
	PartyName string    `json:"party_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Election struct {
	ElectionID   int       `json:"election_id"`
	ElectionName string    `json:"election_name"`
	ElectionDate time.Time `json:"election_date"`
	Description  *string   `json:"description,omitempty"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Voter struct {
	VoterID          int       `json:"voter_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	RegistrationDate time.Time `json:"registration_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type NewVoterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Candidate struct {
	CandidateID int       `json:"candidate_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	ElectionID  int       `json:"election_id"`
	PartyID     *int      `json:"party_id,omitempty"`
	Bio         *string   `json:"bio,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Election    *Election `json:"election,omitempty"`
	Party       *Party    `json:"party,omitempty"`
}

type Vote struct {
	VoteID        int        `json:"vote_id"`
	ElectionID    int        `json:"election_id"`
	CandidateID   int        `json:"candidate_id"`
	VoterID       int        `json:"voter_id"`
	VoteTimestamp time.Time  `json:"vote_timestamp"`
	Election      *Election  `json:"election,omitempty"`
	Candidate     *Candidate `json:"candidate,omitempty"`
	Voter         *Voter     `json:"voter,omitempty"`
}

type NewVoteRequest struct {
	VoteID      int `json:"vote_id"`
	ElectionID  int `json:"election_id"`
	CandidateID int `json:"candidate_id"`
	VoterID     int `json:"voter_id"`
}

package api

import "time"

type Party struct {
	PartyID          int       `json:"party_id"`
	PartyName        string    `json:"party_name"`
	PartyDescription *string   `json:"party_description,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type NewPartyRequest struct {
	PartyName        string  `json:"party_name"`
	PartyDescription *string `json:"party_description,omitempty"`
}

type Voter struct {
	VoterID       int       `json:"voter_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	ContactNumber string    `json:"contact_number"`
}

type NewVoterRequest struct {
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	ContactNumber string    `json:"contact_number"`
}

type Election struct {
	ElectionID   int       `json:"election_id"`
	ElectionName int       `json:"election_name"`
	ElectionDate time.Time `json:"election_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type NewElectionRequest struct {
	ElectionName int       `json:"election_name"`
	ElectionDate time.Time `json:"election_date"`
}

type Candidate struct {
	CandidateID int       `json:"candidate_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PartyID     *int      `json:"party_id,omitempty"`
	Bio         *string   `json:"bio,omitempty"`
	Party       *Party    `json:"party,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewCandidateRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	PartyID   *int    `json:"party_id,omitempty"`
	Bio       *string `json:"bio,omitempty"`
	Party     *Party  `json:"party,omitempty"`
}

type Vote struct {
	VoteID        int       `json:"vote_id"`
	CandidateID   int       `json:"candidate_id"`
	VoterID       int       `json:"voter_id"`
	VoteTimestamp time.Time `json:"vote_timestamp"`
}

type NewVoteRequest struct {
	VoteID      int `json:"vote_id"`
	ElectionID  int `json:"election_id"`
	CandidateID int `json:"candidate_id"`
	VoterID     int `json:"voter_id"`
}

package repository

import (
	"database/sql"
	"errors"
	"log"
	"pala/backend/pkg/api"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Storage interface {
	RunMigrations(connectionString string) error
	CreateVoter(request api.NewVoterRequest) (int, error)
	CreateVoteEntry(request api.NewVoteRequest) error
	CreateElection(request api.NewElectionRequest) error
	CreateParty(request api.NewPartyRequest) error
	CreateCandidate(request api.NewCandidateRequest) error
	GetVoter(voterID int) (api.Voter, error)
	GetVoteByVoter(voterID int) (api.Vote, error)
	GetCandidate(candidateID int) (api.Candidate, error)
	GetElection(electionID int) (api.Election, error)
	GetAllCandidates() ([]api.Candidate, error)
	GetParty(candidateID int) (api.Party, error)
	GetAllParties() ([]api.Party, error)
	HasVoted(voterID int) (bool, error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) RunMigrations(connectionString string) error {
	if connectionString == "" {
		return errors.New("repository: the connString was empty")
	}

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	migrationsPath := filepath.Join("file://", basePath, "/pkg/repository/migrations/")

	m, err := migrate.New(migrationsPath, connectionString)

	if err != nil {
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil
	}

	return nil
}

func (s *storage) CreateVoter(request api.NewVoterRequest) (int, error) {
	newVoterStatement := `
		INSERT INTO "voter" (first_name, last_name, date_of_birth, contact_number)
		VALUES ($1, $2, $3, $4) RETURNING voter_id;
		`

	var voterID int
	err := s.db.QueryRow(newVoterStatement, request.FirstName, request.LastName, request.DateOfBirth, request.ContactNumber).Scan(&voterID)

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return 0, err
	}

	return voterID, nil
}

func (s *storage) CreateVoteEntry(request api.NewVoteRequest) error {
	newVoteStatement := `
		INSERT INTO vote (party_id, voter_id)
		VALUES ($1, $2)
		RETURNING vote_id;
		`

	err := s.db.QueryRow(newVoteStatement, request.PartyID, request.VoterID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}

	return nil
}

func (s *storage) CreateElection(request api.NewElectionRequest) error {
	newVoteStatement := `
		INSERT INTO "election" (election_name, election_date)
		VALUES ($1, $2)
		RETURNING election_id;
		`

	err := s.db.QueryRow(newVoteStatement, request.ElectionName, request.ElectionDate).Err()

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}

	return nil
}

func (s *storage) CreateParty(request api.NewPartyRequest) error {
	newVoteStatement := `
		INSERT INTO "party" (party_name, party_description)
		VALUES ($1, $2)
		RETURNING election_id;
		`

	err := s.db.QueryRow(newVoteStatement, request.PartyName, request.PartyDescription).Err()

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}

	return nil
}

func (s *storage) CreateCandidate(request api.NewCandidateRequest) error {
	query := `
		INSERT INTO candidate (first_name, last_name, party_id, bio)
		VALUES ($1, $2, $3, $4)
		RETURNING candidate_id;
	`

	err := s.db.QueryRow(query, request.FirstName, request.LastName, request.PartyID, request.Bio).Err()
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}
	return nil
}

func (s *storage) GetVoter(voterID int) (api.Voter, error) {
	getVoterStatement := `
		SELECT voter_id, first_name, last_name, date_of_birth, contact_number FROM "voter"
		where voter_id=$1;
		`

	var voter api.Voter
	err := s.db.QueryRow(getVoterStatement, voterID).Scan(&voter.VoterID, &voter.FirstName, &voter.LastName, &voter.DateOfBirth, &voter.ContactNumber)

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return api.Voter{}, err
	}

	return voter, nil
}

func (s *storage) GetVoteByVoter(voterID int) (api.Vote, error) {
	getVoteStatement := `
		SELECT vote_id, candidate_id, voter_id, vote_timestamp
		WHERE voter_id = $1;
						`

	var vote api.Vote
	err := s.db.QueryRow(getVoteStatement, voterID).Scan(&vote.VoteID, &vote.PartyID, &vote.VoterID, &vote.VoteTimestamp)

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return api.Vote{}, err
	}
	return vote, nil
}

func (s *storage) GetCandidate(candidateID int) (api.Candidate, error) {
	query := `
		SELECT candidate_id, first_name, last_name, party_id, bio, created_at, updated_at
		FROM candidate
		WHERE candidate_id = $1;
	`
	var c api.Candidate
	err := s.db.QueryRow(query, candidateID).Scan(
		&c.CandidateID,
		&c.FirstName,
		&c.LastName,
		&c.PartyID,
		&c.Bio,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return api.Candidate{}, err
	}
	return c, nil
}

func (s *storage) GetElection(electionID int) (api.Election, error) {
	query := `
		SELECT election_id, election_name, election_date, created_at, updated_at
		FROM election
		WHERE election_id = $1;
	`
	var e api.Election
	err := s.db.QueryRow(query, electionID).Scan(
		&e.ElectionID,
		&e.ElectionName,
		&e.ElectionDate,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return api.Election{}, err
	}
	return e, nil
}

func (s *storage) GetParty(partyID int) (api.Party, error) {
	query := `
		SELECT party_id, party_name, created_at, updated_at
		FROM party
		WHERE party_id = $1;
	`
	var p api.Party
	err := s.db.QueryRow(query, partyID).Scan(
		&p.PartyID,
		&p.PartyName,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return api.Party{}, err
	}
	return p, nil
}

func (s *storage) GetAllCandidates() ([]api.Candidate, error) {
	query := `
		SELECT candidate_id, first_name, last_name, party_id, bio
		FROM candidate;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return nil, err
	}
	defer rows.Close()

	var candidates []api.Candidate
	for rows.Next() {
		var c api.Candidate
		err := rows.Scan(
			&c.CandidateID,
			&c.FirstName,
			&c.LastName,
			&c.PartyID,
			&c.Bio,
		)
		if err != nil {
			log.Printf("this was the error: %v", err.Error())
			return nil, err
		}
		candidates = append(candidates, c)
	}

	if err := rows.Err(); err != nil {
		log.Printf("this was the error: %v", err.Error())
		return nil, err
	}

	return candidates, nil
}

func (s *storage) GetAllParties() ([]api.Party, error) {
	query := `
		SELECT party_id, party_name, party_description, created_at, updated_at
		FROM party;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return nil, err
	}
	defer rows.Close()

	var parties []api.Party
	for rows.Next() {
		var p api.Party
		err := rows.Scan(
			&p.PartyID,
			&p.PartyName,
			&p.PartyDescription,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			log.Printf("this was the error: %v", err.Error())
			return nil, err
		}
		parties = append(parties, p)
	}

	if err := rows.Err(); err != nil {
		log.Printf("this was the error: %v", err.Error())
		return nil, err
	}

	return parties, nil
}

func (s *storage) HasVoted(voterID int) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM vote
			WHERE voter_id = $1
		);
	`

	var hasVoted bool
	err := s.db.QueryRow(query, voterID).Scan(&hasVoted)

	if err != nil {
		log.Printf("Error checking if voter has voted: %v", err.Error())
		return false, err
	}

	return hasVoted, nil
}
